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

func (r *UserRepository) GetByUsername(db *gorm.DB, user *entity.User, username string) error {
	return db.Debug().Where("username = ?", username).Take(user).Error
}
