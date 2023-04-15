package storage

import (
	"context"
	"fmt"
	"time"

	"ShamanEstBanan-GophKeeper-server/internal/domain/entity"
	"ShamanEstBanan-GophKeeper-server/internal/errs"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

const defaultTimeout time.Duration = time.Millisecond * 500

func (s *storage) CreateUser(ctx context.Context, user entity.User) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	query := `INSERT INTO users (login, password) VALUES ($1, $2)`
	_, err := s.db.Exec(ctx, query, user.Login, user.Password)
	switch e := err.(type) {
	case *pgconn.PgError:
		switch e.Code {
		case pgerrcode.UniqueViolation:
			return errs.ErrLoginAlreadyExist
		}
	default:
		return err
	}
	return nil
}

func (s *storage) AuthenticateUser(ctx context.Context, user entity.User) (userID entity.UserID, err error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	var exist bool
	//Проверка, есть ли записи с указанными логином и паролем
	query := fmt.Sprintf(
		"SELECT password = '%s' AS pswmatch FROM users WHERE login = '%s'",
		user.Password, user.Login)

	err = s.db.QueryRow(ctx, query).Scan(&exist)
	if err != nil {
		return "", err
	}
	if !exist {
		return "", errs.ErrInvalidLoginOrPassword
	}
	query = fmt.Sprintf("SELECT id FROM users WHERE login = '%s'", user.Login)
	err = s.db.QueryRow(ctx, query).Scan(&userID)
	if err != nil {
		return "", err
	}
	return userID, err
}
