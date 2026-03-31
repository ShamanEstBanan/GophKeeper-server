package service

import (
	"context"

	"ShamanEstBanan-GophKeeper-server/internal/domain/entity"

	"go.uber.org/zap"
)

//go:generate mockgen -destination=./mock/storage.go -package=mock . Storage
type Storage interface {
	CreateUser(context.Context, entity.User) error
	AuthenticateUser(context.Context, entity.User) (entity.UserID, error)
	GetAllRecords(context.Context, entity.UserID) (*[]entity.RecordInfo, error)
	GetRecordsByType(context.Context, entity.UserID, entity.DataType) (*[]entity.RecordInfo, error)
	CreateRecord(context.Context, entity.Record) (*entity.Record, error)
	GetRecord(context.Context, entity.RecordID, entity.UserID) (*entity.Record, error)
	UpdateRecord(context.Context, entity.Record) (*entity.Record, error)
	DeleteRecord(context.Context, entity.RecordID, entity.UserID) error
}

type service struct {
	lg      *zap.Logger
	storage Storage
}

func New(lg *zap.Logger, storage Storage) *service {
	return &service{
		lg:      lg,
		storage: storage,
	}
}
