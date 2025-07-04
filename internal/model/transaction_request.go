package model

type CreateTransactionRequest struct {
	AssetID string `json:"asset_id" validate:"required,ulid"`
	Tenor   int8   `json:"tenor" validate:"required,min=1,max=4"`
}
