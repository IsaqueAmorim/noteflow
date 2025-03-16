package user

import (
	"errors"
	"strings"
	"time"

	"github.com/IsaqueAmorim/noteflow/internal/domain/notification"
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

func NewUser(username, emailAdress, password string, role Role) (*User, *notification.Notification) {

	pass, notification := valueobjects.NewPassword(password)
	email, emailNotification := valueobjects.NewEmail(emailAdress)

	notification.Merge(emailNotification)

	user := User{
		id:           uuid.New().String(),
		username:     username,
		email:        email,
		password:     pass,
		createAt:     time.Now().UTC(),
		updatedAt:    time.Now().UTC(),
		role:         role,
		isActive:     false,
		lastActiveAt: time.Now().UTC(),
	}

	notification.Merge(user.validate())

	return &user, notification
}

func (u *User) validate() *notification.Notification {
	notification := notification.NewNotification()

	if strings.TrimSpace(u.username) == "" {
		notification.AddError(errors.New("username cannot be empty"))
	}

	// if u.email == nil || !u.email.IsVerified() {
	// 	notification.AddError(errors.New("invalid email address"))
	// }

	// if u.password == nil || !u.password.IsValid() {
	// 	notification.AddError(errors.New("invalid password"))
	// }

	if u.role != admin && u.role != user {
		notification.AddError(errors.New("invalid role"))
	}

	if u.createAt.IsZero() {
		notification.AddError(errors.New("creation date cannot be zero"))
	}

	if u.updatedAt.IsZero() {
		notification.AddError(errors.New("updated date cannot be zero"))
	}

	if u.lastActiveAt.IsZero() {
		notification.AddError(errors.New("last active date cannot be zero"))
	}

	return notification
}

func (u *User) ID() string {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Email() *valueobjects.Email {
	return u.email
}

func (u *User) Password() *valueobjects.Password {
	return u.password
}

func (u *User) CreatedAt() time.Time {
	return u.createAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) Role() Role {
	return u.role
}

func (u *User) IsActive() bool {
	return u.isActive
}

func (u *User) ActiveAt() time.Time {
	return u.activeAt
}

func (u *User) LastActiveAt() time.Time {
	return u.lastActiveAt
}

func (u *User) Activate() {
	u.isActive = true
	u.activeAt = time.Now().UTC()
	u.validate()
}

func (u *User) ChangeEmail(emailAdress string) *notification.Notification {
	email, notification := valueobjects.NewEmail(emailAdress)

	if notification.HasErrors() {
		return notification
	}

	u.email = email
	u.updatedAt = time.Now().UTC()

	notification = u.validate()

	return notification
	//TODO: LOG EMAIL CHANGED (USER ID, OLD EMAIL, NEW EMAIL)
}

func (u *User) ChangePassword(password string) *notification.Notification {

	pass, notification := valueobjects.NewPassword(password)

	if notification.HasErrors() {
		return notification
	}

	u.password = pass
	u.updatedAt = time.Now().UTC()

	notification = u.validate()

	return notification
}

func (u *User) ChangeRole(role Role) *notification.Notification {
	u.role = role
	u.updatedAt = time.Now().UTC()
	return u.validate()
}

func (u *User) ChangeUsername(username string) *notification.Notification {
	u.username = username
	u.updatedAt = time.Now().UTC()
	return u.validate()
}
