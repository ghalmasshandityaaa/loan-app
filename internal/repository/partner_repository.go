package repository

import (
	"loan-app/internal/entity"

	"github.com/sirupsen/logrus"
)

type PartnerRepository struct {
	Repository[entity.Partner]
	Log *logrus.Logger
}

func NewPartnerRepository(log *logrus.Logger) *PartnerRepository {
	return &PartnerRepository{
		Log: log,
	}
}
