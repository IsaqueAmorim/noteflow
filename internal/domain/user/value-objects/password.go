package valueobjects

import (
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	hash string
	salt string
}

func NewPassword(plainText string) *Password {
	validate(plainText)

	hash, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		return &Password{}
	}

	return &Password{
		hash: string(hash),
		salt: "",
	}
}

func validate(s string) {

	// if len(s) < 8 {
	// 	println("Error: Password must be at least 8 characters long")
	// }
	// if !strings.ContainsAny(s, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
	// 	println("Error: Password must contain at least one uppercase letter")
	// }
	// if !strings.ContainsAny(s, "abcdefghijklmnopqrstuvwxyz") {
	// 	println("Error: Password must contain at least one lowercase letter")
	// }
	// if !strings.ContainsAny(s, "0123456789") {
	// 	println("Error: Password must contain at least one number")
	// }
	// if !strings.ContainsAny(s, "!@#$%^&*()-_=+[]{}|;:'\",.<>?/`~") {
	// 	println("Error: Password must contain at least one special character")
	// }
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
