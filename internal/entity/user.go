package entity

import (
	"time"

	"loan-app/pkg/database/gorm"

	"github.com/oklog/ulid/v2"
)

type User struct {
	ID             gorm.ULID  `json:"id" gorm:"column:id;type:ulid;primaryKey"`
	NIK            string     `json:"nik" gorm:"column:nik;size:16;not null;unique"`
	FullName       string     `json:"full_name" gorm:"column:full_name;size:100;not null"`
	LegalName      string     `json:"legal_name" gorm:"column:legal_name;size:100;not null"`
	PlaceOfBirth   string     `json:"place_of_birth" gorm:"column:place_of_birth;size:100;not null"`
	DateOfBirth    time.Time  `json:"date_of_birth" gorm:"column:date_of_birth;type:date;not null"`
	Salary         int64      `json:"salary" gorm:"column:salary;type:bigint;not null"`
	IDCardPhotoURL string     `json:"id_card_photo_url" gorm:"column:id_card_photo_url;type:text"`
	SelfiePhotoURL string     `json:"selfie_photo_url" gorm:"column:selfie_photo_url;type:text"`
	Password       string     `json:"-" gorm:"column:password;type:text;not null"`
	IsAdmin        bool       `json:"is_admin" gorm:"column:is_admin;type:boolean;not null;default:false"`
	CreatedAt      time.Time  `json:"created_at" gorm:"column:created_at;type:timestamp with time zone;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt      *time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp with time zone"`
}

type CreateUserProps struct {
	NIK            string
	FullName       string
	LegalName      string
	PlaceOfBirth   string
	DateOfBirth    time.Time
	Salary         int64
	IDCardPhotoURL string
	SelfiePhotoURL string
	Password       string
}

func NewUser(props *CreateUserProps) *User {
	return &User{
		ID:             gorm.ULID(ulid.Make()),
		NIK:            props.NIK,
		FullName:       props.FullName,
		LegalName:      props.LegalName,
		PlaceOfBirth:   props.PlaceOfBirth,
		DateOfBirth:    props.DateOfBirth,
		Salary:         props.Salary,
		IDCardPhotoURL: props.IDCardPhotoURL,
		SelfiePhotoURL: props.SelfiePhotoURL,
		Password:       props.Password,
		IsAdmin:        false,
		CreatedAt:      time.Now(),
	}
}

func (e *User) TableName() string {
	return "users"
}
