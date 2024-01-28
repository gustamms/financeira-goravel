package services

import (
	"errors"
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

func (r *UserService) Store(name string, username string, password string, cpf string, email string) (models.User, error) {
	repository := userRepository.NewUserRepository()

	userInformationEmail := repository.FindBy("email", email)
	if userInformationEmail.ID != 0 {
		return models.User{}, errors.New("Usuário já registrado!")
	}

	userInformationCPF := repository.FindBy("cpf", cpf)
	if userInformationCPF.ID != 0 {
		return models.User{}, errors.New("Usuário já registrado!")
	}

	user := repository.Store(name, username, password, cpf, email)
	return user, nil
}
