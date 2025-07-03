package model

import "loan-app/pkg/database/gorm"

type Role string

const (
	RoleAdmin    Role = "admin"
	RoleCustomer Role = "customer"
)

type Auth struct {
	ID      gorm.ULID
	IsAdmin bool
}
