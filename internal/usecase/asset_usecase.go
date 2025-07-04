package usecase

import (
	"context"
	"errors"
	"fmt"
	"loan-app/internal/entity"
	"loan-app/internal/model"
	"loan-app/internal/repository"
	ulid "loan-app/pkg/database/gorm"
	"strings"

	v2 "github.com/oklog/ulid/v2"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AssetUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	AssetRepository   *repository.AssetRepository
	PartnerRepository *repository.PartnerRepository
}

func NewAssetUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	userRepository *repository.AssetRepository,
	partnerRepository *repository.PartnerRepository,
) *AssetUseCase {
	return &AssetUseCase{
		DB:                db,
		Log:               log,
		AssetRepository:   userRepository,
		PartnerRepository: partnerRepository,
	}
}

func (a *AssetUseCase) Create(ctx context.Context, id ulid.ULID, request *model.CreateAssetRequest) error {
	method := "AssetUseCase.Create"
	log := a.Log.WithField("method", method)
	log.Trace("[BEGIN]")
	log.WithField("request", id).Debug("request")

	db := a.DB.WithContext(ctx)

	// Validate partner exists
	partner := new(entity.Partner)
	if err := a.PartnerRepository.FindById(db, partner, ulid.ULID(v2.MustParse(request.PartnerID))); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("partner/not-found")
		}
		panic(err)
	}

	asset := entity.NewAsset(&entity.CreateAssetProps{
		Name:      request.Name,
		PartnerID: ulid.ULID(v2.MustParse(request.PartnerID)),
		Price:     request.Price,
		CreatedBy: id,
	})

	if err := a.AssetRepository.Create(db, asset); err != nil {
		if strings.Contains(err.Error(), "uq_asset") {
			return fmt.Errorf("asset/already-exists")
		}
		panic(err)
	}

	a.Log.Trace("[END]")
	return nil
}

func (a *AssetUseCase) Lists(ctx context.Context) ([]entity.Asset, error) {
	method := "AssetUseCase.Lists"
	log := a.Log.WithField("method", method)
	log.Trace("[BEGIN]")

	db := a.DB.WithContext(ctx)
	assets := make([]entity.Asset, 0)
	if err := a.AssetRepository.FindAll(db, &assets); err != nil {
		panic(err)
	}

	a.Log.Trace("[END]")
	return assets, nil
}
