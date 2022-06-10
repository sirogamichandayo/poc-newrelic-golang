package repository

import (
	"context"

	"github.com/dijsilva/golang-api-newrelic/entities"
	"github.com/google/uuid"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type UserRepository interface {
	Create(user *entities.User, ctx context.Context) error
	FindByUserName(userName string, ctx context.Context) (entities.User, error)
	Find(ctx context.Context) ([]entities.User, error)
}

type userRepository struct{}

func CreateUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *entities.User, ctx context.Context) error {
	txn := newrelic.FromContext(ctx)
	defer txn.StartSegment("repository.user.Create").End()
	ctxDB := newrelic.NewContext(ctx, txn)
	tracedDB := database.WithContext(ctxDB)
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	user.ID = uuid.String()

	err = tracedDB.Create(&user).Error
	return err
}

func (r *userRepository) FindByUserName(userName string, ctx context.Context) (userEntity entities.User, err error) {
	txn := newrelic.FromContext(ctx)
	defer txn.StartSegment("repostory.user.FindByUserName").End()
	ctxDB := newrelic.NewContext(ctx, txn)
	tracedDB := database.WithContext(ctxDB)
	err = tracedDB.Where("user_name = ?", userName).Limit(1).Find(&userEntity).Error

	return userEntity, err
}

func (r *userRepository) Find(ctx context.Context) (usersEntity []entities.User, err error) {
	txn := newrelic.FromContext(ctx)
	defer txn.StartSegment("repository.user.Find").End()
	ctxDB := newrelic.NewContext(ctx, txn)
	tracedDB := database.WithContext(ctxDB)
	err = tracedDB.Find(&usersEntity).Error

	return usersEntity, err
}
