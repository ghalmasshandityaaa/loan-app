package handler

import (
	"loan-app/config"
	"loan-app/internal/model"
	"loan-app/internal/usecase"
	"loan-app/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	Log       *logrus.Logger
	Config    *config.Config
	UseCase   *usecase.AuthUseCase
	Validator *validator.Validator
}

func NewAuthHandler(
	useCase *usecase.AuthUseCase,
	log *logrus.Logger,
	config *config.Config,
	validator *validator.Validator,
) *AuthHandler {
	return &AuthHandler{
		Log:       log,
		UseCase:   useCase,
		Config:    config,
		Validator: validator,
	}
}

// SignIn authenticates a user and returns access and refresh tokens
// @Summary Sign in user
// @Description Authenticate user with nik and password to get access and refresh tokens
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body model.SignInRequest true "Sign in credentials"
// @Router /auth/sign-in [post]
// @Success 200 {object} model.SignInResponseWrapper
func (h *AuthHandler) SignIn(ctx *fiber.Ctx) error {
	method := "AuthHandler.SignIn"
	log := h.Log.WithField("method", method)
	log.Trace("[BEGIN]")

	request := new(model.SignInRequest)
	if err := ctx.BodyParser(request); err != nil {
		return fiber.ErrBadRequest
	}

	errValidation := h.Validator.ValidateStruct(request)
	if errValidation != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[*model.SignInResponse]{
			Ok:     false,
			Errors: errValidation,
		})
	}

	requestCtx := ctx.UserContext()
	accessToken, refreshToken, err := h.UseCase.SignIn(requestCtx, request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Ok:     false,
			Errors: err.Error(),
		})
	}

	log.Trace("[END]")

	return ctx.JSON(model.WebResponse[*model.SignInResponse]{
		Ok: true,
		Data: &model.SignInResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	})
}

// SignUp registers a new user and returns a success response
// @Summary Sign up user
// @Description Register a new user with nik, full_name, and other details
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body model.SignUpRequest true "Sign up credentials"
// @Router /auth/sign-up [post]
// @Success 200 {object} model.SignUpResponseWrapper
func (h *AuthHandler) SignUp(ctx *fiber.Ctx) error {
	method := "AuthHandler.SignUp"
	log := h.Log.WithField("method", method)
	log.Trace("[BEGIN]")

	request := new(model.SignUpRequest)
	if err := ctx.BodyParser(request); err != nil {
		return fiber.ErrBadRequest
	}

	errValidation := h.Validator.ValidateStruct(request)
	if errValidation != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[*model.SignUpResponse]{
			Ok:     false,
			Errors: errValidation,
		})
	}

	requestCtx := ctx.UserContext()
	if err := h.UseCase.SignUp(requestCtx, request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Ok:     false,
			Errors: err.Error(),
		})
	}

	log.Trace("[END]")

	return ctx.JSON(model.WebResponse[*model.SignUpResponse]{
		Ok: true,
	})
}
