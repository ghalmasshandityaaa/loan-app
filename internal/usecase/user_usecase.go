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
	DB                      *gorm.DB
	Log                     *logrus.Logger
	UserRepository          *repository.UserRepository
	CustomerLimitRepository *repository.CustomerLimitRepository
}

func NewUserUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	userRepository *repository.UserRepository,
	customerLimitRepository *repository.CustomerLimitRepository,
) *UserUseCase {
	return &UserUseCase{
		DB:                      db,
		Log:                     log,
		UserRepository:          userRepository,
		CustomerLimitRepository: customerLimitRepository,
	}
}

func (a *UserUseCase) GetById(ctx context.Context, id ulid.ULID) (*entity.User, error) {
	method := "UserUseCase.GetById"
	log := a.Log.WithField("method", method)
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

func (a *UserUseCase) FindLimits(ctx context.Context, userID ulid.ULID) ([]entity.CustomerLimit, error) {
	method := "UserUseCase.GetById"
	log := a.Log.WithField("method", method)
	log.Trace("[BEGIN]")
	log.WithField("request", userID).Debug("request")

	db := a.DB.WithContext(ctx)

	limits := make([]entity.CustomerLimit, 0)
	if err := a.CustomerLimitRepository.ListUserLimits(db, &limits, userID); err != nil {
		panic(err)
	}

	a.Log.Trace("[END]")
	return limits, nil
}
