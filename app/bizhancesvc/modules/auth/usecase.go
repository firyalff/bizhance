package auth

import (
	"bizhancesvc/shared"
	"context"

	"github.com/jackc/pgx/v5"
)

func getUserByCredentials(ctx context.Context, credentials loginPayload) (user userDB, err error) {
	userRecord, err := getUserByEmail(ctx, credentials.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return userDB{}, shared.ErrNotFound
		}
		return
	}

	err = validateUserPassword(credentials.Password, userRecord.Password)
	if err != nil {
		return userDB{}, shared.ErrUnauthorized
	}

	return userRecord, nil
}

func generateAuthenticationToken(userID string, secretToken string, authExpirationDays int) (token string, err error) {
	return generateJWT(userID, secretToken, authExpirationDays)
}
