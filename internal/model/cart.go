package model

import (
	"ShoppingMall/pkg/errcode"
	"github.com/jinzhu/gorm"
)

// Cart 购物车表结构
type Cart struct {
	ID          string `json:"id" gorm:"primary_key;column:ID"`                 // 主键
	UserID      string `json:"user_id" gorm:"not null;column:UserID"`           // 用户id
	CommodityID string `json:"commodity_id" gorm:"not null;column:CommodityID"` // 商品id
	Count       int    `json:"count" gorm:"not null;column:Count"`              // 商品数量                                          // 商品数量
	Price       int    `json:"price" gorm:"not null;column:Price"`              // 商品价格                          // 商品价格
}

// GetCart 获取购物车所有订单
func (c *Cart) GetCart(db *gorm.DB) ([]Cart, error) {
	var cart []Cart
	err := db.Table("carts").Find(&cart).Error
	if err != nil {
		return nil, err
	}
	return cart, nil
}

// GetCartByID 根据用户ID获取他的所有订单
func (c *Cart) GetCartByID(db *gorm.DB, id string) ([]Cart, error) {
	var cart []Cart
	err := db.Table("carts").Where("UserID =?", id).Find(&cart).Error
	if err != nil {
		return nil, err
	}
	return cart, nil
}

// PutCart 修改订单种类
func (c *Cart) PutCart(db *gorm.DB, cart *Cart) error {
	if flag1, _ := c.CheckCartID(db, cart.ID); flag1 { // 判断该订单是否存在
		err := db.Table("carts").Model(&cart).Updates(cart).Error
		if err != nil {
			return err
		}
	} else {
		return errcode.NotFound
	}
	return nil
}

// PostCart 添加订单
func (c *Cart) PostCart(db *gorm.DB, cart *Cart) error {
	err := db.Table("carts").Create(&cart).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteCart 删除订单
func (c *Cart) DeleteCart(db *gorm.DB, id string) error {
	if flag1, _ := c.CheckCartID(db, id); flag1 { // 判断该订单是否存在
		err := db.Table("carts").Where("ID =?", id).Delete(&Cart{}).Error
		if err != nil {
			return err
		}
		return nil
	}
	return errcode.NotFound
}

// CheckCartName 确认该用户是否存在
func (c *Cart) CheckCartName(db *gorm.DB, Name string) bool {
	var cart Cart
	err := db.Table("carts").Where("Name =?", Name).First(&cart).Error
	if err != nil {
		return false
	}
	return true
}

// CheckCartID 确认该订单是否存在
func (c *Cart) CheckCartID(db *gorm.DB, ID string) (bool, Cart) {
	var cart Cart
	err := db.Table("carts").Where("ID =?", ID).First(&cart).Error
	if err != nil {
		return false, cart
	}
	return true, cart
}

// DeleteCartByID 删除特定用户的特定购物车的一类数据
func (c *Cart) DeleteCartByID(db *gorm.DB, UserID string, id string) error {
	err := db.Table("carts").Where("CommodityID = ? AND UserID = ?", id, UserID).Delete(&Cart{}).Error
	if err != nil {
		return err
	}
	return nil
}

// CheckCartUserIdAndID 确认该订单是否存在
func (c *Cart) CheckCartUserIdAndID(db *gorm.DB, userID string, ID string) (bool, Cart) {
	var cart Cart
	err := db.Table("carts").Where("CommodityID =? AND UserID = ?", ID, userID).First(&cart).Error
	if err != nil {
		return false, cart
	}
	return true, cart
}
