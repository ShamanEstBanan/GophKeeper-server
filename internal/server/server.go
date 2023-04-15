package server

import (
	"ShamanEstBanan-GophKeeper-server/internal/domain/entity"
	pb "ShamanEstBanan-GophKeeper-server/internal/proto"
	"context"
)

type service interface {
	SignUp(context.Context, entity.User) error
	LogIn(context.Context, entity.User) (string, error)
	GetAllRecords(context.Context, entity.UserID) (*[]entity.RecordInfo, error)
	GetRecordsByType(ctx context.Context, userID entity.UserID, datatype entity.DataType) (*[]entity.RecordInfo, error)
	CreateRecord(context.Context, entity.Record) (*entity.Record, error)
	GetRecord(context.Context, entity.RecordID, entity.UserID) (*entity.Record, error)
	UpdateRecord(context.Context, entity.Record) (*entity.Record, error)
	DeleteRecord(context.Context, entity.RecordID, entity.UserID) error
}
type KeeperService struct {
	Service service
	pb.UnimplementedKeeperServiceServer
}
