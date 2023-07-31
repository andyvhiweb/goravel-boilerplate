package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"goravel/app/models"
)

type AuthController struct {
	//Dependent services
}

func InitAuthController() *AuthController {
	return &AuthController{
		//Inject services
	}
}

func (r *AuthController) Login(ctx http.Context) {
	var user models.User

	ctx.Request().Bind(&user)

	password := user.Password

	facades.Orm().Query().Where("email = ?", user.Email).First(&user)

	if facades.Hash().Check(password, user.Password) {
		// The passwords match...

		token, err := facades.Auth().LoginUsingID(ctx, 1)

		if err != nil {
			facades.Log().Debug(err)
		}

		res := map[string]any {
			"access_token": token,
		}
		ctx.Response().Success().Json(res)

	} else {
		res := map[string]any {
			"error": "User not found.",
		}
		ctx.Response().Success().Json(res)
	}
}
