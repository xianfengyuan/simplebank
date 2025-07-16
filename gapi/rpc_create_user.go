package gapi

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	db "github.com/xianfengyuan/simplebank/db/sqlc"
	"github.com/xianfengyuan/simplebank/pb"
	"github.com/xianfengyuan/simplebank/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPassword, _ := util.HashedPassword(req.GetPassword())
	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
		FullName:       req.GetFullName(),
		Email:          req.GetEmail(),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pgerr, ok := err.(*pgconn.PgError); ok {
			switch pgerr.Code {
			case "23503", "23505":
				return nil, status.Errorf(codes.InvalidArgument, "Failed to create user: %s", pgerr)
			}
		}
		return nil, status.Errorf(codes.Internal, "Failed to create user: %s", err)
	}
	rsp := &pb.CreateUserResponse{
		User: convertUser(user),
	}

	return rsp, nil
}
