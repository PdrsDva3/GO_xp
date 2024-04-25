package user

import (
	"GO_xp/internal/entities"
	"GO_xp/internal/repository"
	"context"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
	db *sqlx.DB
}

func InitUserRepo(db *sqlx.DB) repository.User {
	return UserRepo{
		db: db,
	}
}

func (usr UserRepo) Create(ctx context.Context, user entities.UserCreate) (int, error) {
	var userID int
	transaction, err := usr.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	row := transaction.QueryRowContext(ctx, `INSERT INTO users (login, username, hashed_password) VALUES ($1, $2, $3) RETURNING id;`,
		user.Login, user.Name, hashedPassword)

	err = row.Scan(&userID)
	if err != nil {
		return 0, err
	}

	if err = transaction.Commit(); err != nil {
		return 0, err
	}

	return userID, nil
}

func (usr UserRepo) Get(ctx context.Context, userID int) (*entities.User, error) {
	var user entities.User

	row := usr.db.QueryRowContext(ctx, `SELECT id, login, users.username FROM users WHERE users.id = $1;`, userID)

	err := row.Scan(&user.UserID, &user.Login, &user.Name)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (usr UserRepo) GetPassword(ctx context.Context, login string) (int, string, error) {
	var (
		userID          int
		hashed_password string
	)

	row := usr.db.QueryRowContext(ctx, `SELECT hashed_password, id FROM users WHERE users.login=$1;`, login)

	err := row.Scan(&hashed_password, &userID)
	if err != nil {
		return 0, "", err
	}

	return userID, hashed_password, nil
}

func (usr UserRepo) UpdatePassword(ctx context.Context, userID int, newPassword string) error {
	transaction, err := usr.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), 10)
	result, err := transaction.ExecContext(ctx, `UPDATE users SET hashed_password = $2 WHERE users.id = $1;`, userID, hashedPassword)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if count != 1 {
		return err //todo кастомная ошибка!
	}

	if err = transaction.Commit(); err != nil {
		return err
	}

	return nil
}

func (usr UserRepo) Delete(ctx context.Context, userID int) error {
	transaction, err := usr.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	result, err := transaction.ExecContext(ctx, `DELETE FROM users WHERE users.id=$1;`, userID)
	if err != nil {
		return err
	}
	//todo разве здесь не нужно что-то сделать? >>
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if count != 1 {
		return err //todo кастомная ошибка!
	}
	// <<
	if err = transaction.Commit(); err != nil {
		return err
	}

	return nil
}
