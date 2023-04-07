package service

import (
	"ShamanEstBanan-GophKeeper-server/internal/errs"
	"context"
	"fmt"
	"time"

	"ShamanEstBanan-GophKeeper-server/internal/domain/entity"

	"github.com/dgrijalva/jwt-go/v4"
	"go.uber.org/zap"
)

func (s *service) SignUp(ctx context.Context, user *entity.User) error {
	err := ValidateUser(user)
	if err != nil {
		s.lg.Error("Validation error:", zap.Error(err))
		return err
	}
	err = s.storage.CreateUser(ctx, user)
	if err != nil {
		s.lg.Error("Creation user error:", zap.Error(err))
		return err
	}
	return nil
}

func (s *service) LogIn(ctx context.Context, user *entity.User) (string, error) {
	err := ValidateUser(user)
	if err != nil {
		s.lg.Error("Validation error:", zap.Error(err))
		return "", err
	}
	userID, err := s.storage.AuthenticateUser(ctx, user)
	if err != nil {
		s.lg.Error("Authenticate user error:", zap.Error(err))
		return "", err
	}
	token, err := s.generateToken(userID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateUser(user *entity.User) error {
	if user.Login == "" {
		return errs.ErrLoginIsEmpty
	}
	if user.Password == "" {
		return errs.ErrPasswordIsEmpty
	}
	return nil
}

// TODO убрать в env
const signingKey = "BestCryptoSign"

type Claims struct {
	jwt.StandardClaims
	userID string
}

func (s *service) generateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt: jwt.At(time.Now()),
		},
		userID: userID,
	})
	stringToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		s.lg.Error("Generate auth-token error: ", zap.Error(err))
		return "", err
	}
	return stringToken, nil
}

func ParseToken(accessToken string, signingKey []byte) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
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
