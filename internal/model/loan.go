package model

import (
	"time"
)

func (Loan) TableName() string {
	return "loan"
}

// nb: if using custom type data, eg=json, array, etc. always define the type for gorm
type Loan struct {
	ID                string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	BorrowerID        string    `gorm:"column:borrower_id"`
	Amount            int64     `gorm:"column:amount"`
	Interest          int64     `gorm:"column:interest"`
	TotalAmount       int64     `gorm:"column:total_amount"`
	WeeklyInstallment int64     `gorm:"column:weekly_installment"`
	TotalWeeks        int64     `gorm:"column:total_weeks"`
	CreatedAt         time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt         time.Time `gorm:"column:updated_at;type:datetime"`
}
