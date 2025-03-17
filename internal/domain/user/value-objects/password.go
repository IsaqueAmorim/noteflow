package valueobjects

import (
	"errors"
	"strings"

	"github.com/IsaqueAmorim/noteflow/internal/domain/notification"
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	hash string
	salt string
}

var bcryptFunc = bcrypt.GenerateFromPassword

func NewPassword(plainText string) (*Password, *notification.Notification) {
	notification := validate(plainText)

	if notification.HasErrors() {
		return &Password{}, notification
	}

	hash, err := bcryptFunc([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		notification.AddError(err)
		return &Password{}, notification
	}

	return &Password{
		hash: string(hash),
		salt: "",
	}, notification
}

func validate(s string) *notification.Notification {

	notification := notification.NewNotification()

	if strings.TrimSpace(s) == "" {
		notification.AddError(errors.New("Password cannot be empty"))
		return notification
	}
	if len(s) < 8 {
		notification.AddError(errors.New("Password must be at least 8 characters long"))
	}
	if !strings.ContainsAny(s, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		notification.AddError(errors.New("Password must contain at least one uppercase letter"))
	}
	if !strings.ContainsAny(s, "abcdefghijklmnopqrstuvwxyz") {
		notification.AddError(errors.New("Password must contain at least one lowercase letter"))
	}
	if !strings.ContainsAny(s, "0123456789") {
		notification.AddError(errors.New("Password must contain at least one number"))
	}
	if !strings.ContainsAny(s, "!@#$%^&*()-_=+[]{}|;:'\",.<>?/`~") {
		notification.AddError(errors.New("Password must contain at least one special character"))
	}

	return notification
}

func (p Password) Check(plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.hash), []byte(plainText))
	return err == nil
}

func (p Password) Hash() string {
	return p.hash
}

func (p Password) Salt() string {
	return p.salt
}

func (p Password) IsValid() bool {
	return true
}
