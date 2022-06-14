package dtos

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
