package config

import (
	"errors"
	"net/http"
)

var (

	ErrInvalidSMTPPort     = errors.New("invalid SMTP_PORT: must be an integer")
	ErrMissingEnvVariables = errors.New("missing required environment variables")

	ErrInvalidWebhookSignature = errors.New("invalid webhook signature")
	ErrInvalidPaymentData      = errors.New("invalid payment payload structure")
	ErrTransactionNotFound     = errors.New("transaction not found")
	ErrPaymentAlreadyHandled   = errors.New("payment has already been processed")
	ErrPaymentFailed           = errors.New("payment failed")
	ErrUnsupportedCurrency     = errors.New("unsupported currency")
	ErrWebhookTimeout          = errors.New("webhook processing timed out")

	
	ErrInvalidEmailFormat   = errors.New("invalid email format")
	ErrEmailAlreadyExists   = errors.New("email already registered")
	ErrUserNotFound         = errors.New("user not found")
	ErrIncorrectCredentials = errors.New("incorrect email or password")
	ErrTokenExpired         = errors.New("token has expired or is invalid")
	ErrJWTSigningFailed     = errors.New("failed to sign JWT")
	ErrUnauthorizedAccess   = errors.New("unauthorized access")

	ErrInternalServer = errors.New("internal server error")
	ErrBadRequest     = errors.New("bad request")

	ErrUnsupportedWebhookEvent = errors.New("unsupported webhook event")
)

func GetStatusCode(err error) int {
	switch err {
	case nil:
		return http.StatusOK

	case ErrInvalidSMTPPort, ErrMissingEnvVariables, ErrInvalidEmailFormat, ErrInvalidPaymentData, ErrBadRequest:
		return http.StatusBadRequest

	case ErrUnauthorizedAccess, ErrIncorrectCredentials, ErrTokenExpired, ErrJWTSigningFailed, ErrInvalidWebhookSignature:
		return http.StatusUnauthorized

	case ErrEmailAlreadyExists:
		return http.StatusConflict

	case ErrUserNotFound, ErrTransactionNotFound:
		return http.StatusNotFound

	case ErrPaymentAlreadyHandled, ErrUnsupportedCurrency, ErrPaymentFailed, ErrWebhookTimeout:
		return http.StatusUnprocessableEntity

	default:
		return http.StatusInternalServerError
	}
}
