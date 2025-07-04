package repository

import (
	"loan-app/internal/entity"

	ulid "loan-app/pkg/database/gorm"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	Repository[entity.Transaction]
	Log *logrus.Logger
}

func NewTransactionRepository(log *logrus.Logger) *TransactionRepository {
	return &TransactionRepository{
		Log: log,
	}
}

func (r *TransactionRepository) FindUserTransactions(db *gorm.DB, entities *[]entity.Transaction, userID ulid.ULID) error {
	return db.Debug().Find(entities).Where("user_id = ?", userID).Error
}
