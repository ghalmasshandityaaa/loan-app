package model

type SignInResponseWrapper struct {
	Ok   bool
	Data SignInResponse
}

type SignUpResponseWrapper struct {
	Ok   bool
	Data SignUpResponse
}
