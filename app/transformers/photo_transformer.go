package transformers

import (
	"github.com/goravel/framework/contracts/http"
	// "github.com/goravel/framework/facades"
	
	"goravel/app/models"
)

type PhotoTransformer struct {
	ID uint
	Likes int
	Path string
}

type PhotoTransformerArray struct {
	Data []PhotoTransformer
}

func (t *PhotoTransformer) Bind (data models.Photo) {
	t.ID	= data.ID
	t.Likes	= data.Likes
	t.Path	= data.Path
}

func (t *PhotoTransformer) ReturnWithItem (data models.Photo, ctx http.Context) any {
	t.Bind(data)
	// facades.Log().Debug(t)
	
	var m Transformer
	res := m.Item(t, ctx)

	return res
}

func (t *PhotoTransformer) ReturnWithPaginator (data []models.Photo, total int64, ctx http.Context) any {
	var u PhotoTransformerArray
	var s PhotoTransformer

	count := len(data)
	for i := 0; i < count; i++ {
		s.Bind(data[i])

		u.Data = append(u.Data, s)
	}

	var m Transformer
	res := m.Paginator(u.Data, count, total, ctx)

	return res
}

func (t *PhotoTransformer) ReturnWithCollection (data []models.Photo, ctx http.Context) any {
	var u PhotoTransformerArray
	var s PhotoTransformer

	for i := 0; i < len(data); i++ {
		s.Bind(data[i])

		u.Data = append(u.Data, s)
	}

	var m Transformer
	res := m.Items(u.Data, ctx)

	return res
}