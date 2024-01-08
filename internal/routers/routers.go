package routers

import (
	"ShoppingMall/internal/dao"
)

type router struct {
}

func NewRouter() *router {
	return &router{}
}

var Dao = dao.NewDao()
