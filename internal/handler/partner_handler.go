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

type PartnerHandler struct {
	Log       *logrus.Logger
	Config    *config.Config
	UseCase   *usecase.PartnerUseCase
	Validator *validator.Validator
}

func NewPartnerHandler(
	useCase *usecase.PartnerUseCase,
	log *logrus.Logger,
	config *config.Config,
	validator *validator.Validator,
) *PartnerHandler {
	return &PartnerHandler{
		Log:       log,
		UseCase:   useCase,
		Config:    config,
		Validator: validator,
	}
}

func (h *PartnerHandler) Create(ctx *fiber.Ctx) error {
	method := "PartnerHandler.Create"
	log := h.Log.WithField("method", method)
	log.Trace("[BEGIN]")

	auth := middleware.GetAuth(ctx)
	requestCtx := ctx.UserContext()

	request := new(model.CreatePartnerRequest)
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

	return ctx.JSON(model.WebResponse[*entity.Partner]{
		Ok: true,
	})
}

func (h *PartnerHandler) Lists(ctx *fiber.Ctx) error {
	method := "PartnerHandler.Lists"
	log := h.Log.WithField("method", method)
	log.Trace("[BEGIN]")

	requestCtx := ctx.UserContext()
	partners, err := h.UseCase.Lists(requestCtx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Ok:     false,
			Errors: err.Error(),
		})
	}

	log.Trace("[END]")

	return ctx.JSON(model.WebResponseWithData[[]entity.Partner]{
		Ok:   true,
		Data: partners,
	})
}
