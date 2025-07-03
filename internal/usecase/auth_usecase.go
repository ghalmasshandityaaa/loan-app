package usecase

import (
	"context"
	"errors"
	"fmt"
	"loan-app/internal/entity"
	"loan-app/internal/model"
	"loan-app/internal/repository"
	"loan-app/internal/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	JwtUtil        *utils.JwtUtil
	UserRepository *repository.UserRepository
}

func NewAuthUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	jwtUtil *utils.JwtUtil,
	userRepository *repository.UserRepository,
) *AuthUseCase {
	return &AuthUseCase{
		DB:             db,
		Log:            log,
		JwtUtil:        jwtUtil,
		UserRepository: userRepository,
	}
}

func (a *AuthUseCase) SignIn(ctx context.Context, request *model.SignInRequest) (string, string, error) {
	method := "AuthUseCase.SignIn"
	a.Log.WithField("method", method).Trace("[BEGIN]")
	a.Log.WithField("method", method).WithField("request", request).Debug("request")

	db := a.DB.WithContext(ctx)

	user := new(entity.User)
	if err := a.UserRepository.GetByUsername(db, user, request.Username); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", "", fmt.Errorf("user/not-found")
		}
		panic(err)
	}

	match := utils.ComparePassword(request.Password, user.Password)
	if !match {
		a.Log.WithContext(ctx).Error("password mismatch")
		return "", "", fmt.Errorf("auth/password-mismatch")
	}

	a.Log.WithContext(ctx).Debug("Password match, generating tokens...")

	// Use goroutines to generate access and refresh tokens in parallel
	type tokenResult struct {
		token string
		err   error
	}

	accessTokenChan := make(chan tokenResult, 1)
	refreshTokenChan := make(chan tokenResult, 1)

	// Generate access token in goroutine
	go func() {
		token, err := a.JwtUtil.GenerateAccessToken(user)
		accessTokenChan <- tokenResult{token: token, err: err}
	}()

	// Generate refresh token in goroutine
	go func() {
		token, err := a.JwtUtil.GenerateRefreshToken(user)
		refreshTokenChan <- tokenResult{token: token, err: err}
	}()

	// Wait for both tokens to be generated
	var accessToken, refreshToken string
	var accessErr, refreshErr error

	// Collect results from both goroutines
	for i := 0; i < 2; i++ {
		select {
		case result := <-accessTokenChan:
			if result.err != nil {
				a.Log.WithContext(ctx).WithError(result.err).Error("failed to generate access token")
				accessErr = fmt.Errorf("internal/server-error")
			} else {
				accessToken = result.token
			}
		case result := <-refreshTokenChan:
			if result.err != nil {
				a.Log.WithContext(ctx).WithError(result.err).Error("failed to generate refresh token")
				refreshErr = fmt.Errorf("internal/server-error")
			} else {
				refreshToken = result.token
			}
		case <-ctx.Done():
			return "", "", ctx.Err()
		}
	}

	a.Log.WithContext(ctx).Debug("Tokens generated successfully, checking for errors...")

	// Check for error
	if accessErr != nil {
		return "", "", accessErr
	}
	if refreshErr != nil {
		return "", "", refreshErr
	}

	a.Log.WithContext(ctx).WithField("method", method).Trace("[END]")

	return accessToken, refreshToken, nil
}
