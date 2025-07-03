package model

import (
	"time"
)

type SignInResponseWrapper struct {
	Ok   bool
	Data SignInResponse
}

type SignUpResponseWrapper struct {
	Ok   bool
	Data SignUpResponse
}

type UserWrapper struct {
	ID             string     `json:"id"`
	NIK            string     `json:"nik"`
	FullName       string     `json:"full_name"`
	LegalName      string     `json:"legal_name"`
	PlaceOfBirth   string     `json:"place_of_birth"`
	DateOfBirth    time.Time  `json:"date_of_birth"`
	Salary         int64      `json:"salary"`
	IDCardPhotoURL string     `json:"id_card_photo_url"`
	SelfiePhotoURL string     `json:"selfie_photo_url"`
	Password       string     `json:"-"`
	IsAdmin        bool       `json:"is_admin"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

type FindSelfResponseWrapper struct {
	Ok   bool
	Data UserWrapper
}

type CustomerLimitWrapper struct {
	ID              string     `json:"id"`
	UserID          string     `json:"user_id"`
	Tenor           int8       `json:"tenor"`
	LimitAmount     int64      `json:"limit_amount"`
	UsedAmount      int64      `json:"used_amount"`
	AvailableAmount int64      `json:"available_amount"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}

type FindCustomerLimitWrapper struct {
	Ok   bool
	Data []CustomerLimitWrapper
}
