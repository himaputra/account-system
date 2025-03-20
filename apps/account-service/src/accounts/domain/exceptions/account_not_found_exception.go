package exception

import "fmt"

type AccountNotFoundException struct {
	Message string
}

func NewAccountNotFoundException(message string) *AccountNotFoundException {
	return &AccountNotFoundException{
		Message: message,
	}
}

func (e *AccountNotFoundException) Error() string {
	return fmt.Sprintf("AccountDomainException: %s", e.Message)
}
