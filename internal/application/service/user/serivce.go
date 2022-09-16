package user

import (
	"github.com/PanovAlexey/learn-subtitles/internal/domain/dto"
	"github.com/PanovAlexey/learn-subtitles/internal/domain/entity"
)

type UserRepository interface {
	SaveUser(user entity.User) (dto.UserDatabaseDto, error)
	GetUserByLogin(login string) (dto.UserDatabaseDto, error)
}

type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) UserService {
	return UserService{userRepository: userRepository}
}

func (s UserService) SaveUser(user entity.User) (dto.UserDatabaseDto, error) {
	return s.userRepository.SaveUser(user)
}

func (s UserService) GetUserByLogin(login string) (dto.UserDatabaseDto, error) {
	user, err := s.userRepository.GetUserByLogin(login)

	if err != nil {
		return user, err
	}

	return user, err
}
