package authServer

import (
	"ShamanEstBanan-GophKeeper-server/internal/domain/entity"
	"ShamanEstBanan-GophKeeper-server/internal/errs"
	pb "ShamanEstBanan-GophKeeper-server/internal/proto"
	mainService "ShamanEstBanan-GophKeeper-server/internal/server/authServer/mock"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestAuthServer_SignUp(t *testing.T) {

	type args struct {
		ctx context.Context
		in  *pb.SignUpRequest
	}
	tests := []struct {
		name     string
		args     args
		mockCall func(
			args args,
			service *mainService.MockService,
		)
		want    *pb.SignUpResponse
		wantErr bool
	}{
		{
			name: "success signUp",
			args: args{
				ctx: context.Background(),
				in: &pb.SignUpRequest{
					Login:    "test login",
					Password: "test password",
				},
			},
			mockCall: func(
				args args,
				service *mainService.MockService,
			) {
				user := entity.User{
					UserID:   "",
					Login:    "test login",
					Password: "test password",
				}
				service.
					EXPECT().
					SignUp(args.ctx, user).
					Return(nil).
					Times(1)

			},
			want:    &pb.SignUpResponse{},
			wantErr: false,
		},
		{
			name: "error: empty login",
			args: args{
				ctx: context.Background(),
				in: &pb.SignUpRequest{
					Login:    "",
					Password: "test password",
				},
			},
			mockCall: func(
				args args,
				service *mainService.MockService,
			) {
				user := entity.User{
					UserID:   "",
					Login:    "",
					Password: "test password",
				}
				err := errs.ErrLoginIsEmpty
				service.
					EXPECT().
					SignUp(args.ctx, user).
					Return(err).
					Times(1)

			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error: empty password",
			args: args{
				ctx: context.Background(),
				in: &pb.SignUpRequest{
					Login:    "test login",
					Password: "",
				},
			},
			mockCall: func(
				args args,
				service *mainService.MockService,
			) {
				user := entity.User{
					UserID:   "",
					Login:    "test login",
					Password: "",
				}
				err := errs.ErrPasswordIsEmpty
				service.
					EXPECT().
					SignUp(args.ctx, user).
					Return(err).
					Times(1)

			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error: empty login and password",
			args: args{
				ctx: context.Background(),
				in: &pb.SignUpRequest{
					Login:    "",
					Password: "",
				},
			},
			mockCall: func(
				args args,
				service *mainService.MockService,
			) {
				user := entity.User{
					UserID:   "",
					Login:    "",
					Password: "",
				}
				err := errs.ErrInvalidLoginOrPassword
				service.
					EXPECT().
					SignUp(args.ctx, user).
					Return(err).
					Times(1)

			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error: some service error",
			args: args{
				ctx: context.Background(),
				in: &pb.SignUpRequest{
					Login:    "test login",
					Password: "test password",
				},
			},
			mockCall: func(
				args args,
				service *mainService.MockService,
			) {
				user := entity.User{
					UserID:   "",
					Login:    "test login",
					Password: "test password",
				}
				err := errors.New("some error")
				service.
					EXPECT().
					SignUp(args.ctx, user).
					Return(err).
					Times(1)

			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error: login already exist",
			args: args{
				ctx: context.Background(),
				in: &pb.SignUpRequest{
					Login:    "test login",
					Password: "test password",
				},
			},
			mockCall: func(
				args args,
				service *mainService.MockService,
			) {
				user := entity.User{
					UserID:   "",
					Login:    "test login",
					Password: "test password",
				}
				err := errs.ErrLoginAlreadyExist
				service.
					EXPECT().
					SignUp(args.ctx, user).
					Return(err).
					Times(1)

			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			service := mainService.NewMockService(ctrl)
			tt.mockCall(tt.args, service)

			k := &AuthServer{
				Service:                        service,
				UnimplementedAuthServiceServer: pb.UnimplementedAuthServiceServer{},
			}

			got, err := k.SignUp(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignUp() got = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestAuthServer_LogIn(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pb.LogInRequest
	}
	tests := []struct {
		name     string
		args     args
		mockCall func(
			args args,
			service *mainService.MockService,
		)
		want    *pb.LogInResponse
		wantErr bool
	}{
		{
			name: "success logIn",
			args: args{
				ctx: context.Background(),
				in: &pb.LogInRequest{
					Login:    "test user",
					Password: "test password",
				},
			},
			mockCall: func(
				args args,
				service *mainService.MockService) {
				user := entity.User{
					UserID:   "",
					Login:    "test user",
					Password: "test password",
				}
				token := "test token"
				service.
					EXPECT().
					LogIn(args.ctx, user).
					Return(token, nil).
					Times(1)
			},

			want:    &pb.LogInResponse{},
			wantErr: false,
		},
		{
			name: "error: empty login",
			args: args{
				ctx: context.Background(),
				in: &pb.LogInRequest{
					Login:    "",
					Password: "test password",
				},
			},
			mockCall: func(
				args args,
				service *mainService.MockService) {
				user := entity.User{
					UserID:   "",
					Login:    "",
					Password: "test password",
				}
				err := errs.ErrLoginIsEmpty
				service.
					EXPECT().
					LogIn(args.ctx, user).
					Return("", err).
					Times(1)
			},

			want:    nil,
			wantErr: true,
		},
		{
			name: "error: empty login",
			args: args{
				ctx: context.Background(),
				in: &pb.LogInRequest{
					Login:    "test user",
					Password: "",
				},
			},
			mockCall: func(
				args args,
				service *mainService.MockService) {
				user := entity.User{
					UserID:   "",
					Login:    "test user",
					Password: "",
				}
				err := errs.ErrPasswordIsEmpty
				service.
					EXPECT().
					LogIn(args.ctx, user).
					Return("", err).
					Times(1)
			},

			want:    nil,
			wantErr: true,
		},
		{
			name: "error: some error",
			args: args{
				ctx: context.Background(),
				in: &pb.LogInRequest{
					Login:    "test user",
					Password: "test password",
				},
			},
			mockCall: func(
				args args,
				service *mainService.MockService) {
				user := entity.User{
					UserID:   "",
					Login:    "test user",
					Password: "test password",
				}
				err := errors.New("some error")
				service.
					EXPECT().
					LogIn(args.ctx, user).
					Return("", err).
					Times(1)
			},

			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			service := mainService.NewMockService(ctrl)
			tt.mockCall(tt.args, service)

			k := &AuthServer{
				Service:                        service,
				UnimplementedAuthServiceServer: pb.UnimplementedAuthServiceServer{},
			}

			got, err := k.LogIn(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignUp() got = %v, want %v", got, tt.want)
			}

		})
	}
}
