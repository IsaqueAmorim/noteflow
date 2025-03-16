package notification

import (
	"errors"
	"testing"
)

func TestNewNotification(t *testing.T) {
	notification := NewNotification()
	if notification == nil {
		t.Error("Expected NewNotification to return a non-nil instance")
	}
	if len(notification.Errors()) != 0 {
		t.Error("Expected new notification to have no errors")
	}
}

func TestAddError(t *testing.T) {
	notification := NewNotification()
	err := errors.New("test error")
	notification.AddError(err)

	if len(notification.Errors()) != 1 {
		t.Errorf("Expected 1 error, got %d", len(notification.Errors()))
	}
	if notification.Errors()[0] != err {
		t.Error("Expected error to be added to notification")
	}
}

func TestHasErrors(t *testing.T) {
	notification := NewNotification()
	if notification.HasErrors() {
		t.Error("Expected HasErrors to return false for empty notification")
	}

	notification.AddError(errors.New("test error"))
	if !notification.HasErrors() {
		t.Error("Expected HasErrors to return true after adding an error")
	}
}

func TestString(t *testing.T) {
	notification := NewNotification()
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	notification.AddError(err1)
	notification.AddError(err2)

	expected := "error 1\nerror 2\n"
	if notification.String() != expected {
		t.Errorf("Expected String to return '%s', got '%s'", expected, notification.String())
	}
}

func TestCountErrors(t *testing.T) {
	notification := NewNotification()
	if notification.CountErrors() != 0 {
		t.Errorf("Expected 0 errors, got %d", notification.CountErrors())
	}

	notification.AddError(errors.New("test error"))
	if notification.CountErrors() != 1 {
		t.Errorf("Expected 1 error, got %d", notification.CountErrors())
	}
}

func TestMerge(t *testing.T) {
	notification1 := NewNotification()
	notification2 := NewNotification()

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	notification1.AddError(err1)
	notification2.AddError(err2)

	notification1.Merge(notification2)

	if notification1.CountErrors() != 2 {
		t.Errorf("Expected 2 errors after merge, got %d", notification1.CountErrors())
	}
	if notification1.Errors()[1] != err2 {
		t.Error("Expected second error to be merged from notification2")
	}
}

func TestClear(t *testing.T) {
	notification := NewNotification()
	notification.AddError(errors.New("test error"))

	notification.Clear()
	if notification.CountErrors() != 0 {
		t.Errorf("Expected 0 errors after Clear, got %d", notification.CountErrors())
	}
	if notification.HasErrors() {
		t.Error("Expected HasErrors to return false after Clear")
	}
}
