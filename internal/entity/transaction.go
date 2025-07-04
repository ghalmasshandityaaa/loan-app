package entity

import (
	"fmt"
	"time"

	"loan-app/pkg/database/gorm"

	"github.com/oklog/ulid/v2"
)

type Transaction struct {
	ID                gorm.ULID  `json:"id" gorm:"column:id;type:varchar(26);primaryKey"`
	UserID            gorm.ULID  `json:"user_id" gorm:"column:user_id;type:varchar(26);not null"`
	AssetID           gorm.ULID  `json:"asset_id" gorm:"column:asset_id;type:varchar(26);not null"`
	ContractNumber    string     `json:"contract_number" gorm:"column:contract_number;type:varchar(50);not null;unique"`
	OtrPrice          int64      `json:"otr_price" gorm:"column:otr_price;type:bigint;not null"`
	AdminFee          int64      `json:"admin_fee" gorm:"column:admin_fee;type:bigint;not null"`
	InstallmentAmount int64      `json:"installment_amount" gorm:"column:installment_amount;type:bigint;not null"`
	InterestAmount    int64      `json:"interest_amount" gorm:"column:interest_amount;type:bigint;not null"`
	CreatedAt         time.Time  `json:"created_at" gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP"`
	CreatedBy         gorm.ULID  `json:"created_by" gorm:"column:created_by;type:varchar(26);not null"`
	UpdatedAt         *time.Time `json:"updated_at" gorm:"column:updated_at;type:datetime"`
	UpdatedBy         *gorm.ULID `json:"updated_by" gorm:"column:updated_by;type:varchar(26)"`
}

type CreateTransactionProps struct {
	UserID gorm.ULID
	Asset  *Asset
	Limit  *CustomerLimit
}

func NewTransaction(props *CreateTransactionProps) (*Transaction, error) {
	// tenor = interest
	// admin fee = 0.05 * otr price

	interestAmount := int64(props.Limit.Tenor) * (props.Asset.Price * int64(props.Limit.Tenor) / 100)
	adminFee := int64(float64(props.Asset.Price) * 0.05)
	InstallmentAmount := (props.Asset.Price + interestAmount + adminFee)

	transaction := &Transaction{
		ID:                gorm.ULID(ulid.Make()),
		UserID:            props.UserID,
		AssetID:           props.Asset.ID,
		ContractNumber:    "TRX-" + props.Asset.ID.String() + "-" + time.Now().Format("20060102150405"),
		OtrPrice:          props.Asset.Price,
		AdminFee:          adminFee,
		InstallmentAmount: InstallmentAmount,
		InterestAmount:    interestAmount,
		CreatedAt:         time.Now(),
		CreatedBy:         props.UserID,
		UpdatedAt:         nil,
		UpdatedBy:         nil,
	}

	if props.Limit.AvailableAmount < InstallmentAmount {
		return nil, fmt.Errorf("transaction/limit-exceeded")
	}

	err := props.Limit.CreateTransaction(InstallmentAmount, props.UserID)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (e *Transaction) TableName() string {
	return "transactions"
}
