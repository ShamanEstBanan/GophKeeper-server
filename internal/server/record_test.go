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
