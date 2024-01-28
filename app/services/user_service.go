package services

import (
	"goravel/app/models"
	userRepository "goravel/app/repositories"
)

type UserService struct {
	//Dependent services
	userRepository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		//Inject services
	}
}

func (r *UserService) Store(name string, username string, password string, cpf string, email string) models.User {
	repository := userRepository.NewUserRepository()
	user := repository.Store(name, username, password, cpf, email)
	return user
}
