package server

import (
	"ShamanEstBanan-GophKeeper-server/internal/domain/entity"
	pb "ShamanEstBanan-GophKeeper-server/internal/proto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k *KeeperService) GetAllRecords(ctx context.Context, in *pb.GetAllRecordsRequest) (*pb.GetAllRecordsResponse, error) {

	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	records, err := k.Service.GetAllRecords(ctx, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	resp := pb.GetAllRecordsResponse{
		Records: buildRecordsResponce(*records),
	}
	return &resp, nil
}

func (k *KeeperService) GetRecordsByType(ctx context.Context, in *pb.GetRecordsByTypeRequest) (*pb.GetRecordsByTypeResponse, error) {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "auth error")
	}
	records, err := k.Service.GetRecordsByType(ctx, userID, in.Type)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	if len(*records) == 0 {
		return &pb.GetRecordsByTypeResponse{}, nil
	}
	respRecords := make([]*pb.GetRecordsByTypeResponse_Record, 0, len(*records))
	for _, v := range *records {
		rec := &pb.GetRecordsByTypeResponse_Record{
			Id:   v.Id,
			Name: v.Name,
			Type: v.Type,
		}

		respRecords = append(respRecords, rec)
	}

	resp := pb.GetRecordsByTypeResponse{
		Records: respRecords,
	}
	return &resp, nil
}

func (k *KeeperService) CreateRecord(ctx context.Context, in *pb.CreateRecordRequest) (*pb.CreateRecordResponse, error) {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "auth error")
	}
	record := entity.Record{
		Id:     "",
		Name:   in.Name,
		Type:   in.Type,
		Data:   in.Data,
		UserID: userID,
	}
	createdRecord, err := k.Service.CreateRecord(ctx, record)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	resp := &pb.CreateRecordResponse{
		Id:   createdRecord.Id,
		Name: createdRecord.Name,
		Type: createdRecord.Type,
		Data: createdRecord.Data,
	}
	return resp, nil
}

func (k *KeeperService) EditRecord(ctx context.Context, in *pb.EditRecordRequest) (*pb.EditRecordResponse, error) {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "auth error")
	}
	record := entity.Record{
		Id:     "",
		Name:   in.Name,
		Type:   in.Type,
		Data:   in.Data,
		UserID: userID,
	}
	updatedRecord, err := k.Service.UpdateRecord(ctx, record)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	resp := &pb.EditRecordResponse{
		Id:   updatedRecord.Id,
		Name: updatedRecord.Name,
		Type: updatedRecord.Type,
		Data: updatedRecord.Data,
	}
	return resp, nil
}

func (k *KeeperService) DeleteRecord(ctx context.Context, in *pb.DeleteRecordRequest) (*pb.DeleteRecordResponse, error) {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "auth error")
	}
	err = k.Service.DeleteRecord(ctx, in.Id, userID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "data not found")
	}
	return nil, nil
}

func buildRecordsResponce(records []entity.RecordInfo) []*pb.GetAllRecordsResponse_Record {
	if len(records) == 0 {
		return nil
	}
	respRecords := make([]*pb.GetAllRecordsResponse_Record, 0, len(records))
	for _, v := range records {
		rec := &pb.GetAllRecordsResponse_Record{
			Id:   v.Id,
			Name: v.Name,
			Type: v.Type,
		}

		respRecords = append(respRecords, rec)
	}
	return respRecords
}

func getUserIDFromContext(ctx context.Context) (userID string, err error) {
	usID := ctx.Value("userID")
	if usID.(string) == "" {
		return "", status.Error(codes.Unauthenticated, "authenticated error")
	}
	return usID.(string), nil
}
