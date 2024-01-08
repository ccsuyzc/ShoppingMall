package routers

import (
	"ShoppingMall/global"
	. "ShoppingMall/internal/model"
	"github.com/gin-gonic/gin"
)

// GetCartByID 通过用户id获取该用户的所有购物车信息
func (r *router) GetCartByID(c *gin.Context) {
	userId := c.Param("id")
	carts, err := Dao.GetCartByID(global.DBEngine, userId)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "数据库查询失败",
			"err":  err,
		})
	} else {
		c.JSON(200, gin.H{
			"code":   200,
			"msg":    "获取成功",
			"data":   carts,
			"length": len(carts),
		})
	}
}

// PutCart 修改购物车信息
func (r *router) PutCart(c *gin.Context) {
	var cart Cart
	err := c.ShouldBindJSON(&cart)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "数据格式错误",
			"err":  err,
		})
	} else {
		if flag1, _ := Dao.CheckCartID(global.DBEngine, cart.ID); flag1 { // 判断该订单是否存在
			err = Dao.PutCart(global.DBEngine, &cart)
			if err != nil {
				c.JSON(200, gin.H{
					"code": 500,
					"msg":  "数据库更新失败",
					"err":  err,
				})
			} else {
				c.JSON(200, gin.H{
					"code": 200,
					"msg":  "更新成功",
				})
			}
		} else {
			c.JSON(200, gin.H{
				"code": 500,
				"msg":  "该订单不存在",
			})
		}
	}
}

// DeleteCart 根据id删除购物车信息
func (r *router) DeleteCart(c *gin.Context) {
	id := c.Param("id")
	if flag1, _ := Dao.CheckCartID(global.DBEngine, id); flag1 { // 判断该订单是否存在
		err := Dao.DeleteCart(global.DBEngine, id)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 500,
				"msg":  "数据库删除失败",
				"err":  err,
			})
		} else {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "删除成功",
			})
		}
	} else {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "该订单不存在",
		})
	}
}

// PostCart 新增一条购物车数据
func (r *router) PostCart(c *gin.Context) {
	var cart Cart
	err := c.ShouldBindJSON(&cart)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "数据格式错误",
			"err":  err,
		})
	} else {
		err = Dao.PostCart(global.DBEngine, &cart)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 500,
				"msg":  "数据库插入失败",
				"err":  err,
			})
		} else {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "插入成功",
			})
		}
	}
}

// DeleteCartByID 删除指定用户的指定商品
func (r *router) DeleteCartByID(c *gin.Context) {
	userId := c.Param("userid")
	commodityId := c.Param("id")
	if flag1, _ := Dao.CheckCartUserIdAndID(global.DBEngine, userId, commodityId); flag1 { // 判断该订单是否存在
		err := Dao.DeleteCartByID(global.DBEngine, userId, commodityId)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 500,
				"msg":  "数据库删除失败",
				"err":  err,
			})
		} else {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "删除成功",
			})
		}
	} else {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "该订单不存在",
		})
	}

}
