package command

type CreateUserCommand struct {
	Name        string
	Nik         string
	PhoneNumber string
}

func NewCreateUserCommand(name string, nik string, phoneNumber string) *CreateUserCommand {
	return &CreateUserCommand{
		Name:        name,
		Nik:         nik,
		PhoneNumber: phoneNumber,
	}
}
