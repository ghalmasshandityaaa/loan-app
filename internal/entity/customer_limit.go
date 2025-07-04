package entity

import (
	"fmt"
	"time"

	"loan-app/pkg/database/gorm"

	"github.com/oklog/ulid/v2"
)

type CustomerLimit struct {
	ID              gorm.ULID  `json:"id" gorm:"column:id;type:ulid;primaryKey"`
	UserID          gorm.ULID  `json:"user_id" gorm:"column:user_id;not null"`
	Tenor           int8       `json:"tenor" gorm:"column:tenor;type:smallint;not null"`
	LimitAmount     int64      `json:"limit_amount" gorm:"column:limit_amount;type:bigint;not null"`
	UsedAmount      int64      `json:"used_amount" gorm:"column:used_amount;type:bigint;not null;default:0"`
	AvailableAmount int64      `json:"available_amount" gorm:"column:available_amount;type:bigint;not null;default:0"`
	CreatedAt       time.Time  `json:"created_at" gorm:"column:created_at;type:timestamp with time zone;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt       *time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp with time zone"`
}

type CreateCustomerLimitProps struct {
	UserID          gorm.ULID
	Tenor           int8
	LimitAmount     int64
	UsedAmount      int64
	AvailableAmount int64
}

func NewCustomerLimit(props *CreateCustomerLimitProps) *CustomerLimit {
	return &CustomerLimit{
		ID:              gorm.ULID(ulid.Make()),
		UserID:          props.UserID,
		Tenor:           props.Tenor,
		LimitAmount:     props.LimitAmount,
		UsedAmount:      props.UsedAmount,
		AvailableAmount: props.AvailableAmount,
		CreatedAt:       time.Now(),
		UpdatedAt:       nil,
	}
}

func (e *CustomerLimit) TableName() string {
	return "customer_limits"
}

func (e *CustomerLimit) CreateTransaction(amount int64, userID gorm.ULID) error {
	now := time.Now()

	e.AvailableAmount -= amount
	e.UsedAmount += amount
	e.UpdatedAt = &now

	if e.AvailableAmount < 0 || e.UsedAmount > e.LimitAmount {
		return fmt.Errorf("customer-limit/insufficient-funds")
	}

	return nil
}
