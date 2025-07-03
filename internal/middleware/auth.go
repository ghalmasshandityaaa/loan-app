package middleware

import (
	"loan-app/internal/model"
	"loan-app/internal/usecase"
	"loan-app/internal/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func NewAuthMiddleware(userUseCase *usecase.UserUseCase, util *utils.JwtUtil) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := &model.VerifyAccountRequest{
			Token: ctx.Get("Authorization", "NOT_FOUND"),
		}

		token := request.Token
		if token == "NOT_FOUND" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(model.WebResponse[any]{
				Ok:     false,
				Errors: "auth/unauthorized",
			})
		} else if strings.HasPrefix(token, "Bearer ") {
			token = strings.TrimPrefix(token, "Bearer ")
		}

		claim, err := util.ValidateToken(token)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(model.WebResponse[any]{
				Ok:     false,
				Errors: "auth/unauthorized",
			})
		}

		user, err := userUseCase.GetById(ctx.UserContext(), claim.ID)
		if err != nil || user == nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(model.WebResponse[any]{
				Ok:     false,
				Errors: "auth/unauthorized",
			})
		}

		ctx.Locals("auth", &model.Auth{
			ID:      user.ID,
			IsAdmin: user.IsAdmin,
		})
		return ctx.Next()
	}
}

func GetAuth(ctx *fiber.Ctx) *model.Auth {
	return ctx.Locals("auth").(*model.Auth)
}
