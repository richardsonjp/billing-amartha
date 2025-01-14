package model

import (
	"time"
)

func (Borrower) TableName() string {
	return "borrower"
}

// nb: if using custom type data, eg=json, array, etc. always define the type for gorm
type Borrower struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"column:name"`
	Status    string    `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime"`
}
