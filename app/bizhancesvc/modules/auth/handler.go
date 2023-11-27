package auth

import (
	"bizhancesvc/shared"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func registerHandlers(router fiber.Router) {
	router.Post("/auth/login", AuthModuleInstance.authLoginHandler)
	router.Post("/auth/register", AuthModuleInstance.authRegisterHandler)
	router.Post("/auth/account-activation", AuthModuleInstance.authAccountActivationHandler)
}

type loginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}

const JWT_EXPIRATION_DAYS = 7

func (authModule AuthModule) authLoginHandler(ctx *fiber.Ctx) error {
	ctx.Context()
	payload := new(loginPayload)
	if err := ctx.BodyParser(payload); err != nil {
		return err
	}

	err := shared.Validator().Struct(payload)
	if err != nil {
		errBody := shared.GenerateErrorResponse("BAD_REQ", shared.ParseValidatorError(err))
		return ctx.Status(400).JSON(errBody)
	}

	userRecord, err := getUserByCredentials(ctx.Context(), *payload)
	if err != nil {
		if err == shared.ErrNotFound || err == shared.ErrUnauthorized {
			return ctx.Status(http.StatusUnauthorized).JSON(shared.GenerateErrorResponse("UNAUTHORIZED", nil))
		} else {
			return ctx.Status(http.StatusInternalServerError).JSON(shared.GenerateErrorResponse("INTERNALERR", nil))
		}
	}

	tokenString, err := generateAuthenticationToken(userRecord.Id.String(), authModule.serverConfig.JWTSecret, JWT_EXPIRATION_DAYS)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(shared.GenerateErrorResponse("INTERNALERR", nil))
	}

	return ctx.Status(200).JSON(map[string]interface{}{
		"token": tokenString,
	})
}

type registerPayload struct {
	FirstName            string `json:"first_name" validate:"required"`
	LastName             string `json:"last_name" validate:"required"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required,min=8,max=64"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,min=8,max=64,eqfield=Password"`
}

func (authModule AuthModule) authRegisterHandler(ctx *fiber.Ctx) error {
	var payload registerPayload
	if err := ctx.BodyParser(payload); err != nil {
		return err
	}

	return ctx.JSON(shared.GenerateErrorResponse("NOT_IMPLEMENTED", map[string]interface{}{
		"detail": "not implemented",
	}))
}

func (authModule AuthModule) authAccountActivationHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(shared.GenerateErrorResponse("NOT_IMPLEMENTED", map[string]interface{}{
		"detail": "not implemented",
	}))
}
