package service

import (
	"ShamanEstBanan-GophKeeper-server/internal/domain/entity"
	"ShamanEstBanan-GophKeeper-server/internal/errs"
	"ShamanEstBanan-GophKeeper-server/internal/utils/authtoken"
	"context"

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
	token, err := authtoken.GenerateToken(userID)
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
