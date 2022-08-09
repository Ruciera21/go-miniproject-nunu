package service

import (
	"log"
	"mini-project-nunu/dto"
	"mini-project-nunu/entity"
	"mini-project-nunu/repository"

	"github.com/mashingan/smapping"
)

// * UserService is a contract
type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

// * NewUserService is a new instance of UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepository: userRepo}
}

func (s *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("failed mapping %v: ", err)
	}
	updatedUser := s.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (s *userService) Profile(userID string) entity.User {
	return s.userRepository.ProfileUser(userID)
}
