package server

import (
	"ShamanEstBanan-GophKeeper-server/internal/domain/entity"
	pb "ShamanEstBanan-GophKeeper-server/internal/proto"
	mainService "ShamanEstBanan-GophKeeper-server/internal/server/mock"
	"context"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestKeeperService_CreateRecord(t *testing.T) {

	type args struct {
		ctx context.Context
		in  *pb.CreateRecordRequest
	}
	tests := []struct {
		name     string
		args     args
		mockCall func(
			args args,
			service *mainService.MockService,
		)
		want    *pb.CreateRecordResponse
		wantErr bool
	}{
		{
			name: "success create record",
			args: args{
				ctx: context.WithValue(context.Background(), "userID", "1"),
				in: &pb.CreateRecordRequest{
					Name: "test Name",
					Type: "test type",
					Data: []byte("password"),
				},
			},
			mockCall: func(
				args args,
				service *mainService.MockService,
			) {
				record := entity.Record{
					Id:     "",
					Name:   "test Name",
					Type:   "test type",
					Data:   []byte("password"),
					UserID: "1",
				}
				createdRec := &entity.Record{
					Id:     "1",
					Name:   "test Name",
					Type:   "test type",
					Data:   []byte("password"),
					UserID: "1",
				}
				service.
					EXPECT().
					CreateRecord(args.ctx, record).
					Return(createdRec, nil).
					Times(1)
			},
			want: &pb.CreateRecordResponse{
				Id:   "1",
				Name: "test Name",
				Type: "test type",
				Data: []byte("password"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			service := mainService.NewMockService(ctrl)
			tt.mockCall(tt.args, service)

			k := &KeeperService{
				Service:                          service,
				UnimplementedKeeperServiceServer: pb.UnimplementedKeeperServiceServer{},
			}

			got, err := k.CreateRecord(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRecord() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//
//func TestKeeperService_DeleteRecord(t *testing.T) {
//	type fields struct {
//		Service                          Service
//		UnimplementedKeeperServiceServer pb.UnimplementedKeeperServiceServer
//	}
//	type args struct {
//		ctx context.Context
//		in  *pb.DeleteRecordRequest
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *pb.DeleteRecordResponse
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			k := &KeeperService{
//				Service:                          tt.fields.Service,
//				UnimplementedKeeperServiceServer: tt.fields.UnimplementedKeeperServiceServer,
//			}
//			got, err := k.DeleteRecord(tt.args.ctx, tt.args.in)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("DeleteRecord() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("DeleteRecord() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestKeeperService_EditRecord(t *testing.T) {
//	type fields struct {
//		Service                          Service
//		UnimplementedKeeperServiceServer pb.UnimplementedKeeperServiceServer
//	}
//	type args struct {
//		ctx context.Context
//		in  *pb.EditRecordRequest
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *pb.EditRecordResponse
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			k := &KeeperService{
//				Service:                          tt.fields.Service,
//				UnimplementedKeeperServiceServer: tt.fields.UnimplementedKeeperServiceServer,
//			}
//			got, err := k.EditRecord(tt.args.ctx, tt.args.in)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("EditRecord() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("EditRecord() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestKeeperService_GetAllRecords(t *testing.T) {
//	type fields struct {
//		Service                          Service
//		UnimplementedKeeperServiceServer pb.UnimplementedKeeperServiceServer
//	}
//	type args struct {
//		ctx context.Context
//		in  *pb.GetAllRecordsRequest
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *pb.GetAllRecordsResponse
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			k := &KeeperService{
//				Service:                          tt.fields.Service,
//				UnimplementedKeeperServiceServer: tt.fields.UnimplementedKeeperServiceServer,
//			}
//			got, err := k.GetAllRecords(tt.args.ctx, tt.args.in)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetAllRecords() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetAllRecords() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestKeeperService_GetRecord(t *testing.T) {
//	type fields struct {
//		Service                          Service
//		UnimplementedKeeperServiceServer pb.UnimplementedKeeperServiceServer
//	}
//	type args struct {
//		ctx context.Context
//		in  *pb.GetRecordRequest
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *pb.GetRecordResponse
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			k := &KeeperService{
//				Service:                          tt.fields.Service,
//				UnimplementedKeeperServiceServer: tt.fields.UnimplementedKeeperServiceServer,
//			}
//			got, err := k.GetRecord(tt.args.ctx, tt.args.in)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetRecord() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetRecord() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestKeeperService_GetRecordsByType(t *testing.T) {
//	type fields struct {
//		Service                          Service
//		UnimplementedKeeperServiceServer pb.UnimplementedKeeperServiceServer
//	}
//	type args struct {
//		ctx context.Context
//		in  *pb.GetRecordsByTypeRequest
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *pb.GetRecordsByTypeResponse
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			k := &KeeperService{
//				Service:                          tt.fields.Service,
//				UnimplementedKeeperServiceServer: tt.fields.UnimplementedKeeperServiceServer,
//			}
//			got, err := k.GetRecordsByType(tt.args.ctx, tt.args.in)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetRecordsByType() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetRecordsByType() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_buildRecordsResponce(t *testing.T) {
//	type args struct {
//		records []entity.RecordInfo
//	}
//	tests := []struct {
//		name string
//		args args
//		want []*pb.GetAllRecordsResponse_Record
//	}{
//		// TODO: Add test cases.
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := buildRecordsResponce(tt.args.records); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("buildRecordsResponce() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func Test_getUserIDFromContext(t *testing.T) {
//	type args struct {
//		ctx context.Context
//	}
//	tests := []struct {
//		name       string
//		args       args
//		wantUserID string
//		wantErr    bool
//	}{
//		// TODO: Add test cases.
//		{},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			gotUserID, err := getUserIDFromContext(tt.args.ctx)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("getUserIDFromContext() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if gotUserID != tt.wantUserID {
//				t.Errorf("getUserIDFromContext() gotUserID = %v, want %v", gotUserID, tt.wantUserID)
//			}
//		})
//	}
//}
