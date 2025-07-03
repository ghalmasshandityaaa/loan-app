package app

import (
	"loan-app/config"
	"loan-app/internal/handler"
	"loan-app/internal/middleware"
	"loan-app/internal/repository"
	"loan-app/internal/route"
	"loan-app/internal/usecase"
	"loan-app/internal/utils"
	"loan-app/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	App       *fiber.App
	Log       *logrus.Logger
	Config    *config.Config
	DB        *gorm.DB
	Validator *validator.Validator
}

func Bootstrap(config *BootstrapConfig) {
	// init utils
	jwtUtil := utils.NewJwtUtil(config.Config)

	// init repositories
	userRepository := repository.NewUserRepository(config.Log)
	customerLimitRepository := repository.NewCustomerLimitRepository(config.Log)
	partnerRepository := repository.NewPartnerRepository(config.Log)
	assetRepository := repository.NewAssetRepository(config.Log)

	// init use cases
	authUseCase := usecase.NewAuthUseCase(config.DB, config.Log, jwtUtil, userRepository, customerLimitRepository)
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, userRepository, customerLimitRepository)
	partnerUseCase := usecase.NewPartnerUseCase(config.DB, config.Log, partnerRepository)
	assetUseCase := usecase.NewAssetUseCase(config.DB, config.Log, assetRepository, partnerRepository)

	// init handlers
	authHandler := handler.NewAuthHandler(authUseCase, config.Log, config.Config, config.Validator)
	userHandler := handler.NewUserHandler(userUseCase, config.Log, config.Config, config.Validator)
	partnerHandler := handler.NewPartnerHandler(partnerUseCase, config.Log, config.Config, config.Validator)
	assetHandler := handler.NewAssetHandler(assetUseCase, config.Log, config.Config, config.Validator)

	// init middleware
	authMiddleware := middleware.NewAuthMiddleware(userUseCase, jwtUtil)

	// init routes
	appRoute := route.NewRoute(
		config.App,
		config.Log,
		authMiddleware,
		authHandler,
		userHandler,
		partnerHandler,
		assetHandler,
	)

	// setup routes
	appRoute.Setup()
}
