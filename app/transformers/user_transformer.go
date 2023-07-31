package transformers

import (
	"github.com/goravel/framework/contracts/http"
	// "github.com/goravel/framework/facades"
	
	"goravel/app/models"
)

type UserTransformer struct {
	Name string
	Email string
}

type UserTransformerArray struct {
	Data []UserTransformer
}

func (t *UserTransformer) Bind (data models.User) {
	t.Name	= data.Name
	t.Email	= data.Email
}

func (t *UserTransformer) ReturnWithItem (data models.User, ctx http.Context) any {
	t.Bind(data)
	// facades.Log().Debug(t)
	
	var m Transformer
	res := m.Item(t, ctx)

	return res
}

func (t *UserTransformer) ReturnWithPaginator (data []models.User, total int64, ctx http.Context) any {
	var u UserTransformerArray
	var s UserTransformer

	count := len(data)
	for i := 0; i < count; i++ {
		s.Bind(data[i])

		u.Data = append(u.Data, s)
	}

	var m Transformer
	res := m.Paginator(u.Data, count, total, ctx)

	return res
}

func (t *UserTransformer) ReturnWithCollection (data []models.User, ctx http.Context) any {
	var u UserTransformerArray
	var s UserTransformer

	for i := 0; i < len(data); i++ {
		s.Bind(data[i])

		u.Data = append(u.Data, s)
	}

	var m Transformer
	res := m.Items(u.Data, ctx)

	return res
}