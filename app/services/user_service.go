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

func (r *UserService) GetAll() ([]models.User, error) {
	repository := userRepository.NewUserRepository()
	return repository.GetAll()
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

func (r *UserService) GetById(id string) (models.User, error) {
	repository := userRepository.NewUserRepository()

	user := repository.FindBy("ID", id)

	if user.ID == 0 {
		return models.User{}, errors.New("Usuário não encontrado!")
	}

	return user, nil
}

func (r *UserService) DeleteById(id string) (bool, error) {
	repository := userRepository.NewUserRepository()

	user := repository.FindBy("ID", id)

	if user.ID == 0 {
		return false, errors.New("Usuário não encontrado!")
	}

	repository.DeleteBy("ID", id)

	user = repository.FindBy("ID", id)

	if user.ID != 0 {
		return false, errors.New("Erro ao excluir conta!")
	}

	return true, nil
}

func (r *UserService) Update(id string, name string, username string, password string, cpf string, email string) models.User {
	repository := userRepository.NewUserRepository()

	user := repository.FindBy("ID", id)

	if name != "" {
		user.Name = name
	}
	if username != "" {
		user.Username = username
	}
	if password != "" {
		user.Password = password
	}
	if cpf != "" {
		user.CPF = cpf
	}
	if email != "" {
		user.Email = email
	}

	user = repository.Save(user)

	return user
}
