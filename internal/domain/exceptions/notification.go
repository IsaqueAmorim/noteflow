package exceptions

// type Notification struct {

// }

var errors []error

type Notification interface {
	Add()
	AddRange()
	IsEmpty()
	GetErrors()
	Clean()
}

func Add(err error) {
	errors = append(errors, err)
}

func AddRange(errs []error) {

	for _, err := range errs {
		errors = append(errors, err)
	}

}

func IsEmpty() bool {
	return !(len(errors) > 0)
}
