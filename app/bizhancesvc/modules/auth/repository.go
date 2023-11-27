package auth

import (
	"bizhancesvc/shared"
	"context"
	"log"
	"time"

	"github.com/gofrs/uuid"
	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/jackc/pgx/pgtype"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type userDB struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Password  string
	Status    string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

const (
	USERSTATUS_ACTIVE            string = "active"
	USERSTATUS_INACTIVE          string = "inactive"
	USERSTATUS_NEED_VERIFICATION string = "need_verification"
)

func getUserByEmail(ctx context.Context, email string) (user userDB, err error) {
	query := "SELECT * FROM users WHERE email=$1"
	row := AuthModuleInstance.dbPool.QueryRow(ctx, query, email)

	if err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Status, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err != pgx.ErrNoRows {
			log.Print(err)
		}
		return userDB{}, err
	}
	return
}
func validateUserPassword(inputPassword, userPassword string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(inputPassword))
}

func generateJWT(userID string, secret string, dayToExpire int) (string, error) {
	claims := shared.UserJWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{time.Now().AddDate(0, 0, dayToExpire)},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
