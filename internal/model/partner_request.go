package model

import "loan-app/internal/entity"

type CreatePartnerRequest struct {
	Name string             `json:"name" validate:"required,min=3,max=100"`
	Type entity.PartnerType `json:"type" validate:"required,oneof=ecommerce dealer"`
}
