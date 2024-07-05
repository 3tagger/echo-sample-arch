package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

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

func (r *UserRepositoryPostgreSQL) GetAll(ctx context.Context) ([]*user.User, error) {
	res := []*user.User{}
	q := `
		SELECT
			u.user_id,
			u.name,
			u.email,
			u.about
		FROM 
			users u
	`

	rows, err := r.db.QueryContext(ctx, q)
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

func (r *UserRepositoryPostgreSQL) GetOneById(ctx context.Context, userId int64) (*user.User, error) {
	res := &user.User{}
	q := `
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

	err := r.db.QueryRowContext(ctx, q, userId).Scan(&res.Id, &res.Name, &res.Email, &res.About)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return res, nil
}

func (r *UserRepositoryPostgreSQL) InsertOne(ctx context.Context, user user.User) (*user.User, error) {
	res := user
	q := `
		INSERT INTO users
			(name, email, about)
		VALUES
			($1, $2, $3)
		RETURNING
			user_id
	`

	err := r.db.QueryRowContext(ctx, q, user.Name, user.Email, user.About).Scan(&res.Id)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *UserRepositoryPostgreSQL) InsertMany(ctx context.Context, users []user.User) error {
	if len(users) == 0 {
		return nil
	}

	const numOfParamPerItem int = 3
	var (
		sb         strings.Builder
		paramCount int = 1
		paramStr   string
	)
	params := []interface{}{}

	for _, u := range users {
		sb.WriteString(fmt.Sprintf("($%d, $%d, $%d),", paramCount, paramCount+1, paramCount+2))
		params = append(params, u.Name, u.Email, u.About)
		paramCount += numOfParamPerItem
	}

	paramStr = strings.TrimSuffix(sb.String(), ",")

	q := fmt.Sprintf(`
		INSERT INTO users
			(name, email, about)
		VALUES
			%s
	`, paramStr)

	_, err := r.db.ExecContext(ctx, q, params...)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryPostgreSQL) DeleteOne(ctx context.Context, userId int64) error {
	q := `
		DELETE FROM 
			users
		WHERE
			users.user_id = $1
	`

	_, err := r.db.ExecContext(ctx, q, userId)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryPostgreSQL) UpdateOneById(ctx context.Context, userId int64, user user.User) error {
	q := `
		UPDATE 
			users
		SET
			users.name = $2,
			users.email = $3,
			users.about = $4
		WHERE
			user_id = $1
	`

	_, err := r.db.ExecContext(ctx, q, userId, user.Name, user.Email, user.About)
	if err != nil {
		return err
	}

	return nil
}
