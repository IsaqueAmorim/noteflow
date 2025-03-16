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
	id        string
	username  string
	email     *valueobjects.Email
	password  *valueobjects.Password
	createAt  time.Time
	updatedAt time.Time
	role      Role
	isActive  bool
	activeAt  time.Time
}

func NewUser(username, emailAdress, password string, role Role) (*User, *notification.Notification) {

	pass, notification := valueobjects.NewPassword(password)

	user := User{
		id:        uuid.New().String(),
		username:  username,
		email:     valueobjects.NewEmail(emailAdress),
		password:  pass,
		createAt:  time.Now().UTC(),
		updatedAt: time.Now().UTC(),
		role:      role,
		isActive:  false,
		activeAt:  time.Now().UTC(),
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

	if u.role != admin && u.role != user {
		notification.AddError(errors.New("invalid role"))
	}

	if u.createAt.IsZero() {
		notification.AddError(errors.New("creation date cannot be zero"))
	}

	if u.updatedAt.IsZero() {
		notification.AddError(errors.New("updated date cannot be zero"))
	}

	if u.activeAt.IsZero() {
		notification.AddError(errors.New("active date cannot be zero"))
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

func (u *User) Activate() *notification.Notification {

	notification := notification.NewNotification()

	if u.isActive {
		notification.AddError(errors.New("user is already active"))
		return notification
	}

	u.isActive = true
	u.validate()

	return notification
}

func (u *User) ChageEmail(emailAdress string) {

	u.email = valueobjects.NewEmail(emailAdress)
	u.updatedAt = time.Now().UTC()
	u.validate()
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

func (u *User) Deactivate() *notification.Notification {

	notification := notification.NewNotification()

	if !u.isActive {
		notification.AddError(errors.New("user is already deactivated"))
		return notification
	}

	u.isActive = false
	u.updatedAt = time.Now().UTC()
	u.validate()

	return notification
}

func (u *User) UpdateActiveAt() {
	u.activeAt = time.Now().UTC()
	u.updatedAt = time.Now().UTC()
}
