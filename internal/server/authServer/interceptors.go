package authServer

import (
	"ShamanEstBanan-GophKeeper-server/internal/errs"
	"ShamanEstBanan-GophKeeper-server/internal/utils/authtoken"
	"context"
	"errors"
	"log"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var tokenFromReq string
	methodInfo := info.FullMethod
	md, ok := metadata.FromIncomingContext(ctx)
	if strings.Contains(methodInfo, "LogIn") || strings.Contains(methodInfo, "SignUp") {
		return handler(ctx, req)
	}
	if ok {
		values := md.Get("jwt-token")
		if len(values) > 0 {
			tokenFromReq = values[0]
		}
	}

	if len(tokenFromReq) == 0 {
		return nil, status.Error(codes.Unauthenticated, "missing token")
	}
	userID, err := authtoken.ParseToken(tokenFromReq)
	if errors.Is(err, errs.ErrInvalidAccessToken) {
		log.Printf("invalid token: %s", tokenFromReq)
		return nil, status.Errorf(codes.Unauthenticated, "Неверный auth-token")
	}
	if err != nil {
		log.Printf("invalid token: %s", tokenFromReq)
		return nil, status.Errorf(codes.Unauthenticated, "Неверный auth-token")
	}

	newContext := context.WithValue(ctx, "userID", userID)

	return handler(newContext, req)
}
