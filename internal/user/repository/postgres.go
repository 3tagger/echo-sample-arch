package repository

import (
	"context"
	"database/sql"

	"github.com/3tagger/echo-sample-arch/internal/user"
)

type UserRepositoryPostgreSQL struct {
	db *sql.DB
}

func NewUserRepositoryPostgreSQL(db *sql.DB) *UserRepositoryPostgreSQL {
	return &UserRepositoryPostgreSQL{
		db: db,
	}
}

func (r *UserRepositoryPostgreSQL) GetAllUsers(ctx context.Context) ([]*user.User, error) {
	res := []*user.User{}
	sql := `
		SELECT
			u.user_id,
			u.name,
			u.email,
			u.about
		FROM 
			users u
	`

	rows, err := r.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		u := &user.User{}

		err = rows.Scan(&u.Id, &u.Name, &u.Email, &u.About)
		if err != nil {
			return nil, err
		}

		res = append(res, u)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *UserRepositoryPostgreSQL) GetOneUserById(ctx context.Context, userId int64) (*user.User, error) {
	res := &user.User{}
	sql := `
		SELECT
			u.user_id,
			u.name,
			u.email,
			u.about
		FROM 
			users u
		WHERE 
			u.user_id = $1
	`

	err := r.db.QueryRowContext(ctx, sql, userId).Scan(&res.Id, &res.Name, &res.Email, &res.About)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *UserRepositoryPostgreSQL) InsertOneUser(ctx context.Context, user user.User) (*user.User, error) {
	res := user
	sql := `
		INSERT INTO users
			(name, email, about)
		VALUES
			($1, $2, $3)
		RETURNING
			user_id
	`

	err := r.db.QueryRowContext(ctx, sql, user.Name, user.Email, user.About).Scan(&res.Id)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *UserRepositoryPostgreSQL) DeleteOneUser(ctx context.Context, userId int64) error {
	sql := `
		DELETE FROM 
			users
		WHERE
			users.user_id = $1
	`

	_, err := r.db.ExecContext(ctx, sql, userId)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryPostgreSQL) UpdateOneUserById(ctx context.Context, userId int64, user user.User) error {
	sql := `
		UPDATE 
			users
		SET
			users.name = $2,
			users.email = $3,
			users.about = $4
		WHERE
			user_id = $1
	`

	_, err := r.db.ExecContext(ctx, sql, userId, user.Name, user.Email, user.About)
	if err != nil {
		return err
	}

	return nil
}
