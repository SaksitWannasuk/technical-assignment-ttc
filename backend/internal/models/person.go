package models

import "time"

type Person struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	BirthDate time.Time `json:"birthDate"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
