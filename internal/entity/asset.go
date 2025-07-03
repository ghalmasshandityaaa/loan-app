package entity

import (
	"time"

	"loan-app/pkg/database/gorm"

	"github.com/oklog/ulid/v2"
)

type Asset struct {
	ID        gorm.ULID  `json:"id" gorm:"column:id;type:varchar(26);primaryKey"`
	PartnerID gorm.ULID  `json:"partner_id" gorm:"column:partner_id;type:varchar(26);not null"`
	Name      string     `json:"name" gorm:"column:name;type:varchar(100);not null"`
	Price     int64      `json:"price" gorm:"column:price;type:bigint;not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP"`
	CreatedBy gorm.ULID  `json:"created_by" gorm:"column:created_by;type:varchar(26);not null"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;type:datetime"`
	UpdatedBy *gorm.ULID `json:"updated_by" gorm:"column:updated_by;type:varchar(26)"`
}

type CreateAssetProps struct {
	Name      string
	PartnerID gorm.ULID
	Price     int64
	CreatedBy gorm.ULID
}

func NewAsset(props *CreateAssetProps) *Asset {
	return &Asset{
		ID:        gorm.ULID(ulid.Make()),
		Name:      props.Name,
		PartnerID: props.PartnerID,
		Price:     props.Price,
		CreatedAt: time.Now(),
		CreatedBy: props.CreatedBy,
		UpdatedAt: nil,
		UpdatedBy: nil,
	}
}

func (e *Asset) TableName() string {
	return "assets"
}
