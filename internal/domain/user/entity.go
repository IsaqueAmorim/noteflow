package user

import (
	"strings"
	"time"

	valueobjects "github.com/IsaqueAmorim/noteflow/internal/domain/user/value-objects"
	"github.com/google/uuid"
)

type Role uint

const (
	admin Role = iota
	user
)

type User struct {
	id           string
	username     string
	email        *valueobjects.Email
	password     *valueobjects.Password
	createAt     time.Time
	updatedAt    time.Time
	role         Role
	isActive     bool
	activeAt     time.Time
	lastActiveAt time.Time
}

func NewUser(username, emailAdress, password string, role Role) *User {

	user := User{
		id:           uuid.New().String(),
		username:     username,
		email:        valueobjects.NewEmail(emailAdress),
		password:     valueobjects.NewPassword(password),
		createAt:     time.Now().UTC(),
		updatedAt:    time.Now().UTC(),
		role:         role,
		isActive:     false,
		lastActiveAt: time.Now().UTC(),
	}

	user.validate()
	return &user
}

func (u *User) validate() {
	if strings.TrimSpace(u.username) == "" {
		panic("username cannot be empty")
	}

	if u.email == nil || !u.email.IsVerified() {
		panic("invalid email address")
	}

	if u.password == nil || !u.password.IsValid() {
		panic("invalid password")
	}

	if u.role != admin && u.role != user {
		panic("invalid role")
	}

	if u.createAt.IsZero() {
		panic("creation date cannot be zero")
	}

	if u.updatedAt.IsZero() {
		panic("updated date cannot be zero")
	}

	if u.lastActiveAt.IsZero() {
		panic("last active date cannot be zero")
	}
}

func (u *User) Activate() {
	u.isActive = true
	u.activeAt = time.Now().UTC()
	u.validate()
}

func (u *User) ChageEmail(emailAdress string) {

	u.email = valueobjects.NewEmail(emailAdress)
	u.updatedAt = time.Now().UTC()
	u.validate()
	//TODO: LOG EMAIL CHANGED (USER ID, OLD EMAIL, NEW EMAIL)
}

func (u *User) ChangePassword(password string) {
	u.password = valueobjects.NewPassword(password)
	u.updatedAt = time.Now().UTC()
	u.validate()
}

func (u *User) ChangeRole(role Role) {
	u.role = role
	u.updatedAt = time.Now().UTC()
	u.validate()
}

func (u *User) ChangeUsername(username string) {
	u.username = username
	u.updatedAt = time.Now().UTC()
	u.validate()
}
