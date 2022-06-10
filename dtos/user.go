package dtos

type User struct {
	Name     string `json:"name,omitempty"`
	UserName string `json:"user_name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
}
