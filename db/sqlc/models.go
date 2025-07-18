// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package db

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	ID        int64
	Owner     string
	Balance   int64
	Currency  string
	CreatedAt pgtype.Timestamptz
}

type Entry struct {
	ID        int64
	AccountID int64
	Amount    int64
	CreatedAt pgtype.Timestamptz
}

type Session struct {
	ID           uuid.UUID
	Username     string
	RefreshToken string
	UserAgent    string
	ClientIp     string
	IsBlocked    bool
	ExpiresAt    pgtype.Timestamptz
	CreatedAt    pgtype.Timestamptz
}

type Transfer struct {
	ID            int64
	FromAccountID int64
	ToAccountID   int64
	Amount        int64
	CreatedAt     pgtype.Timestamptz
}

type User struct {
	Username          string
	HashedPassword    string
	FullName          string
	Email             string
	PasswordChangedAt pgtype.Timestamptz
	CreatedAt         pgtype.Timestamptz
}
