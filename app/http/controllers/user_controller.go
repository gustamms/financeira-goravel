package controllers

import (
	userService "goravel/app/services"

	"github.com/goravel/framework/contracts/http"
)

type UserController struct {
	//Dependent services
	userService.UserService
}

func NewUserController() *UserController {
	return &UserController{
		//Inject services
	}
}

func (r *UserController) Index(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"Hello": "Goravel",
	})
}

func (r *UserController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Input("id")
	user := userService.NewUserService()
	userInformation, err := user.GetById(id)

	if err != nil {
		return ctx.Response().Status(404).Json(http.Json{
			"message": err.Error(),
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"message": userInformation,
	})
}

func (r *UserController) Store(ctx http.Context) http.Response {
	validator, _ := ctx.Request().Validate(map[string]string{
		"name":     "required|max_len:255",
		"username": "required|max_len:255",
		"password": "required|max_len:255",
		"cpf":      "required|max_len:11|min_len:10",
		"email":    "required|max_len:255",
	})

	if validator.Fails() {
		return ctx.Response().Success().Json(http.Json{
			"message": validator.Errors().All(),
		})
	}

	name := ctx.Request().Input("name")
	username := ctx.Request().Input("username")
	password := ctx.Request().Input("password")
	cpf := ctx.Request().Input("cpf")
	email := ctx.Request().Input("email")

	user := userService.NewUserService()
	userInformation, err := user.Store(name, username, password, cpf, email)

	if err != nil {
		return ctx.Response().Status(400).Json(http.Json{
			"message": err.Error(),
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"message": userInformation,
	})
}

func (r *UserController) Update(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"Hello": "Goravel",
	})
}

func (r *UserController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Input("id")
	user := userService.NewUserService()
	deleteConfirmation, err := user.DeleteById(id)

	if deleteConfirmation == false {
		return ctx.Response().Status(400).Json(http.Json{
			"message": err.Error(),
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"message": "Sucesso!",
	})
}
