package authtoken

import (
	"ShamanEstBanan-GophKeeper-server/internal/errs"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"go.uber.org/zap"
	"log"
	"time"
)

// TODO убрать в env
const signingKey = "BestCryptoSign"

type Claims struct {
	jwt.StandardClaims
	userID string
}

func GenerateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt: jwt.At(time.Now()),
		},
		userID: userID,
	})
	stringToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		log.Println("Generate auth-token error: ", zap.Error(err))
		return "", err
	}
	return stringToken, nil
}

func ParseToken(userToken string) (userID string, err error) {
	token, err := jwt.ParseWithClaims(userToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.userID, nil
	}
	return "", errs.ErrInvalidAccessToken
}
