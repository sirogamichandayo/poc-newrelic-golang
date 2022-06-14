package services

import (
	"context"
	"errors"
	"net/http"

	apperrors "github.com/dijsilva/golang-api-newrelic/app_errors"
	"github.com/dijsilva/golang-api-newrelic/dtos"
	"github.com/dijsilva/golang-api-newrelic/entities"
	"github.com/dijsilva/golang-api-newrelic/pkg/github"
	"github.com/dijsilva/golang-api-newrelic/repository"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type UserService interface {
	Create(user dtos.User, ctx context.Context) apperrors.AppError
	List(ctx context.Context) ([]dtos.User, apperrors.AppError)
}

type userService struct {
	repository   repository.UserRepository
	githubClient github.GithubClient
}

func CreateUserService(repository repository.UserRepository, githubClient github.GithubClient) UserService {
	return &userService{
		repository:   repository,
		githubClient: githubClient,
	}
}

func (s *userService) Create(user dtos.User, ctx context.Context) apperrors.AppError {
	txn := newrelic.FromContext(ctx)
	defer txn.StartSegment("services.user.Create").End()

	userAlreadyInDatabase, err := s.repository.FindByUserName(user.UserName, ctx)

	if err != nil {
		return apperrors.AppError{
			Err:       err,
			ErrStatus: http.StatusInternalServerError,
		}
	}

	if userAlreadyInDatabase.ID != "" {
		return apperrors.AppError{
			Err:       errors.New("resource already in database"),
			ErrStatus: http.StatusConflict,
		}
	}

	githubUserData, appError := s.githubClient.GetUserData(user.UserName, ctx)

	if appError.Err != nil {
		return apperrors.AppError{
			Err:       errors.New("error to try get user data from github"),
			ErrStatus: http.StatusInternalServerError,
		}
	}

	userEntity := &entities.User{
		Name:      user.Name,
		UserName:  user.UserName,
		Age:       user.Age,
		Email:     user.Email,
		AvatarUrl: githubUserData.AvatarUrl,
		Location:  githubUserData.Location,
	}

	err = s.repository.Create(userEntity, ctx)
	if err != nil {
		return apperrors.AppError{
			Err:       err,
			ErrStatus: http.StatusInternalServerError,
		}
	}
	return apperrors.AppError{}
}

func (s *userService) List(ctx context.Context) ([]dtos.User, apperrors.AppError) {
	txn := newrelic.FromContext(ctx)
	defer txn.StartSegment("service.user.List").End()
	users, err := s.repository.Find(ctx)

	if err != nil {
		return nil, apperrors.AppError{
			Err:       err,
			ErrStatus: http.StatusInternalServerError,
		}
	}

	usersDTO := []dtos.User{}
	for _, user := range users {
		userDTO := dtos.User{
			Name:     user.Name,
			UserName: user.UserName,
			Age:      user.Age,
			Email:    user.Email,
			GithubUserData: dtos.GithubUserData{
				AvatarUrl: user.AvatarUrl,
				Location:  user.Location,
			},
		}
		usersDTO = append(usersDTO, userDTO)
	}

	return usersDTO, apperrors.AppError{}
}
