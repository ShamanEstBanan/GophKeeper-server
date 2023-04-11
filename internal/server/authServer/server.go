package authServer

import (
	"ShamanEstBanan-GophKeeper-server/internal/domain/entity"
	"ShamanEstBanan-GophKeeper-server/internal/errs"
	pb "ShamanEstBanan-GophKeeper-server/internal/proto"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
)

type service interface {
	SignUp(context.Context, *entity.User) error
	LogIn(context.Context, *entity.User) (string, error)
}
type AuthServer struct {
	Service service
	pb.UnimplementedAuthServiceServer
}

func (k *AuthServer) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	var resp pb.SignUpResponse
	user := &entity.User{
		UserID:   "",
		Login:    in.Login,
		Password: in.Password,
	}
	err := k.Service.SignUp(ctx, user)
	if err != nil {
		if errors.Is(err, errs.ErrLoginIsEmpty) ||
			errors.Is(err, errs.ErrPasswordIsEmpty) ||
			errors.Is(err, errs.ErrInvalidLoginOrPassword) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		} else if errors.Is(err, errs.ErrLoginAlreadyExist) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		} else {
			return nil, status.Error(codes.Internal, "internal problem")
		}
	}
	if err != nil {
		log.Println(fmt.Errorf("invalid request"))
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	return &resp, nil
}

func (k *AuthServer) LogIn(ctx context.Context, in *pb.LogInRequest) (*pb.LogInResponse, error) {
	var resp pb.LogInResponse

	mdReq, ok := metadata.FromIncomingContext(ctx)
	if ok {
		values := mdReq.Get("authorization")
		if len(values) > 0 {
			return nil, status.Error(codes.AlreadyExists, "already authenticated token")
		}
	}

	user := &entity.User{
		UserID:   "",
		Login:    in.Login,
		Password: in.Password,
	}

	token, err := k.Service.LogIn(ctx, user)
	if err != nil {
		if errors.Is(err, errs.ErrLoginIsEmpty) ||
			errors.Is(err, errs.ErrPasswordIsEmpty) ||
			errors.Is(err, errs.ErrInvalidLoginOrPassword) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		} else {
			return nil, status.Error(codes.Internal, "internal problem")
		}
	}

	header := metadata.Pairs("jwt-token", token) //в key нельзя ставить пробел
	err = grpc.SendHeader(ctx, header)

	return &resp, nil
}
