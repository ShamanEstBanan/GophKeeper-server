package service

import (
	"ShamanEstBanan-GophKeeper-server/internal/domain/entity"
	"context"
	"go.uber.org/zap"
)

type storage interface {
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
	storage storage
}

func New(lg *zap.Logger, storage storage) *service {
	return &service{
		lg:      lg,
		storage: storage,
	}
}
