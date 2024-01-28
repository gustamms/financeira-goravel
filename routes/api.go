package routes

import (
	"goravel/app/http/controllers"

	"github.com/goravel/framework/facades"
)

func Api() {
	facades.Route().Resource("/users", controllers.NewUserController())
}
