package exception

import "fmt"

type AccountDomainException struct {
	Message string
}

func NewAccountDomainException(message string) *AccountDomainException {
	return &AccountDomainException{
		Message: message,
	}
}

func (e *AccountDomainException) Error() string {
	return fmt.Sprintf("AccountDomainException: %s", e.Message)
}
