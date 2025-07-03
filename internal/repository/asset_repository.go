package repository

import (
	"loan-app/internal/entity"

	"github.com/sirupsen/logrus"
)

type AssetRepository struct {
	Repository[entity.Asset]
	Log *logrus.Logger
}

func NewAssetRepository(log *logrus.Logger) *AssetRepository {
	return &AssetRepository{
		Log: log,
	}
}
