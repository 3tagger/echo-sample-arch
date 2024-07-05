package seeder

import (
	"context"

	"github.com/3tagger/echo-sample-arch/internal/user"
	"github.com/go-faker/faker/v4"
)

type userSeederService struct {
	userRepository user.Repository
}

func NewuserSeederService(userRepository user.Repository) *userSeederService {
	return &userSeederService{
		userRepository: userRepository,
	}
}

func (s *userSeederService) Fake(ctx context.Context) user.User {
	return user.User{
		Name:  faker.Name(),
		Email: faker.Email(),
		About: faker.Sentence(),
	}
}

func (s *userSeederService) InsertMany(ctx context.Context, arr []user.User) error {
	return s.userRepository.InsertMany(ctx, arr)
}
