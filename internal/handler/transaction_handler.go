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

type TransactionHandler struct {
	Log       *logrus.Logger
	Config    *config.Config
	UseCase   *usecase.TransactionUseCase
	Validator *validator.Validator
}

func NewTransactionHandler(
	useCase *usecase.TransactionUseCase,
	log *logrus.Logger,
	config *config.Config,
	validator *validator.Validator,
) *TransactionHandler {
	return &TransactionHandler{
		Log:       log,
		UseCase:   useCase,
		Config:    config,
		Validator: validator,
	}
}

func (h *TransactionHandler) Create(ctx *fiber.Ctx) error {
	method := "TransactionHandler.Create"
	log := h.Log.WithField("method", method)
	log.Trace("[BEGIN]")

	auth := middleware.GetAuth(ctx)
	requestCtx := ctx.UserContext()

	request := new(model.CreateTransactionRequest)
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

	return ctx.JSON(model.WebResponse[*entity.Transaction]{
		Ok: true,
	})
}

func (h *TransactionHandler) Lists(ctx *fiber.Ctx) error {
	method := "TransactionHandler.Lists"
	log := h.Log.WithField("method", method)
	log.Trace("[BEGIN]")

	requestCtx := ctx.UserContext()
	auth := middleware.GetAuth(ctx)
	transactions, err := h.UseCase.Lists(requestCtx, auth.ID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Ok:     false,
			Errors: err.Error(),
		})
	}

	log.Trace("[END]")

	return ctx.JSON(model.WebResponseWithData[[]entity.Transaction]{
		Ok:   true,
		Data: transactions,
	})
}
