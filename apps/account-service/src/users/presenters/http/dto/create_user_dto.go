package dto

type CreateUserDto struct {
	Name        string `json:"nama"`
	Nik         string `json:"nik"`
	PhoneNumber string `json:"no_hp"`
}
