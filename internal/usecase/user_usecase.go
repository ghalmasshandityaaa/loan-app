package usecase

import (
	"context"
	"errors"
	"fmt"
	"loan-app/internal/entity"
	"loan-app/internal/repository"
	ulid "loan-app/pkg/database/gorm"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	UserRepository *repository.UserRepository
}

func NewUserUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	userRepository *repository.UserRepository,
) *UserUseCase {
	return &UserUseCase{
		DB:             db,
		Log:            log,
		UserRepository: userRepository,
	}
}

func (a *UserUseCase) GetById(ctx context.Context, id ulid.ULID) (*entity.User, error) {
	method := "UserUseCase.GetById"
	a.Log.WithField("method", method).Trace("[BEGIN]")
	a.Log.WithField("method", method).WithField("request", id).Debug("request")

	db := a.DB.WithContext(ctx)

	user := new(entity.User)
	if err := a.UserRepository.FindById(db, user, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user/not-found")
		}
		panic(err)
	}

	a.Log.WithField("method", method).Trace("[END]")
	return user, nil
}

func (a *UserUseCase) List(ctx context.Context) ([]entity.User, error) {
	method := "UserUseCase.List"
	a.Log.WithField("method", method).Trace("[BEGIN]")
	a.Log.WithField("method", method).Debug("request")

	db := a.DB.WithContext(ctx)

	users := make([]entity.User, 0)
	if err := a.UserRepository.FindAll(db, &users); err != nil {
		panic(err)
	}

	a.Log.WithField("method", method).Trace("[END]")
	return users, nil
}
