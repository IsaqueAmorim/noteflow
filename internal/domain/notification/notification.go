package notification

import (
	"fmt"
)

type Notification struct {
	errors []error
}

func NewNotification() *Notification {
	return &Notification{}
}

func (n *Notification) AddError(err error) {
	n.errors = append(n.errors, err)
}

func (n *Notification) HasErrors() bool {
	return len(n.errors) > 0
}

func (n *Notification) String() string {
	var err string

	for _, error := range n.errors {
		err += fmt.Sprintf("%s\n", error.Error())
	}

	return err
}

func (n *Notification) CountErrors() int {
	return len(n.errors)
}

func (n *Notification) Errors() []error {
	return n.errors
}

func (n *Notification) Merge(notification *Notification) {
	for _, err := range notification.errors {
		n.AddError(err)
	}
}

func (n *Notification) Clear() {
	n.errors = []error{}
}
