package service

import (
	"ShamanEstBanan-GophKeeper-server/internal/domain/entity"
	"context"
	"errors"
	"go.uber.org/zap"
)

type storage interface {
	CreateUser(context.Context, *entity.User) error
	AuthenticateUser(context.Context, entity.User) (*entity.User, error)
	GetAllRecords(context.Context, entity.UserID) ([]entity.Record, error)
}

type service struct {
	lg      *zap.Logger
	storage storage
}

func New(lg *zap.Logger, storage storage) *service {
	return &service{
		lg:      lg,
		storage: storage,
	}
}

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

func ValidateUser(user *entity.User) error {
	if user.Login == "" {
		return errors.New("empty field Login")
	}
	if user.Password == "" {
		return errors.New("empty field Password")
	}
	return nil
}
func (s *service) LogIn(ctx context.Context, request *entity.LogInRequest) (*entity.LogInResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetAllRecords(ctx context.Context) (*entity.GetAllRecordsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetRecordsByType(ctx context.Context, request *entity.GetRecordsByTypeRequest) (*entity.GetRecordsByTypeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) CreateRecord(ctx context.Context, request *entity.CreateRecordRequest) (*entity.CreateRecordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetRecord(ctx context.Context, request *entity.GetRecordRequest) (*entity.GetRecordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) EditRecord(ctx context.Context, request *entity.EditRecordRequest) (*entity.EditRecordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteRecord(ctx context.Context, request *entity.DeleteRecordRequest) error {
	//TODO implement me
	panic("implement me")
}
