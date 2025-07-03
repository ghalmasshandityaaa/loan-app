package entity

import (
	"time"

	"loan-app/pkg/database/gorm"

	"github.com/oklog/ulid/v2"
)

// User represents an user in the system
// swagger:model User
type User struct {
	// Unique identifier for the user
	// example: "01HXYZ123456789ABCDEFGHIJK"
	ID gorm.ULID `json:"id" gorm:"column:id;type:ulid;primaryKey"`

	// Username for authentication (unique)
	// example: "john.doe"
	Username string `json:"username" gorm:"column:username;size:50;not null;unique"`

	// Password hash (not exposed in JSON)
	Password string `json:"-" gorm:"column:password;type:text;not null"`

	// Whether the user has admin privileges
	// example: false
	IsAdmin bool `json:"is_admin" gorm:"column:is_admin;type:boolean;not null;default:false"`

	// Timestamp when the user was created
	// example: "2024-01-15T08:00:00Z"
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp with time zone;not null;default:CURRENT_TIMESTAMP"`

	// Timestamp when the user was last updated
	// example: "2024-01-15T08:00:00Z"
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp with time zone"`
}

// CreateUserProps represents the properties needed to create a new user
// swagger:model CreateUserProps
type CreateUserProps struct {
	// Username for the new user
	Username string
	// Password for the new user
	Password string
}

func NewUser(props *CreateUserProps) *User {
	return &User{
		ID:        gorm.ULID(ulid.Make()),
		Username:  props.Username,
		Password:  props.Password,
		IsAdmin:   false,
		CreatedAt: time.Now(),
	}
}

func (e *User) TableName() string {
	return "users"
}
