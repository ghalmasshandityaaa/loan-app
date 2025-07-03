package repository

import (
	"loan-app/internal/entity"

	"gorm.io/gorm"

	ulid "loan-app/pkg/database/gorm"

	"github.com/sirupsen/logrus"
)

type CustomerLimitRepository struct {
	Repository[entity.CustomerLimit]
	Log *logrus.Logger
}

func NewCustomerLimitRepository(log *logrus.Logger) *CustomerLimitRepository {
	return &CustomerLimitRepository{
		Log: log,
	}
}

func (r *CustomerLimitRepository) ListUserLimits(db *gorm.DB, entities *[]entity.CustomerLimit, userID ulid.ULID) error {
	return db.Debug().Find(entities).Where("user_id = ?", userID).Error
}
