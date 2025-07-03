package model

type CreateAssetRequest struct {
	Name      string `json:"name" validate:"required,min=3,max=100"`
	Price     int64  `json:"price" validate:"required,min=1"`
	PartnerID string `json:"partner_id" validate:"required,ulid"`
}
