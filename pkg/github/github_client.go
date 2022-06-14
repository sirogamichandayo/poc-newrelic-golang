package github

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	apperrors "github.com/dijsilva/golang-api-newrelic/app_errors"
	"github.com/dijsilva/golang-api-newrelic/config"
	"github.com/dijsilva/golang-api-newrelic/dtos"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type GithubClient interface {
	GetUserData(string, context.Context) (dtos.GithubUserData, apperrors.AppError)
}

type githubClient struct{}

func NewGithubClient() GithubClient {
	return &githubClient{}
}

func (g *githubClient) GetUserData(userName string, ctx context.Context) (dtos.GithubUserData, apperrors.AppError) {
	txn := newrelic.FromContext(ctx)
	defer txn.StartSegment("pkg.github.github_client.GetUserData").End()

	endpoint := config.Configuration.GithubUserApiURI + userName
	response, err := makeRequest("GET", endpoint, nil, ctx)
	if err != nil {
		return dtos.GithubUserData{}, apperrors.AppError{
			Err:       err,
			ErrStatus: http.StatusFailedDependency,
		}
	}

	if response.StatusCode == http.StatusNotFound {
		return dtos.GithubUserData{}, apperrors.AppError{
			Err:       errors.New("user not found"),
			ErrStatus: http.StatusNotFound,
		}
	}

	if response.StatusCode != http.StatusOK {
		return dtos.GithubUserData{}, apperrors.AppError{
			Err:       errors.New("error to try get user infomation from github"),
			ErrStatus: http.StatusFailedDependency,
		}
	}

	responseJson, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return dtos.GithubUserData{}, apperrors.AppError{
			Err:       errors.New("error to try read user infomation from github"),
			ErrStatus: http.StatusInternalServerError,
		}
	}

	var githubUserData dtos.GithubUserData
	err = json.Unmarshal(responseJson, &githubUserData)
	if err != nil {
		return dtos.GithubUserData{}, apperrors.AppError{
			Err:       errors.New("error to try unmarshal user infomation from github"),
			ErrStatus: http.StatusInternalServerError,
		}
	}

	return githubUserData, apperrors.AppError{}
}

func makeRequest(method string, url string, body io.Reader, ctx context.Context) (*http.Response, error) {
	txn := newrelic.FromContext(ctx)
	client := &http.Client{}
	client.Transport = newrelic.NewRoundTripper(client.Transport)
	request, _ := http.NewRequest(method, url, body)
	request = newrelic.RequestWithTransactionContext(request, txn)
	return client.Do(request)
}
