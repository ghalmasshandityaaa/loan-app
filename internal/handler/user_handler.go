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

type UserHandler struct {
	Log       *logrus.Logger
	Config    *config.Config
	UseCase   *usecase.UserUseCase
	Validator *validator.Validator
}

func NewUserHandler(
	useCase *usecase.UserUseCase,
	log *logrus.Logger,
	config *config.Config,
	validator *validator.Validator,
) *UserHandler {
	return &UserHandler{
		Log:       log,
		UseCase:   useCase,
		Config:    config,
		Validator: validator,
	}
}

// FindSelf retrieves the authenticated user's information.
// @Summary Find Self User
// @Description Retrieve the authenticated user's information.
// @Tags User
// @Accept json
// @Produce json
// @Security bearer
// @Router /user/me [get]
// @Success 200 {object} model.FindSelfResponseWrapper
func (h *UserHandler) FindSelf(ctx *fiber.Ctx) error {
	method := "UserHandler.FindSelf"
	log := h.Log.WithField("method", method)
	log.Trace("[BEGIN]")

	auth := middleware.GetAuth(ctx)
	requestCtx := ctx.UserContext()
	user, err := h.UseCase.GetById(requestCtx, auth.ID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Ok:     false,
			Errors: err.Error(),
		})
	}

	log.Trace("[END]")

	return ctx.JSON(model.WebResponse[*entity.User]{
		Ok:   true,
		Data: user,
	})
}

// FindLimits retrieves the customer limits for the authenticated user.
// @Summary Find Customer Limits
// @Description Retrieve the customer limits for the authenticated user.
// @Tags User
// @Accept json
// @Produce json
// @Security bearer
// @Router /user/limit [get]
// @Success 200 {object} model.FindCustomerLimitWrapper
func (h *UserHandler) FindLimits(ctx *fiber.Ctx) error {
	method := "UserHandler.FindLimits"
	log := h.Log.WithField("method", method)
	log.Trace("[BEGIN]")

	auth := middleware.GetAuth(ctx)
	requestCtx := ctx.UserContext()
	limits, err := h.UseCase.FindLimits(requestCtx, auth.ID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Ok:     false,
			Errors: err.Error(),
		})
	}

	log.Trace("[END]")

	return ctx.JSON(model.WebResponseWithData[[]entity.CustomerLimit]{
		Ok:   true,
		Data: limits,
	})
}
