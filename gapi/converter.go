package gapi

import (
	db "github.com/xianfengyuan/simplebank/db/sqlc"
	"github.com/xianfengyuan/simplebank/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username: user.Username,
		FullName: user.FullName,
		Email: user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt.Time),
		CreatedAt: timestamppb.New(user.CreatedAt.Time),
	}
}