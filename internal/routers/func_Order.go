package routers

import (
	"ShoppingMall/global"
	. "ShoppingMall/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

// GetOrder 获取全部订单
func (r *router) GetOrder(c *gin.Context) {
	order, err := Dao.GetOrder(global.DBEngine)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "获取订单失败",
			"err":  err,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "获取订单成功",
			"data": order,
		})
	}
}

// GetOrderByUserID 根据用户ID获取订单信息
func (r *router) GetOrderByUserID(c *gin.Context) {
	userID := c.Param("userid")
	order, err := Dao.GetOrderByUserID(global.DBEngine, userID)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "获取订单失败",
			"err":  err,
		})
	} else {
		c.JSON(200, gin.H{
			"code":   200,
			"msg":    "获取订单成功",
			"data":   order,
			"length": len(order),
		})
	}
}

// GetOrderByID 根据订单ID获取订单信息
func (r *router) GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	order, err := Dao.GetOrderByID(global.DBEngine, id)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "获取订单失败",
			"err":  err,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "获取订单成功",
			"data": order,
		})
	}
}

// PutOrder 修改订单信息
func (r *router) PutOrder(c *gin.Context) {
	var order Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "传参错误",
			"err":  err,
		})
	} else {
		err := Dao.PostOrder(global.DBEngine, &order)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 500,
				"msg":  "修改订单失败",
				"err":  err,
			})
		} else {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "修改订单成功",
			})
		}
	}
}

// PostOrder 添加订单
func (r *router) PostOrder(c *gin.Context) {
	var order Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "传参错误",
			"err":  err,
		})
	} else {
		err := Dao.PostOrder(global.DBEngine, &order)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 500,
				"msg":  "添加订单失败",
				"err":  err,
			})
		} else {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "添加订单成功",
			})
		}
	}
}

// DeleteOrder 通过订单号删除订单
func (r *router) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	err := Dao.DeleteOrder(global.DBEngine, id)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "删除订单失败",
			"err":  err,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "删除订单成功",
		})
	}
}

// GetOrderByOrderID 根据订单号查询订单信息
func (r *router) GetOrderByOrderID(c *gin.Context) {
	orderID := c.Param("orderid")
	order, err := Dao.GetOrderByOrderID(global.DBEngine, orderID)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "查询订单失败",
			"err":  err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询订单成功",
		"data": order,
	})

}

// GetOrderByUserIDaaa 根据用户名查询该用户的所有订单信息
func (r *router) GetOrderByUserIDaaa(c *gin.Context) {
	username := c.Param("username")
	order, err := Dao.GetOrderByUserIDaaa(global.DBEngine, username)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "查询订单失败",
			"err":  err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询订单成功",
		"data": order,
	})
}

// GetOrderByPhone 通过手机号查询订单信息
func (r *router) GetOrderByPhone(c *gin.Context) {
	phone := c.Param("phone")
	order, err := Dao.GetOrderByPhone(global.DBEngine, phone)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "查询订单失败",
			"err":  err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询订单成功",
		"data": order,
	})
}

// GetOrderByPhoneAndUsername 通过手机号和用户名查询到订单信息
func (r *router) GetOrderByPhoneAndUsername(c *gin.Context) {
	phone := c.Param("phone")
	username := c.Param("username")
	fmt.Printf("%v %v \n", phone, username)
	order, err := Dao.GetOrderByPhoneAndUsername(global.DBEngine, phone, username)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "查询订单失败",
			"err":  err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询订单成功",
		"data": order,
	})

}
