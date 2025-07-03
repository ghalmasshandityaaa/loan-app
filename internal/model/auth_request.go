package model

type SignInRequest struct {
	NIK      string `json:"nik" validate:"required,len=16,numeric"`
	Password string `json:"password" validate:"required,is-strong-password"`
}

type SignUpRequest struct {
	NIK            string `json:"nik" validate:"required,len=16,numeric"`
	FullName       string `json:"full_name" validate:"required,min=3,max=100"`
	LegalName      string `json:"legal_name" validate:"required,min=3,max=100"`
	PlaceOfBirth   string `json:"place_of_birth" validate:"required,min=2,max=100"`
	DateOfBirth    string `json:"date_of_birth" validate:"required,is-valid-date"`
	Salary         int64  `json:"salary" validate:"required,min=0"`
	IDCardPhotoURL string `json:"id_card_photo_url" validate:"required,url"`
	SelfiePhotoURL string `json:"selfie_photo_url" validate:"required,url"`
	Password       string `json:"password" validate:"required,is-strong-password"`
}

type VerifyAccountRequest struct {
	Token string `validate:"required"`
}
