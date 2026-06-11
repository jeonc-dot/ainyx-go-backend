package repository

import (
	"context"

	"github.com/jeon-prince/ainyx-go-backend/db/sqlc"
)

// UserRepository wraps the SQLC generated queries to provide a clean interface.
// In enterprise apps, we use interfaces here to allow for mock testing.
type UserRepository interface {
	CreateUser(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error)
	GetUser(ctx context.Context, id int32) (sqlc.User, error)
	UpdateUser(ctx context.Context, arg sqlc.UpdateUserParams) (sqlc.User, error)
	DeleteUser(ctx context.Context, id int32) error
	ListUsers(ctx context.Context, arg sqlc.ListUsersParams) ([]sqlc.User, error)
}

type userRepo struct {
	queries *sqlc.Queries
}

// NewUserRepository creates a new UserRepository.
func NewUserRepository(queries *sqlc.Queries) UserRepository {
	return &userRepo{
		queries: queries,
	}
}

func (r *userRepo) CreateUser(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error) {
	return r.queries.CreateUser(ctx, arg)
}

func (r *userRepo) GetUser(ctx context.Context, id int32) (sqlc.User, error) {
	return r.queries.GetUser(ctx, id)
}

func (r *userRepo) UpdateUser(ctx context.Context, arg sqlc.UpdateUserParams) (sqlc.User, error) {
	return r.queries.UpdateUser(ctx, arg)
}

func (r *userRepo) DeleteUser(ctx context.Context, id int32) error {
	return r.queries.DeleteUser(ctx, id)
}

func (r *userRepo) ListUsers(ctx context.Context, arg sqlc.ListUsersParams) ([]sqlc.User, error) {
	return r.queries.ListUsers(ctx, arg)
}
