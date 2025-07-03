package entity

import (
	"time"

	"loan-app/pkg/database/gorm"

	"github.com/oklog/ulid/v2"
)

// PartnerType defines the enum for partner types
type PartnerType string

const (
	PartnerTypeEcommerce PartnerType = "ecommerce"
	PartnerTypeDealer    PartnerType = "dealer"
)

type Partner struct {
	ID          gorm.ULID   `json:"id" gorm:"column:id;type:varchar(26);primaryKey"`
	Name        string      `json:"name" gorm:"column:name;type:varchar(100);not null"`
	PartnerType PartnerType `json:"partner_type" gorm:"column:partner_type;type:enum('ecommerce','dealer');not null"`
	CreatedAt   time.Time   `json:"created_at" gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP"`
	CreatedBy   gorm.ULID   `json:"created_by" gorm:"column:created_by;type:varchar(26);not null"`
	UpdatedAt   *time.Time  `json:"updated_at" gorm:"column:updated_at;type:datetime"`
	UpdatedBy   *gorm.ULID  `json:"updated_by" gorm:"column:updated_by;type:varchar(26)"`
}

type CreatePartnerProps struct {
	Name        string
	PartnerType PartnerType
	CreatedBy   gorm.ULID
}

func NewPartner(props *CreatePartnerProps) *Partner {
	return &Partner{
		ID:          gorm.ULID(ulid.Make()),
		Name:        props.Name,
		PartnerType: PartnerType(props.PartnerType),
		CreatedAt:   time.Now(),
		CreatedBy:   props.CreatedBy,
		UpdatedAt:   nil,
		UpdatedBy:   nil,
	}
}

func (e *Partner) TableName() string {
	return "partners"
}
