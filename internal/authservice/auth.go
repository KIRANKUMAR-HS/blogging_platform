package authservice

import (
	"errors"
	"time"

	db "github.com/KIRANKUMAR-HS/blogging_platform/internal/psql"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Db        *db.PsqlClient // User data access
	SecretKey string         // Secret key for JWT
}

func NewAuthService(Db *db.PsqlClient, secretkey string) (*AuthService, error) {
	return &AuthService{
		Db:        Db,
		SecretKey: secretkey,
	}, nil
}

// Authenticate the user and return a JWT token
func (a *AuthService) Authenticate(username, password string) (string, error) {
	user, err := a.Db.FindByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// Compare the password with the stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password_hash), []byte(password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	// Create a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{

		"username": user.Name,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	tokenString, err := token.SignedString([]byte(a.SecretKey))
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return tokenString, nil
}
