package valueobjects

import (
	"errors"
	"regexp"
	"strings"

	"github.com/IsaqueAmorim/noteflow/internal/domain/notification"
)

type Email struct {
	address    string
	local      string
	domain     string
	isVerified bool
}

func NewEmail(address string) (*Email, *notification.Notification) {
	address = strings.TrimSpace(address)
	notification := validateEmail(address)

	if notification.HasErrors() {
		return &Email{}, notification
	}

	atIndex := strings.LastIndex(address, "@")
	local := address[:atIndex]
	domain := address[atIndex+1:]

	email := &Email{
		address:    address,
		local:      local,
		domain:     domain,
		isVerified: false,
	}

	return email, notification
}

func validateEmail(address string) *notification.Notification {
	notification := notification.NewNotification()

	if address == "" {
		notification.AddError(errors.New("email address cannot be empty"))
		return notification
	}

	atIndex := strings.LastIndex(address, "@")
	if atIndex < 0 {
		notification.AddError(errors.New("invalid email format: '@' is missing"))
		return notification
	}

	if atIndex == len(address)-1 || atIndex == 0 {
		notification.AddError(errors.New("invalid email format: '@' is misplaced"))
	}

	local := address[:atIndex]
	domain := address[atIndex+1:]

	if len(local) < 1 {
		notification.AddError(errors.New("local part of the email must have at least one character"))
	}

	if len(domain) < 3 || !strings.Contains(domain, ".") {
		notification.AddError(errors.New("domain part of the email is invalid"))
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(address) {
		notification.AddError(errors.New("email address does not match the required format"))
	}

	return notification
}

func (e *Email) Address() string {
	return e.address
}

func (e *Email) Local() string {
	return e.local
}

func (e *Email) IsVerified() bool {
	return e.isVerified
}

func (e *Email) Validate() {
	e.isVerified = true
}
