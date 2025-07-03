package handler

import (
	"loan-app/config"
	"loan-app/internal/entity"
	"loan-app/internal/middleware"
	"loan-app/internal/model"
	"loan-app/internal/usecase"
	"loan-app/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AssetHandler struct {
	Log       *logrus.Logger
	Config    *config.Config
	UseCase   *usecase.AssetUseCase
	Validator *validator.Validator
}

func NewAssetHandler(
	useCase *usecase.AssetUseCase,
	log *logrus.Logger,
	config *config.Config,
	validator *validator.Validator,
) *AssetHandler {
	return &AssetHandler{
		Log:       log,
		UseCase:   useCase,
		Config:    config,
		Validator: validator,
	}
}

func (h *AssetHandler) Create(ctx *fiber.Ctx) error {
	method := "AssetHandler.Create"
	log := h.Log.WithField("method", method)
	log.Trace("[BEGIN]")

	auth := middleware.GetAuth(ctx)
	requestCtx := ctx.UserContext()

	request := new(model.CreateAssetRequest)
	if err := ctx.BodyParser(request); err != nil {
		log.Error("failed parse body: ", err.Error())
		return fiber.ErrBadRequest
	}

	errValidation := h.Validator.ValidateStruct(request)
	if errValidation != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Ok:     false,
			Errors: errValidation,
		})
	}

	if err := h.UseCase.Create(requestCtx, auth.ID, request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Ok:     false,
			Errors: err.Error(),
		})
	}

	log.Trace("[END]")

	return ctx.JSON(model.WebResponse[*entity.Asset]{
		Ok: true,
	})
}

func (h *AssetHandler) Lists(ctx *fiber.Ctx) error {
	method := "AssetHandler.Lists"
	log := h.Log.WithField("method", method)
	log.Trace("[BEGIN]")

	requestCtx := ctx.UserContext()
	assets, err := h.UseCase.Lists(requestCtx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Ok:     false,
			Errors: err.Error(),
		})
	}

	log.Trace("[END]")

	return ctx.JSON(model.WebResponseWithData[[]entity.Asset]{
		Ok:   true,
		Data: assets,
	})
}
