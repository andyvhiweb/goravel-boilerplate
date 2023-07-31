package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"goravel/app/models"
	"goravel/app/transformers"
	"strconv"
)

type PhotoController struct {
	//Dependent services
}

func InitPhotoController() *PhotoController {
	return &PhotoController{
		//Inject services
	}
}

func (r *PhotoController) List(ctx http.Context) {
	var photos []models.Photo
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

	facades.Orm().Query().Paginate(page, limit, &photos, &total)

	var photoTransformer transformers.PhotoTransformer
	res := photoTransformer.ReturnWithPaginator(photos, total, ctx)

	ctx.Response().Success().Json(res)
}

func (r *PhotoController) Detail(ctx http.Context) {
	var photo models.Photo
	var res any

	//Get photo by ID
	id := ctx.Request().Route("id")
	facades.Orm().Query().Find(&photo, id)
	
	var photoTransformer transformers.PhotoTransformer
	if photo.ID != 0 {
		res = photoTransformer.ReturnWithItem(photo, ctx)
	} else {
		res = r.ReturnError("Photo not found", ctx)
	}

	ctx.Response().Success().Json(res)
}


func (r *PhotoController) Store(ctx http.Context) {
	var res any
	isError := false

	token := ctx.Request().Header("token")

	payload, err := facades.Auth().Parse(ctx, token)
	if err != nil {
		isError = true
		res = r.ReturnError(err, ctx)
	} else {
		facades.Log().Debug(payload)
	}

	if !isError {
		var user models.User

		err = facades.Auth().User(ctx, &user)
		if err != nil {
			isError = true
			res = r.ReturnError(err, ctx)
		}

		if !isError {
			//Get File

			file, err := ctx.Request().File("file")
			if err != nil {
				facades.Log().Debug(err)
			}

			storedFile, err := file.Store("./public")
			if err != nil {
				isError = true
				res = r.ReturnError(err, ctx)
			}
	
			if !isError {
				//Save File Path to Photo
				var photo models.Photo
				photo.UserID = user.ID
				photo.Path = storedFile
			
				facades.Orm().Query().Create(&photo)
		
				var photoTransformer transformers.PhotoTransformer
				res = photoTransformer.ReturnWithItem(photo, ctx)
			}
		}
	}

	ctx.Response().Success().Json(res)
}

func (r *PhotoController) Update(ctx http.Context) {
	var res any
	isError := false

	token := ctx.Request().Header("token")

	payload, err := facades.Auth().Parse(ctx, token)
	if err != nil {
		isError = true
		res = r.ReturnError(err, ctx)
	} else {
		facades.Log().Debug(payload)
	}

	var user models.User
	
	err = facades.Auth().User(ctx, &user)
	if err != nil {
		isError = true
		res = r.ReturnError(err, ctx)
	}

	//Cek photo with user login
	id := ctx.Request().Route("id")

	var photo models.Photo
	facades.Orm().Query().Find(&photo, id)

	if photo.UserID != user.ID {
		isError = true
		res = r.ReturnError("Photo not found", ctx)
	}

	//Get File
	file, err := ctx.Request().File("file")
	if err != nil {
		isError = true
		res = r.ReturnError(err, ctx)
	}

	if (!isError) {
		storedFile, err := file.Store("./public")
		if err != nil {
			isError = true
			res = r.ReturnError(err, ctx)
		} else {
			//Delete Existing File
			err = facades.Storage().Delete(photo.Path)
			if err != nil {
				isError = true
				res = r.ReturnError(err, ctx)
			}

			//Save File Path to Photo
			photo.Path = storedFile
		
			facades.Orm().Query().Save(&photo)
		
			var photoTransformer transformers.PhotoTransformer
			res = photoTransformer.ReturnWithItem(photo, ctx)
		}
	}

	ctx.Response().Success().Json(res)

}

func (r *PhotoController) Like(ctx http.Context) {
	var res any
	// isError := false

	//Get photo
	id := ctx.Request().Route("id")

	var photo models.Photo
	facades.Orm().Query().Find(&photo, id)

	//Check photo is exists
	if photo.ID == 0 {
		res = r.ReturnError("Photo not found", ctx)
	} else {
		//Increment likes
		photo.Likes++
		facades.Orm().Query().Save(&photo)

		var photoTransformer transformers.PhotoTransformer
		res = photoTransformer.ReturnWithItem(photo, ctx)
	}


	ctx.Response().Success().Json(res)
}

func (r *PhotoController) Unlike(ctx http.Context) {
	var res any
	// isError := false

	//Get photo
	id := ctx.Request().Route("id")

	var photo models.Photo
	facades.Orm().Query().Find(&photo, id)

	//Check photo is exists
	if photo.ID == 0 {
		res = r.ReturnError("Photo not found", ctx)
	} else {
		//Increment likes
		photo.Likes--

		if photo.Likes < 0 {
			photo.Likes = 0
		}

		facades.Orm().Query().Save(&photo)

		var photoTransformer transformers.PhotoTransformer
		res = photoTransformer.ReturnWithItem(photo, ctx)
	}


	ctx.Response().Success().Json(res)
}

func (r *PhotoController) ReturnError(message any, ctx http.Context) any {
	res := map[string]any {
		"error": message,
	}
	return res
}