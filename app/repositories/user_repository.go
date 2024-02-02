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

func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	facades.Orm().Query().Find(&users)
	return users, nil
}

func (r *UserRepository) Store(name string, username string, password string, cpf string, email string) models.User {
	user := models.User{
		Name:     name,
		Username: username,
		Password: password,
		CPF:      cpf,
		Email:    email,
	}

	facades.Orm().Query().Where("cpf", cpf).FirstOrCreate(&user, user)
	return user
}

func (r *UserRepository) FindBy(field string, value string) models.User {
	var users models.User
	facades.Orm().Query().Where(field, value).Get(&users)
	return users
}

func (r *UserRepository) DeleteBy(field string, value string) models.User {
	var users models.User
	facades.Orm().Query().Where(field, value).Delete(&users)
	return users
}

func (r *UserRepository) Save(user models.User) models.User {
	facades.Orm().Query().Save(user)
	return user
}
