package domain

import "errors"

var (
	ErrAccountNotFound = errors.New("account not found")

	ErrDuplicateAPIKey = errors.New("api key already exists")

	ErrInvoiceNotFound = errors.New("invoice not found")

	ErrUnauthorizedAcess = errors.New("unauthorized access")
)
