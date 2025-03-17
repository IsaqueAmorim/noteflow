package notification_test

import (
	"errors"
	"testing"

	"github.com/IsaqueAmorim/noteflow/internal/domain/notification"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNotification_AddError(t *testing.T) {
	n := notification.NewNotification()
	err := errors.New("test error")

	n.AddError(err)

	require.Equal(t, 1, n.CountErrors())
	assert.Equal(t, err, n.Errors()[0])
}

func TestNotification_HasErrors(t *testing.T) {
	n := notification.NewNotification()

	assert.False(t, n.HasErrors())

	n.AddError(errors.New("test error"))

	assert.True(t, n.HasErrors())
}

func TestNotification_String(t *testing.T) {
	n := notification.NewNotification()
	n.AddError(errors.New("error 1"))
	n.AddError(errors.New("error 2"))

	expected := "error 1\nerror 2\n"
	assert.Equal(t, expected, n.String())
}

func TestNotification_CountErrors(t *testing.T) {
	n := notification.NewNotification()

	assert.Equal(t, 0, n.CountErrors())

	n.AddError(errors.New("test error"))

	assert.Equal(t, 1, n.CountErrors())
}

func TestNotification_Errors(t *testing.T) {
	n := notification.NewNotification()
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	n.AddError(err1)
	n.AddError(err2)

	assert.Equal(t, []error{err1, err2}, n.Errors())
}

func TestNotification_Merge(t *testing.T) {
	n1 := notification.NewNotification()
	n2 := notification.NewNotification()

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	n1.AddError(err1)
	n2.AddError(err2)

	n1.Merge(n2)

	assert.Equal(t, 2, n1.CountErrors())
	assert.Contains(t, n1.Errors(), err1)
	assert.Contains(t, n1.Errors(), err2)
}

func TestNotification_Clear(t *testing.T) {
	n := notification.NewNotification()
	n.AddError(errors.New("test error"))

	require.True(t, n.HasErrors())

	n.Clear()

	assert.False(t, n.HasErrors())
	assert.Equal(t, 0, n.CountErrors())
}
