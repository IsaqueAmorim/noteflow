package valueobjects

import (
	"regexp"
	"strings"
)

type Email struct {
	address    string
	local      string
	domain     string
	isVerified bool
}

func NewEmail(address string) *Email {

	address = strings.TrimSpace(address)
	if address == "" {
		return nil
	}

	atIndex := strings.LastIndex(address, "@")
	if atIndex < 1 || atIndex == len(address)-1 {
		return nil
	}

	local := address[:atIndex]
	domain := address[atIndex+1:]

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(address) {
		return nil
	}

	if len(local) < 1 {
		return nil
	}

	if len(domain) < 3 || !strings.Contains(domain, ".") {
		return nil
	}

	return &Email{
		address:    address,
		local:      local,
		domain:     domain,
		isVerified: false,
	}
}

func (e *Email) validate() {
	if e == nil {
		return
	}
	e.isVerified = true
}

func (e *Email) Validate() {
	e.isVerified = true
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
