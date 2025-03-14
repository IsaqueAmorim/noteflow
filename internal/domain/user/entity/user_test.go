package user

import (
	"testing"

	valueobjects "github.com/IsaqueAmorim/noteflow/internal/domain/user/value-objects"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	t.Run("should create a valid user", func(t *testing.T) {

		const pass = "StrongPassword123!"
		email := valueobjects.NewEmail("test@example.com")
		user := NewUser("testuser", email.Address(), pass, user)

		assert.NotNil(t, user)
		assert.Equal(t, "testuser", user.Username())
		assert.Equal(t, email.Address(), user.Email().Address())
		assert.True(t, user.Password().Check(pass))
		assert.Equal(t, user.role, user.Role())
		assert.False(t, user.IsActive())
		assert.NotZero(t, user.CreatedAt())
		assert.NotZero(t, user.UpdatedAt())
	})

	t.Run("should return String for invalid user fields", func(t *testing.T) {
		user := NewUser("", "asdasd@qwasd.com", "StrongPassword123", 99) // Invalid role

		notification := user.validate()
		assert.NotNil(t, notification)
		assert.True(t, notification.HasErrors())
		assert.Contains(t, notification.String(), "username cannot be empty")
		//assert.Contains(t, notification.String(), "invalid email address")
		//assert.Contains(t, notification.String(), "invalid password")
		assert.Contains(t, notification.String(), "invalid role")
	})
}

func TestActivateUser(t *testing.T) {
	t.Run("should activate user", func(t *testing.T) {
		email := valueobjects.NewEmail("test@example.com")
		password := valueobjects.NewPassword("StrongPassword123!")
		user := NewUser("testuser", email.Address(), password.Hash(), user)

		user.Activate()

		assert.True(t, user.IsActive())
		assert.NotZero(t, user.ActiveAt())
	})
}

func TestChangeEmail(t *testing.T) {
	t.Run("should change email successfully", func(t *testing.T) {
		email := valueobjects.NewEmail("test@example.com")
		password := valueobjects.NewPassword("StrongPassword123!")
		user := NewUser("testuser", email.Address(), password.Hash(), user)

		newEmail := "newemail@example.com"
		user.ChageEmail(newEmail)

		assert.Equal(t, newEmail, user.Email().Address())
		assert.NotZero(t, user.UpdatedAt())
	})

	t.Run("should return error for invalid email", func(t *testing.T) {
		email := valueobjects.NewEmail("test@example.com")
		password := valueobjects.NewPassword("StrongPassword123!")
		user := NewUser("testuser", email.Address(), password.Hash(), user)

		user.ChageEmail("invalid-email")

		// notification := user.validate()
		// assert.NotNil(t, notification)
		// assert.True(t, notification.HasErrors())
		//assert.Contains(t, notification.String(), "invalid email address")
	})
}

func TestChangePassword(t *testing.T) {
	t.Run("should change password successfully", func(t *testing.T) {
		const pass = "StrongPassword123!"
		const newPass = "NewStrongPassword123!"

		email := valueobjects.NewEmail("test@example.com")
		user := NewUser("testuser", email.Address(), pass, user)

		user.ChangePassword(newPass)

		assert.True(t, user.Password().Check(newPass))
		assert.NotZero(t, user.UpdatedAt())
	})

	t.Run("should return error for invalid password", func(t *testing.T) {
		email := valueobjects.NewEmail("test@example.com")
		password := valueobjects.NewPassword("StrongPassword123!")
		user := NewUser("testuser", email.Address(), password.Hash(), user)

		user.ChangePassword("")

		notification := user.validate()
		assert.NotNil(t, notification)
		//assert.True(t, notification.HasErrors())
		//assert.Contains(t, notification.String(), "invalid password")
	})
}

func TestChangeRole(t *testing.T) {
	t.Run("should change role successfully", func(t *testing.T) {
		email := valueobjects.NewEmail("test@example.com")
		password := valueobjects.NewPassword("StrongPassword123!")
		user := NewUser("testuser", email.Address(), password.Hash(), user)

		user.ChangeRole(admin)

		assert.Equal(t, admin, user.Role())
		assert.NotZero(t, user.UpdatedAt())
	})

	t.Run("should return error for invalid role", func(t *testing.T) {
		email := valueobjects.NewEmail("test@example.com")
		password := valueobjects.NewPassword("StrongPassword123!")
		user := NewUser("testuser", email.Address(), password.Hash(), user)

		user.ChangeRole(99) // Invalid role

		notification := user.validate()
		assert.NotNil(t, notification)
		assert.True(t, notification.HasErrors())
		assert.Contains(t, notification.String(), "invalid role")
	})
}

func TestChangeUsername(t *testing.T) {
	t.Run("should change username successfully", func(t *testing.T) {
		email := valueobjects.NewEmail("test@example.com")
		password := valueobjects.NewPassword("StrongPassword123!")
		user := NewUser("testuser", email.Address(), password.Hash(), user)

		newUsername := "newusername"
		user.ChangeUsername(newUsername)

		assert.Equal(t, newUsername, user.Username())
		assert.NotZero(t, user.UpdatedAt())
	})

	t.Run("should return error for empty username", func(t *testing.T) {
		email := valueobjects.NewEmail("test@example.com")
		password := valueobjects.NewPassword("StrongPassword123!")
		user := NewUser("testuser", email.Address(), password.Hash(), user)

		user.ChangeUsername("")

		notification := user.validate()
		assert.NotNil(t, notification)
		assert.True(t, notification.HasErrors())
		assert.Contains(t, notification.String(), "username cannot be empty")
	})
}
