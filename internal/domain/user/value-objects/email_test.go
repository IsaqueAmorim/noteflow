package valueobjects

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmail(t *testing.T) {
	t.Run("should create a valid email", func(t *testing.T) {
		email, notification := NewEmail("test@example.com")

		assert.NotNil(t, email)
		assert.Equal(t, "test@example.com", email.Address())
		assert.Equal(t, "test", email.Local())
		assert.False(t, email.IsVerified())
		assert.False(t, notification.HasErrors())
	})

	t.Run("should return error for empty email", func(t *testing.T) {
		email, notification := NewEmail("")

		assert.NotNil(t, email)
		assert.True(t, notification.HasErrors())
		assert.Equal(t, notification.CountErrors(), 1)
		assert.Contains(t, notification.String(), "email address cannot be empty")
	})

	t.Run("should return error for email without @", func(t *testing.T) {
		email, notification := NewEmail("testemail.com")

		assert.NotNil(t, email)
		assert.True(t, notification.HasErrors())
		assert.Equal(t, notification.CountErrors(), 1)
		assert.Contains(t, notification.String(), "invalid email format: '@' is missing")
	})

	t.Run("should return error for email with @ at the beginning", func(t *testing.T) {
		email, notification := NewEmail("@example.com")

		assert.NotNil(t, email)
		assert.True(t, notification.HasErrors())
		assert.Equal(t, notification.CountErrors(), 3)
		assert.Contains(t, notification.String(), "local part of the email must have at least one character")
		assert.Contains(t, notification.String(), "email address does not match the required format")
		assert.Contains(t, notification.String(), "invalid email format: '@' is misplaced")
	})

	t.Run("should return error for email with @ at the end", func(t *testing.T) {
		email, notification := NewEmail("teste@")

		assert.NotNil(t, email)
		assert.True(t, notification.HasErrors())
		assert.Equal(t, notification.CountErrors(), 3)
		assert.Contains(t, notification.String(), "invalid email format: '@' is misplaced")
		assert.Contains(t, notification.String(), "email address does not match the required format")
		assert.Contains(t, notification.String(), "domain part of the email is invalid")
	})

	t.Run("should return error for email with invalid domain format", func(t *testing.T) {
		email, notification := NewEmail("test@example")

		assert.NotNil(t, email)
		assert.True(t, notification.HasErrors())
		assert.Equal(t, notification.CountErrors(), 2)
		assert.Contains(t, notification.String(), "domain part of the email is invalid")
		assert.Contains(t, notification.String(), "email address does not match the required format")
	})

	t.Run("should return error for email with domain too short", func(t *testing.T) {
		email, notification := NewEmail("test@e.c")

		assert.NotNil(t, email)
		assert.True(t, notification.HasErrors())
		assert.Equal(t, notification.CountErrors(), 1)
		assert.Contains(t, notification.String(), "email address does not match the required format")
	})

	t.Run("should handle email with multiple @ symbols", func(t *testing.T) {
		email, notification := NewEmail("test@example@domain.com")

		assert.NotNil(t, email)
		assert.True(t, notification.HasErrors())
		assert.Equal(t, notification.CountErrors(), 1)
		assert.Contains(t, notification.String(), "email address does not match the required format")
	})

	t.Run("should trim whitespace from email address", func(t *testing.T) {
		email, notification := NewEmail("  test@example.com  ")

		assert.NotNil(t, email)
		assert.Equal(t, "test@example.com", email.Address())
		assert.False(t, notification.HasErrors())
	})
}

func TestEmailValidate(t *testing.T) {
	t.Run("should mark email as verified", func(t *testing.T) {
		email, _ := NewEmail("test@example.com")

		assert.NotNil(t, email)
		assert.False(t, email.IsVerified())

		email.Validate()

		assert.True(t, email.IsVerified())
	})
}

func TestEmailAccessors(t *testing.T) {
	t.Run("should return correct address", func(t *testing.T) {
		email, _ := NewEmail("test@example.com")
		assert.NotNil(t, email)
		assert.Equal(t, "test@example.com", email.Address())
	})

	t.Run("should return correct local part", func(t *testing.T) {
		email, _ := NewEmail("test@example.com")
		assert.NotNil(t, email)
		assert.Equal(t, "test", email.Local())
	})

	t.Run("should return correct verification status", func(t *testing.T) {
		email, _ := NewEmail("test@example.com")
		assert.NotNil(t, email)
		assert.False(t, email.IsVerified())
	})
}

func TestComplexEmailScenarios(t *testing.T) {
	t.Run("should accept valid email with dots in local part", func(t *testing.T) {
		email, notification := NewEmail("first.last@example.com")

		assert.NotNil(t, email)
		assert.False(t, notification.HasErrors())
	})

	t.Run("should accept valid email with plus in local part", func(t *testing.T) {
		email, notification := NewEmail("test+tag@example.com")

		assert.NotNil(t, email)
		assert.False(t, notification.HasErrors())
	})

	t.Run("should accept valid email with numbers", func(t *testing.T) {
		email, notification := NewEmail("test123@example123.com")

		assert.NotNil(t, email)
		assert.False(t, notification.HasErrors())
	})

	t.Run("should accept valid email with subdomain", func(t *testing.T) {
		email, notification := NewEmail("test@sub.example.com")

		assert.NotNil(t, email)
		assert.False(t, notification.HasErrors())
	})
}
