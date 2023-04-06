package storage

import (
	"ShamanEstBanan-GophKeeper-server/internal/domain/entity"
	"context"
	"time"
)

const defaultTimeout time.Duration = time.Millisecond * 500

func (s *storage) CreateUser(ctx context.Context, user *entity.User) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	query := `INSERT INTO users (login, password) VALUES ($1, $2)`
	_, err := s.db.Query(ctx, query, user.Login, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (s *storage) AuthenticateUser(ctx context.Context, user entity.User) (*entity.User, error) {
	return nil, nil
}
