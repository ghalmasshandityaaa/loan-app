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
	log := logrus.WithField("method", method)
	log.Trace("[BEGIN]")
	log.WithField("request", id).Debug("request")

	db := a.DB.WithContext(ctx)

	user := new(entity.User)
	if err := a.UserRepository.FindById(db, user, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user/not-found")
		}
		panic(err)
	}

	a.Log.Trace("[END]")
	return user, nil
}

func GetByNIK(ctx context.Context, db *gorm.DB, nik string) (*entity.User, error) {
	method := "UserUseCase.GetByNIK"
	log := logrus.WithField("method", method)
	log.Trace("[BEGIN]")
	log.WithField("request", nik).Debug("request")

	user := new(entity.User)
	if err := db.WithContext(ctx).Where("nik = ?", nik).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user/not-found")
		}
		panic(err)
	}

	log.Trace("[END]")
	return user, nil
}
