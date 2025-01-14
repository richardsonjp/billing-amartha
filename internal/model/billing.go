package model

import (
	"time"
)

func (Billing) TableName() string {
	return "billing"
}

// nb: if using custom type data, eg=json, array, etc. always define the type for gorm
type Billing struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	LoanID    string    `gorm:"column:loan_id"`
	Amount    int64     `gorm:"column:loan_id"`
	Status    string    `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime"`
}
