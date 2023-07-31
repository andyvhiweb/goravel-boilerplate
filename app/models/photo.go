package models

import (
	"github.com/goravel/framework/database/orm"
)

type Photo struct {
	orm.Model
	UserID uint
	Likes int
	Path string
}
