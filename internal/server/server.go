package server

import (
	"ShamanEstBanan-GophKeeper-server/internal/domain/entity"
	pb "ShamanEstBanan-GophKeeper-server/internal/proto"
	"context"
)

type service interface {
	SignUp(context.Context, *entity.User) error
	LogIn(context.Context, *entity.User) (string, error)
	GetAllRecords(context.Context) (*entity.GetAllRecordsResponse, error)
	GetRecordsByType(context.Context, *entity.GetRecordsByTypeRequest) (*entity.GetRecordsByTypeResponse, error)
	CreateRecord(context.Context, *entity.CreateRecordRequest) (*entity.CreateRecordResponse, error)
	GetRecord(context.Context, *entity.GetRecordRequest) (*entity.GetRecordResponse, error)
	EditRecord(context.Context, *entity.EditRecordRequest) (*entity.EditRecordResponse, error)
	DeleteRecord(context.Context, *entity.DeleteRecordRequest) error
}
type KeeperService struct {
	Service service
	pb.UnimplementedKeeperServiceServer
}

func (k *KeeperService) GetAllRecords(context.Context, *pb.GetAllRecordsRequest) (*pb.GetAllRecordsResponse, error) {

	return nil, nil
}

func (k *KeeperService) GetRecordsByType(context.Context, *pb.GetRecordsByTypeRequest) (*pb.GetRecordsByTypeResponse, error) {

	return nil, nil
}

func (k *KeeperService) CreateRecord(context.Context, *pb.CreateRecordRequest) (*pb.CreateRecordResponse, error) {

	return nil, nil
}

func (k *KeeperService) GetRecord(context.Context, *pb.GetRecordRequest) (*pb.GetRecordResponse, error) {

	return nil, nil
}

func (k *KeeperService) EditRecord(context.Context, *pb.EditRecordRequest) (*pb.EditRecordResponse, error) {

	return nil, nil
}

func (k *KeeperService) DeleteRecord(context.Context, *pb.DeleteRecordRequest) (*pb.DeleteRecordResponse, error) {

	return nil, nil
}
