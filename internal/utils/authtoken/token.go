package authtoken

import (
	"ShamanEstBanan-GophKeeper-server/internal/errs"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"log"
	"time"
)

// TODO убрать в env
var signingKey = []byte("someSecret")

type UserClaims struct {
	UserID string `json:"UserID"`
	jwt.RegisteredClaims
}

func GenerateToken(userID string) (string, error) {
	claims := UserClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(45 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString(signingKey)
	if err != nil {
		log.Println("Generate auth-token error: ", zap.Error(err))
		return "", err
	}
	return stringToken, nil
}

func ParseToken(userToken string) (uID string, err error) {
	token, err := jwt.ParseWithClaims(userToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims.UserID, nil
	}
	return "", errs.ErrInvalidAccessToken
}
