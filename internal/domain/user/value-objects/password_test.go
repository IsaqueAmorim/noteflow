package valueobjects

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPassword(t *testing.T) {
	t.Run("should create a valid password hash", func(t *testing.T) {
		plainText := "Valid123!"
		password, errs := NewPassword(plainText)

		assert.NotEmpty(t, password.Hash())
		assert.Empty(t, errs.Errors())
		assert.True(t, password.Check(plainText))
	})

	t.Run("should return an empty password on bcrypt error", func(t *testing.T) {
		// Simulate bcrypt error by passing an invalid cost (not directly testable here)
		// This would require mocking bcrypt.GenerateFromPassword
	})
}

func TestValidate(t *testing.T) {
	t.Run("should return no errors for a valid password", func(t *testing.T) {
		_, notification := NewPassword("Valid123!")
		assert.False(t, notification.HasErrors())
	})

	t.Run("should return error for password shorter than 8 characters", func(t *testing.T) {
		_, notification := NewPassword("Short1!")
		assert.True(t, notification.HasErrors())
		assert.Contains(t, notification.String(), "Password must be at least 8 characters long")
	})

	t.Run("should return error for password without uppercase letter", func(t *testing.T) {
		_, notification := NewPassword("lowercase1!")
		assert.True(t, notification.HasErrors())
		assert.Equal(t, 1, notification.CountErrors())
		assert.Contains(t, notification.String(), "Password must contain at least one uppercase letter")
	})

	t.Run("should return error for password without lowercase letter", func(t *testing.T) {
		_, notification := NewPassword("UPPERCASE1!")
		assert.True(t, notification.HasErrors())
		assert.Equal(t, 1, notification.CountErrors())
		assert.Contains(t, notification.String(), "Password must contain at least one lowercase letter")
	})

	t.Run("should return error for password without a number", func(t *testing.T) {
		_, notification := NewPassword("NoNumber!")
		assert.True(t, notification.HasErrors())
		assert.Equal(t, 1, notification.CountErrors())
		assert.Contains(t, notification.String(), "Password must contain at least one number")
	})

	t.Run("should return error for password without special character", func(t *testing.T) {
		_, notification := NewPassword("NoSpecial1")
		assert.True(t, notification.HasErrors())
		assert.Equal(t, 1, notification.CountErrors())
		assert.Contains(t, notification.String(), "Password must contain at least one special character")
	})
}

func TestPasswordCheck(t *testing.T) {
	t.Run("should return true for matching password", func(t *testing.T) {
		plainText := "Valid123!"
		password, errs := NewPassword(plainText)

		assert.Empty(t, errs.Errors())
		assert.True(t, password.Check(plainText))
	})

	t.Run("should return false for non-matching password", func(t *testing.T) {
		plainText := "Valid123!"
		password, errs := NewPassword(plainText)

		assert.Empty(t, errs.Errors())
		assert.False(t, password.Check("Invalid123!"))
	})
}

func TestPasswordHash(t *testing.T) {
	t.Run("should return the hash of the password", func(t *testing.T) {
		plainText := "Valid123!"
		password, errs := NewPassword(plainText)

		assert.Empty(t, errs.Errors())
		assert.NotEmpty(t, password.Hash())
	})
}

func TestPasswordSalt(t *testing.T) {
	t.Run("should return the salt of the password", func(t *testing.T) {
		password, errs := NewPassword("Valid123!")

		assert.Empty(t, errs.String())
		assert.Equal(t, "", password.Salt())
	})
}

func TestPasswordIsValid(t *testing.T) {
	t.Run("should always return true", func(t *testing.T) {
		password, errs := NewPassword("Valid123!")
		assert.Empty(t, errs.Errors())
		assert.True(t, password.IsValid())
	})
}
