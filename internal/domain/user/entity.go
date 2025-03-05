package user

import (
	"strings"
	"time"
)

type User struct {
	Username  string
	Email     string
	Password  string
	CreateAt  time.Time
	UpdatedAt time.Time
	Role      string
}

func NewUser(username, email, password, role string) *User {
	user := User{
		Username:  username,
		Email:     email,
		Password:  password,
		CreateAt:  time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Role:      role,
	}

	user.validate()

	return &user
}

func (u *User) validate() {
	if strings.TrimSpace(u.Username) == "" {
		println("Error")
	}
}
