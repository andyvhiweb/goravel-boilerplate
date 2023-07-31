package transformers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	// mock "github.com/stretchr/testify/mock"

	"strconv"
	"reflect"
)

type Transformer struct {
	// mock.Mock
}

func (trans *Transformer) Paginator(data any, count int, total int64, ctx http.Context) any {

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

	return map[string]any {
		"Data": data,
		"Meta": map[string]any {
			"Pagination": map[string]any {
				"Total": total,
				"Count": count,
				"PerPage": limit,
				"CurrentPage": page,
				"TotalPages": int64(float64(total)/float64(limit))+1,
			},
		},
	}
}

func (trans *Transformer) Items(data any, ctx http.Context) any {
	facades.Log().Debug(reflect.TypeOf(data))
	return map[string]any {
		"data": data,
	}
}

func (trans *Transformer) Item(data any, ctx http.Context) any {

	return map[string]any {
		"data": data,
	}
}

// Bind provides a mock function with given fields: obj
// func (_m *Request) Bind(obj interface{}) error {
// 	ret := _m.Called(obj)

// 	var r0 error
// 	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
// 		r0 = rf(obj)
// 	} else {
// 		r0 = ret.Error(0)
// 	}

// 	return r0
// }
