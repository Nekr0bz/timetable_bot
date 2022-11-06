package repo

import (
	"context"
	"github.com/Nekr0bz/timetable_bot/internal/entity"
	"github.com/Nekr0bz/timetable_bot/pkg/postgres"
	"github.com/jackc/pgx/v4"
)

type UserRepo interface {
	GetUser(ctx context.Context, id int64) (user entity.User, err error)
	CreateUser(ctx context.Context, user *entity.User) (err error)
	GetOrCreateUser(ctx context.Context, user *entity.User) (created bool, err error)
}

type userRepo struct {
	db *postgres.Postgres
}

const userTableName = "public.user"

func NewUserRepo(db *postgres.Postgres) UserRepo {
	return &userRepo{db: db}
}

func (ur *userRepo) GetUser(ctx context.Context, id int64) (user entity.User, err error) {
	q, args, err := ur.db.Builder.Select("*").From(userTableName).Where("id = ?", id).ToSql()
	if err != nil {
		return user, err
	}

	row := ur.db.Conn.QueryRow(ctx, q, args...)
	err = row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.LanguageCode)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *userRepo) CreateUser(ctx context.Context, user *entity.User) (err error) {
	sql, args, err := ur.db.Builder.Insert(userTableName).
		Columns("id", "first_name", "last_name", "username", "language_code", "created_at").
		Values(user.ID, user.FirstName, user.LastName, user.Username, user.LanguageCode, user.CreatedAt).
		ToSql()
	if err != nil {
		return err
	}

	_, err = ur.db.Conn.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

// GetOrCreateUser returns true if user was created
// and false if user already exists
func (ur *userRepo) GetOrCreateUser(ctx context.Context, user *entity.User) (created bool, err error) {
	_, err = ur.GetUser(ctx, user.ID)

	if err == pgx.ErrNoRows {
		err = ur.CreateUser(ctx, user)
		if err == nil {
			created = true
		}
	}
	return created, err
}
