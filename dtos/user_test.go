package dtos

import (
	"testing"
)

func Test_GithubUserData(t *testing.T) {
	gu := GithubUserData{
		AvatarUrl: "https://gitlab.com",
		Location:  "JP",
	}

	b, err := gu.Json()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(b)
}
