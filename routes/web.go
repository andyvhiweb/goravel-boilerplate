package routes

import (
	// "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Web() {
	facades.Route().Prefix("api/v1").Group(func(route route.Route) {
		authController := controllers.InitAuthController()
		route.Post("/auth/login", authController.Login)
	
		userController := controllers.NewUserController()
		route.Get("/users", userController.List)
		route.Post("/users", userController.Store)
		
		route.Get("/users/{id}", userController.Detail)

		photoController := controllers.InitPhotoController()
		route.Get("/photos", photoController.List)
		route.Post("/photos", photoController.Store)
		route.Get("/photos/{id}", photoController.Detail)
		route.Patch("/photos/{id}", photoController.Update)
		route.Post("/photos/{id}/like", photoController.Like)
		route.Post("/photos/{id}/unlike", photoController.Unlike)
	})

}
