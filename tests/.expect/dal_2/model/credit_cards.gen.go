// Code generated by github.com/go-leo/sqlgen. DO NOT EDIT.
// Code generated by github.com/go-leo/sqlgen. DO NOT EDIT.
// Code generated by github.com/go-leo/sqlgen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCreditCard = "credit_cards"

// CreditCard mapped from table <credit_cards>
type CreditCard struct {
	ID            int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"-"`
	CreatedAt     *time.Time     `gorm:"column:created_at" json:"-"`
	UpdatedAt     *time.Time     `gorm:"column:updated_at" json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;index:idx_credit_cards_deleted_at,priority:1" json:"-"`
	Number        *string        `gorm:"column:number" json:"-"`
	CustomerRefer *int64         `gorm:"column:customer_refer" json:"-"`
	BankID        *int64         `gorm:"column:bank_id" json:"-"`
}

// TableName CreditCard's table name
func (*CreditCard) TableName() string {
	return TableNameCreditCard
}
