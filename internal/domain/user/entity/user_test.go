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
		user, notification := NewUser("testuser", email.Address(), pass, user)

		assert.NotNil(t, user)
		assert.Empty(t, notification.Errors())
		assert.Equal(t, "testuser", user.Username())
		assert.Equal(t, email.Address(), user.Email().Address())
		assert.True(t, user.Password().Check(pass))
		assert.Equal(t, user.role, user.Role())
		assert.False(t, user.IsActive())
		assert.NotZero(t, user.CreatedAt())
		assert.NotZero(t, user.UpdatedAt())
	})

	t.Run("should return String for invalid user fields", func(t *testing.T) {
		_, notification := NewUser("", "asdasd@qwasd.com", "StrongPassword123", 99) // Invalid role

		assert.NotNil(t, notification)
		assert.True(t, notification.HasErrors())
		assert.Contains(t, notification.String(), "username cannot be empty")
		//assert.Contains(t, notification.String(), "invalid email address")
		//assert.Contains(t, notification.String(), "invalid password")
		assert.Contains(t, notification.String(), "invalid role")
	})
}

func TestActivateUser(t *testing.T) {
	t.Run("should activate valid user", func(t *testing.T) {
		user, notification := NewUser("testuser", "test@example.com", "StrongPassword123!", user)

		user.Activate()

		assert.Empty(t, notification.Errors())
		assert.True(t, user.IsActive())
		assert.NotZero(t, user.ActiveAt())
	})
}

func TestChangeEmail(t *testing.T) {
	t.Run("should change email successfully", func(t *testing.T) {
		user, notification := NewUser("testuser", "test@example.com", "StrongPassword123!", user)

		assert.Empty(t, notification.Errors())

		newEmail := "newemail@example.com"
		user.ChageEmail(newEmail)

		assert.Equal(t, newEmail, user.Email().Address())
		assert.Empty(t, notification.Errors())
		assert.NotZero(t, user.UpdatedAt())
	})

	t.Run("should return error for invalid email", func(t *testing.T) {
		user, errs := NewUser("testuser", "test@example.com", "StrongPassword123!", user)

		user.ChageEmail("invalid-email")
		errs.Clear()
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

		user, notification := NewUser("testuser", "test@example.com", pass, user)

		assert.Empty(t, notification.Errors())

		notification = user.ChangePassword(newPass)

		assert.Empty(t, notification.Errors())
		assert.True(t, user.Password().Check(newPass))
		assert.NotZero(t, user.UpdatedAt())
	})

	t.Run("should return error for invalid password", func(t *testing.T) {
		user, notification := NewUser("testuser", "test@example.com", "StrongPassword123!", user)

		assert.Empty(t, notification.Errors())

		notification = user.ChangePassword("")

		assert.NotNil(t, notification)
		assert.True(t, notification.HasErrors())
		assert.Equal(t, 1, notification.CountErrors())
		assert.Contains(t, notification.String(), "Password cannot be empty")
	})
}

func TestChangeRole(t *testing.T) {
	t.Run("should change role successfully", func(t *testing.T) {
		user, notification := NewUser("testuser", "test@example.com", "StrongPassword123!", user)

		assert.Empty(t, notification.Errors())

		notification = user.ChangeRole(admin)

		assert.Empty(t, notification.Errors())
		assert.Equal(t, admin, user.Role())
		assert.NotZero(t, user.UpdatedAt())
	})

	t.Run("should return error for invalid role", func(t *testing.T) {
		user, notification := NewUser("testuser", "test@example.com", "StrongPassword123!", user)

		assert.Empty(t, notification.Errors())

		notification = user.ChangeRole(99)

		assert.NotNil(t, notification)
		assert.True(t, notification.HasErrors())
		assert.Equal(t, 1, notification.CountErrors())
		assert.Contains(t, notification.String(), "invalid role")
	})
}

func TestChangeUsername(t *testing.T) {
	t.Run("should change username successfully", func(t *testing.T) {
		user, notification := NewUser("testuser", "test@example.com", "StrongPassword123!", user)

		assert.Empty(t, notification.Errors())

		newUsername := "newusername"
		notification = user.ChangeUsername(newUsername)

		assert.Empty(t, notification.Errors())
		assert.Equal(t, newUsername, user.Username())
		assert.NotZero(t, user.UpdatedAt())
	})

	t.Run("should return error for empty username", func(t *testing.T) {
		user, notification := NewUser("testuser", "test@example.com", "StrongPassword123!", user)

		assert.Empty(t, notification.Errors())

		notification = user.ChangeUsername("")

		assert.NotNil(t, notification)
		assert.True(t, notification.HasErrors())
		assert.Equal(t, 1, notification.CountErrors())
		assert.Contains(t, notification.String(), "username cannot be empty")
	})
}
