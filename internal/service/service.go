package service

import (
	"ShamanEstBanan-GophKeeper-server/internal/domain/entity"
	"context"
	"go.uber.org/zap"
)

type storage interface {
	CreateUser(context.Context, *entity.User) error
	AuthenticateUser(context.Context, *entity.User) (entity.UserID, error)
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
