package gapi

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/xianfengyuan/simplebank/db/sqlc"
	"github.com/xianfengyuan/simplebank/pb"
	"github.com/xianfengyuan/simplebank/util"
	"github.com/xianfengyuan/simplebank/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	violations := validateUpdateUserRequest(req)
	if violations != nil {
		return nil, InvalidArgumentError(violations)
	}

	arg := db.UpdateUserParams{
		Username:       req.GetUsername(),
		FullName: pgtype.Text{
			String: req.GetFullName(),
			Valid:  req.FullName != nil,
		},
		Email: pgtype.Text{
			String: req.GetEmail(),
			Valid:  req.Email != nil,
		},	
	}

	if req.Password != nil {
		hashedPassword, err := util.HashedPassword(req.GetPassword())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
		}

		arg.HashedPassword = pgtype.Text{
			String: hashedPassword,
			Valid: true,
		}
		arg.PasswordChangedAt = pgtype.Timestamptz{
			Time: time.Now(),
			Valid: true,
		}
	}

	user, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		if pgerr, ok := err.(*pgconn.PgError); ok {
			switch pgerr.Code {
			case "23503", "23505":
				return nil, status.Errorf(codes.InvalidArgument, "Failed to Update user: %s", pgerr)
			}
		}
		return nil, status.Errorf(codes.Internal, "Failed to Update user: %s", err)
	}
	rsp := &pb.UpdateUserResponse{
		User: convertUser(user),
	}

	return rsp, nil
}

func validateUpdateUserRequest(req *pb.UpdateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	if req.FullName != nil {
		if err := val.ValidateFullname(req.GetFullName()); err != nil {
			violations = append(violations, fieldViolation("full_name", err))
		}
	}
	if req.Email != nil {
		if err := val.ValidateEmail(req.GetEmail()); err != nil {
			violations = append(violations, fieldViolation("email", err))
		}
	}
	if req.Password != nil {
		if err := val.ValidatePassword(req.GetPassword()); err != nil {
			violations = append(violations, fieldViolation("password", err))
		}
	}

	return violations
}
