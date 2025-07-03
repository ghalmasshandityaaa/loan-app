package repository

import (
	"loan-app/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{
		Log: log,
	}
}

func (r *UserRepository) GetByNIK(db *gorm.DB, user *entity.User, nik string) error {
	return db.Debug().Where("nik = ?", nik).Take(user).Error
}
