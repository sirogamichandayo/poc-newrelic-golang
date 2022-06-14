package entities

type User struct {
	Base
	Name      string
	UserName  string `gorm:"column:user_name"`
	Age       int
	Email     string
	AvatarUrl string
	Location  string
}
