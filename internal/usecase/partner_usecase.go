package usecase

import (
	"context"
	"fmt"
	"loan-app/internal/entity"
	"loan-app/internal/model"
	"loan-app/internal/repository"
	ulid "loan-app/pkg/database/gorm"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PartnerUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	PartnerRepository *repository.PartnerRepository
}

func NewPartnerUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	userRepository *repository.PartnerRepository,
) *PartnerUseCase {
	return &PartnerUseCase{
		DB:                db,
		Log:               log,
		PartnerRepository: userRepository,
	}
}

func (a *PartnerUseCase) Create(ctx context.Context, id ulid.ULID, request *model.CreatePartnerRequest) error {
	method := "PartnerUseCase.Create"
	log := a.Log.WithField("method", method)
	log.Trace("[BEGIN]")
	log.WithField("request", id).Debug("request")

	db := a.DB.WithContext(ctx)

	partner := entity.NewPartner(&entity.CreatePartnerProps{
		Name:        request.Name,
		PartnerType: request.Type,
		CreatedBy:   id,
	})

	if err := a.PartnerRepository.Create(db, partner); err != nil {
		if strings.Contains(err.Error(), "uq_partner_name") {
			return fmt.Errorf("partner/already-exists")
		}
		panic(err)
	}

	a.Log.Trace("[END]")
	return nil
}

func (a *PartnerUseCase) Lists(ctx context.Context) ([]entity.Partner, error) {
	method := "PartnerUseCase.Lists"
	log := a.Log.WithField("method", method)
	log.Trace("[BEGIN]")

	db := a.DB.WithContext(ctx)
	partners := make([]entity.Partner, 0)
	if err := a.PartnerRepository.FindAll(db, &partners); err != nil {
		panic(err)
	}

	a.Log.Trace("[END]")
	return partners, nil
}
