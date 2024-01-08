package dao

import (
	. "ShoppingMall/internal/model"
)

type Dao struct {
	User
	Commodity
	Category
	Cart
	Order
}

func NewDao() *Dao {
	return &Dao{}
}
