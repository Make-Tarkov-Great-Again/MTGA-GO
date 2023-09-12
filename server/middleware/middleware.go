package middleware

import (
	"context"
	"log"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

func CreateJwtMiddleWare(secret string) *jwtmiddleware.JWTMiddleware {
	keyFunc := func(ctx context.Context) (interface{}, error) {
		// Our token must be signed using this data.
		return []byte(secret), nil
	}

	// Set up the validator.
	jwtValidator, err := validator.New(
		keyFunc,
		validator.HS256,
		"https://<issuer-url>/",
		[]string{"<audience>"},
	)
	if err != nil {
		log.Fatalf("failed to set up the validator: %v", err)
	}

	// Set up the middleware.
	middleware := jwtmiddleware.New(jwtValidator.ValidateToken)
	return middleware
}
