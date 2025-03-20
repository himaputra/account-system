package domain

type User struct {
	Id          string
	Name        string
	Nik         string
	PhoneNumber string
}

func NewUser(id string) *User {
	return &User{Id: id}
}
