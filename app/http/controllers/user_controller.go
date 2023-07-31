package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"goravel/app/models"
	"goravel/app/transformers"
	"strconv"
)

type UserController struct {
	//Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		//Inject services
	}
}

func (r *UserController) List(ctx http.Context) {
	var users []models.User
	var total int64
	
	//Get page param
	page, err := strconv.Atoi(ctx.Request().Query("page"))
	if err != nil {
		facades.Log().Debug(err)
		page = 1
	}

	//Get page param
	limit, err := strconv.Atoi(ctx.Request().Query("limit"))
	if err != nil {
		facades.Log().Debug(err)
		limit = 10
	}

	facades.Orm().Query().Paginate(page, limit, &users, &total)

	var userTransformer transformers.UserTransformer
	// res := userTransformer.ReturnWithPaginator(users, total, ctx)
	res := userTransformer.ReturnWithCollection(users, ctx)
	
	ctx.Response().Success().Json(res)
}

func (r *UserController) Detail(ctx http.Context) {
	var user models.User
	var userTransformer transformers.UserTransformer
	
	id := ctx.Request().Route("id")

	facades.Orm().Query().Find(&user, id)
	
	// userTransformer.Bind(user);

	// facades.Log().Debug(userTransformer)

	// var trans transformers.Transformer
	res := userTransformer.ReturnWithItem(user, ctx)

	ctx.Response().Success().Json(res)
}

func (r *UserController) Store(ctx http.Context) {
	var user models.User

	ctx.Request().Bind(&user)

	password, err := facades.Hash().Make(user.Password)

	user.Password = password
	facades.Log().Debug(err)

	facades.Orm().Query().Save(&user)
}

func (r *UserController) Show(ctx http.Context) {
	ctx.Response().Success().Json(http.Json{
		"Hello": "Goravel",
	})
}
