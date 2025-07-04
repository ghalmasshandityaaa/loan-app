package usecase

import (
	"context"
	"errors"
	"fmt"
	"loan-app/internal/entity"
	"loan-app/internal/model"
	"loan-app/internal/repository"
	ulid "loan-app/pkg/database/gorm"
	"sync"
	"time"

	v2 "github.com/oklog/ulid/v2"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	DefaultTimeout    = 2 * time.Second
	AssetFetchTimeout = 1 * time.Second
	LimitFetchTimeout = 1 * time.Second
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
	log := a.Log.WithField("method", method)
	log.Trace("[BEGIN]")
	log.WithField("request", id).Debug("request")

	// Create context with timeout
	ctx, cancel := context.WithTimeout(ctx, DefaultTimeout)
	defer cancel()

	// Start database transaction
	tx := a.DB.WithContext(ctx).Begin()
	db := a.DB.WithContext(ctx)
	defer tx.Rollback()

	// Use channels to collect results from goroutines
	assetChan := make(chan *entity.Asset, 1)
	limitsChan := make(chan []entity.CustomerLimit, 1)
	errChan := make(chan error, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine 1: Fetch asset concurrently
	go func() {
		defer wg.Done()

		// Create separate context with timeout for asset fetch
		assetCtx, assetCancel := context.WithTimeout(ctx, AssetFetchTimeout)
		defer assetCancel()

		asset := new(entity.Asset)
		assetID := ulid.ULID(v2.MustParse(request.AssetID))

		if err := a.AssetRepository.FindById(db, asset, assetID); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				errChan <- fmt.Errorf("asset/not-found")
				return
			}
			log.WithError(err).Error("failed to find asset")
			errChan <- fmt.Errorf("asset fetch failed: %w", err)
			return
		}

		select {
		case assetChan <- asset:
		case <-assetCtx.Done():
			errChan <- fmt.Errorf("asset fetch timeout: %w", assetCtx.Err())
		}
	}()

	// Goroutine 2: Fetch customer limits concurrently
	go func() {
		defer wg.Done()

		// Create separate context with timeout for limits fetch
		limitsCtx, limitsCancel := context.WithTimeout(ctx, LimitFetchTimeout)
		defer limitsCancel()

		limits := make([]entity.CustomerLimit, 0)

		if err := a.CustomerLimitRepository.ListUserLimits(db, &limits, id); err != nil {
			log.WithError(err).Error("failed to find user limit")
			errChan <- fmt.Errorf("customer limits fetch failed: %w", err)
			return
		}

		select {
		case limitsChan <- limits:
		case <-limitsCtx.Done():
			errChan <- fmt.Errorf("limits fetch timeout: %w", limitsCtx.Err())
		}
	}()

	// Wait for goroutines to complete or timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	var asset *entity.Asset
	var limits []entity.CustomerLimit

	// Collect results with timeout handling
	select {
	case <-done:
		// All goroutines completed, collect results
		select {
		case asset = <-assetChan:
		default:
			// Asset fetch failed, check error channel
		}

		select {
		case limits = <-limitsChan:
		default:
			// Limits fetch failed, check error channel
		}

		// Check for any errors
		select {
		case err := <-errChan:
			return err
		default:
			// No errors
		}

	case <-ctx.Done():
		return fmt.Errorf("operation timeout: %w", ctx.Err())
	}

	// Validate results
	if asset == nil {
		return fmt.Errorf("asset/not-found")
	}

	if len(limits) == 0 {
		return fmt.Errorf("customer-limit/not-found")
	}

	// Find tenor limit
	var limit *entity.CustomerLimit
	for _, l := range limits {
		if l.Tenor == request.Tenor {
			limit = &l
			break
		}
	}

	if limit == nil {
		return fmt.Errorf("customer-limit/not-found")
	}

	// Create transaction entity
	transaction, err := entity.NewTransaction(&entity.CreateTransactionProps{
		UserID: id,
		Asset:  asset,
		Limit:  limit,
	})
	if err != nil {
		log.WithError(err).Error("failed to create transaction")
		panic(err)
	}

	// Perform database operations with context cancellation check
	if err := ctx.Err(); err != nil {
		panic("operation cancelled: " + err.Error())
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

	log.Trace("[END]")
	return nil
}

func (a *TransactionUseCase) Lists(ctx context.Context, userID ulid.ULID) ([]entity.Transaction, error) {
	method := "TransactionUseCase.Lists"
	log := a.Log.WithField("method", method)
	log.Trace("[BEGIN]")

	db := a.DB.WithContext(ctx)
	transactions := make([]entity.Transaction, 0)
	if err := a.TransactionRepository.FindUserTransactions(db, &transactions, userID); err != nil {
		panic(err)
	}

	a.Log.Trace("[END]")
	return transactions, nil
}
