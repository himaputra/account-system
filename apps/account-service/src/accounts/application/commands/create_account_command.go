package command

type CreateAccountCommand struct {
	UserId string
}

func NewCreateAccountCommand(userId string) *CreateAccountCommand {
	return &CreateAccountCommand{
		UserId: userId,
	}
}
