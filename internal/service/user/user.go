package user

import (
	"GO_xp/internal/entities"
	"GO_xp/internal/repository"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type UserServ struct {
	UserRepo repository.UserRepo
}

func InitUserServ(userRepo repository.UserRepo) *UserServ {
	return &UserServ{UserRepo: userRepo}
}

func (usrs UserServ) Create(ctx context.Context, user entities.UserCreate) (int, error) {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return 0, nil
	}

	newUser := entities.UserCreate{
		UserBase: user.UserBase,
		Password: string(hashed_password),
	}

	id, err := usrs.UserRepo.Create(ctx, newUser)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (usrs UserServ) Get(ctx context.Context, id int) (*entities.User, error) {
	user, err := usrs.UserRepo.Get(ctx, id)
	if err != nil {
		return &entities.User{}, err
	}
	if user == nil {
		return nil, err
	}

	return user, nil

}

func (usrs UserServ) GetPassword(ctx context.Context, login string) (int, string, error) {
	id, password, err := usrs.UserRepo.GetPassword(ctx, login)
	if err != nil {
		return 0, "", err
	}

	return id, password, nil
}

func (usrs UserServ) UpdatePassword(ctx context.Context, UserID int, newPassword string) error {
	hashed_password, _ := bcrypt.GenerateFromPassword([]byte(newPassword), 10)

	err := usrs.UserRepo.UpdatePassword(ctx, UserID, string(hashed_password))
	if err != nil {
		return err
	}

	return nil
}

func (usrs UserServ) Delete(ctx context.Context, userID int) error {
	err := usrs.UserRepo.Delete(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}
