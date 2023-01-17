package dtos

import (
	"encoding/json"
	"fmt"
)

type User struct {
	GithubUserData
	Name     string `json:"name,omitempty"`
	UserName string `json:"user_name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
}

type GithubUserData struct {
	AvatarUrl string `json:"avatar_url"`
	Location  string `json:"location"`
}

func (u GithubUserData) Json() ([]byte, error) {
	return json.Marshal(u)
}

func (u GithubUserData) Str() string {
	return fmt.Sprint(u)
}
