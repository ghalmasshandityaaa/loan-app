package usecase

import (
	"context"
	"errors"
	"fmt"
	"loan-app/internal/entity"
	"loan-app/internal/model"
	"loan-app/internal/repository"
	ulid "loan-app/pkg/database/gorm"

	v2 "github.com/oklog/ulid/v2"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TransactionUseCase struct {
	DB                      *gorm.DB
	Log                     *logrus.Logger
	TransactionRepository   *repository.TransactionRepository
	AssetRepository         *repository.AssetRepository
	CustomerLimitRepository *repository.CustomerLimitRepository
}

func NewTransactionUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	userRepository *repository.TransactionRepository,
	assetRepository *repository.AssetRepository,
	customerLimitRepository *repository.CustomerLimitRepository,
) *TransactionUseCase {
	return &TransactionUseCase{
		DB:                      db,
		Log:                     log,
		TransactionRepository:   userRepository,
		AssetRepository:         assetRepository,
		CustomerLimitRepository: customerLimitRepository,
	}
}

func (a *TransactionUseCase) Create(ctx context.Context, id ulid.ULID, request *model.CreateTransactionRequest) error {
	method := "TransactionUseCase.Create"
	log := logrus.WithField("method", method)
	log.Trace("[BEGIN]")
	log.WithField("request", id).Debug("request")

	tx := a.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	asset := new(entity.Asset)
	if err := a.AssetRepository.FindById(tx, asset, ulid.ULID(v2.MustParse(request.AssetID))); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("asset/not-found")
		}
		log.WithError(err).Error("failed to find asset")
		panic(err)
	}

	limits := make([]entity.CustomerLimit, 0)
	if err := a.CustomerLimitRepository.ListUserLimits(tx, &limits, id); err != nil {
		log.WithError(err).Error("failed to find user limit")
		panic(err)
	}

	if len(limits) == 0 {
		return fmt.Errorf("customer-limit/not-found")
	}
	// find tenor limit
	var limit *entity.CustomerLimit
	for _, l := range limits {
		if l.Tenor == request.Tenor {
			limit = &l
			break
		}
	}

	transaction, err := entity.NewTransaction(&entity.CreateTransactionProps{
		UserID: id,
		Asset:  asset,
		Limit:  limit,
	})
	if err != nil {
		log.WithError(err).Error("failed to create transaction")
		return err
	}

	if err := a.TransactionRepository.Create(tx, transaction); err != nil {
		log.WithError(err).Error("failed to create transaction")
		panic(err)
	}

	if err := a.CustomerLimitRepository.Update(tx, limit); err != nil {
		log.WithError(err).Error("failed to update customer limit")
		panic(err)
	}

	if err := tx.Commit().Error; err != nil {
		log.WithError(err).Error("failed to commit transaction")
		panic(err)
	}

	a.Log.Trace("[END]")
	return nil
}

func (a *TransactionUseCase) Lists(ctx context.Context, userID ulid.ULID) ([]entity.Transaction, error) {
	method := "TransactionUseCase.Lists"
	log := logrus.WithField("method", method)
	log.Trace("[BEGIN]")

	db := a.DB.WithContext(ctx)
	transactions := make([]entity.Transaction, 0)
	if err := a.TransactionRepository.FindUserTransactions(db, &transactions, userID); err != nil {
		panic(err)
	}

	a.Log.Trace("[END]")
	return transactions, nil
}
