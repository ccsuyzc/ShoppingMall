package model

import (
	"ShoppingMall/pkg/errcode"
	"github.com/jinzhu/gorm"
	"time"
)

// Order 订单表结构
type Order struct {
	ID          string    `json:"id" gorm:"primary_key;column:ID"`                 // 主键
	UserID      string    `json:"user_id" gorm:"not null;column:UserID"`           // 用户id
	CommodityID string    `json:"commodity_id" gorm:"not null;column:CommodityID"` // 商品id`
	Count       int       `json:"count" gorm:"not null;column:Count"`              // 商品数量
	Price       float64   `json:"price"gorm:"not null;column:Price"`               // 总金额
	Status      int       `json:"status" gorm:"not null;column:Status"`            // 状态 0 表示未支付，1 表示已支付，2 表示已发货，3 表示已完成，4 表示已取消
	PaymentType int       `json:"payment_type" gorm:"not null;column:PaymentType"` // 支付方式 0 表示未支付，1 表示微信支付，2 表示支付宝支付
	Address     string    `json:"address" gorm:"not null;column:Address"`          // 地址
	PayTime     time.Time `json:"pay_time" gorm:"column:PayTime"`                  // 支付时间
	Phone       string    `json:"phone" gorm:"not null;column:Phone"`              // 电话
	IsRemove    bool      `json:"is_remove" gorm:"not null;column:IsRemove"`       // 是否删除 false 表示未删除，true 表示删除
	OrderID     string    `json:"orderID" gorm:"column:OrderID"`                   // 订单ID
}

// GetOrderByOrderID 根据订单号查询订单信息
func (c *Order) GetOrderByOrderID(db *gorm.DB, OrderID string) ([]Order, error) {
	var order []Order
	err := db.Table("orders").Where("OrderID =?", OrderID).Find(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

// GetOrderByUserIDaaa 根据用户名查询该用户的所有订单信息
func (c *Order) GetOrderByUserIDaaa(db *gorm.DB, username string) ([]Order, error) {
	var order []Order
	// 先通过用户名查询到用户的ID
	userid, err0 := GetUserByUsername(db, username)
	//println("userid", userid)
	if err0 != nil {
		return nil, err0
	}
	err := db.Table("orders").Where("UserID =?", userid).Find(&order).Error
	if err != nil {
		return nil, err
	}
	//fmt.Printf("order:%v\n", order)
	return order, nil
}

// GetOrderByPhone 通过手机号查询订单信息
func (c *Order) GetOrderByPhone(db *gorm.DB, phone string) ([]Order, error) {
	var order []Order
	err := db.Table("orders").Where("Phone =?", phone).Find(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

// GetOrderByPhoneAndUsername 通过手机号和用户名查询到订单信息
func (c *Order) GetOrderByPhoneAndUsername(db *gorm.DB, phone string, username string) ([]Order, error) {
	var order []Order
	// 先通过用户名查询到用户的ID
	userid, err0 := GetUserByUsername(db, username)
	if err0 != nil {
		return nil, err0
	}
	err := db.Table("orders").Where("Phone =? and UserID =?", phone, userid).Find(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

// GetOrder 获取所有订单列表
func (c *Order) GetOrder(db *gorm.DB) ([]Order, error) {
	var orders []Order
	err := db.Table("orders").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// GetOrderByUserID 根据用户ID获取订单信息
func (c *Order) GetOrderByUserID(db *gorm.DB, UserID string) ([]Order, error) {
	var orders []Order
	err := db.Table("orders").Where("UserID =?", UserID).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// GetOrderByID 根据订单ID获取订单信息
func (c *Order) GetOrderByID(db *gorm.DB, id string) (*Order, error) {
	var order Order
	err := db.Table("orders").Where("ID =?", id).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// PutOrder 修改订单信息
func (c *Order) PutOrder(db *gorm.DB, order *Order) error {
	if flag1, _ := c.CheckOrderID(db, order.ID); flag1 { // 判断该商品种类是否存在

		err := db.Table("orders").Model(&order).Updates(order).Error
		if err != nil {
			return err
		}

	}
	return errcode.NotFound
}

// PostOrder 添加订单
func (c *Order) PostOrder(db *gorm.DB, order *Order) error {
	err := db.Table("orders").Create(&order).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteOrder 删除订单
func (c *Order) DeleteOrder(db *gorm.DB, id string) error {
	if flag1, _ := c.CheckOrderID(db, id); flag1 { // 判断该商品种类是否存在
		err := db.Table("orders").Where("ID =?", id).Delete(&Order{}).Error
		if err != nil {
			return err
		}
		return nil
	}
	return errcode.NotFound
}

// CheckOrderID 确认该订单是否存在
func (c *Order) CheckOrderID(db *gorm.DB, ID string) (bool, Order) {
	var order Order
	err := db.Table("orders").Where("ID =?", ID).First(&order).Error
	if err != nil {
		return false, order
	}
	return true, order
}
