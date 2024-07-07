package usecase

import (
	"context"

	"github.com/3tagger/echo-sample-arch/internal/apperror"
	"github.com/3tagger/echo-sample-arch/internal/user"
)

type UserUsecase struct {
	userRepository user.Repository
}

func NewUserUsecase(userRepository user.Repository) *UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

func (u *UserUsecase) GetAll(ctx context.Context) ([]*user.User, error) {
	res, err := u.userRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *UserUsecase) GetOneById(ctx context.Context, userId int64) (*user.User, error) {
	res, err := u.userRepository.GetOneById(ctx, userId)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, apperror.ErrUserWithIdNotFound(userId)
	}

	return res, nil
}

func (u *UserUsecase) InsertOne(ctx context.Context, user user.User) (*user.User, error) {
	res, err := u.userRepository.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *UserUsecase) InsertMany(ctx context.Context, users []user.User) error {
	err := u.userRepository.InsertMany(ctx, users)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) DeleteOne(ctx context.Context, userId int64) error {
	err := u.userRepository.DeleteOne(ctx, userId)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) UpdateOneById(ctx context.Context, userId int64, user user.User) error {
	err := u.userRepository.UpdateOneById(ctx, userId, user)
	if err != nil {
		return err
	}

	return nil
}
