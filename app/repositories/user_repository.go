package repositories

import (
	"goravel/app/models"

	"github.com/goravel/framework/facades"
)

type UserRepository struct {
	//Dependent services
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		//Inject services
	}
}

func (r *UserRepository) Store(name string, username string, password string, cpf string, email string) models.User {
	user := models.User{
		Name:     name,
		Username: username,
		Password: password,
		CPF:      cpf,
		Email:    email,
	}

	facades.Orm().Query().Where("username", username).FirstOrCreate(&user, user)
	return user
}
