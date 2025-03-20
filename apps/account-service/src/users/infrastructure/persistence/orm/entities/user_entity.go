package entity

type UserEntity struct {
	Id          string `gorm:"primaryKey"`
	Name        string
	Nik         string `gorm:"uniqueIndex"`
	PhoneNumber string `gorm:"uniqueIndex"`
}
