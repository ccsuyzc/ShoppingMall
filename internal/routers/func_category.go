package routers

import (
	"ShoppingMall/global"
	. "ShoppingMall/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCategory 得到全部商品信息
func (r *router) GetCategory(c *gin.Context) {
	category, err := Dao.GetCategory(global.DBEngine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "服务器错误",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "获取成功",
			"obj":  category,
		})
	}
}

// GetCategoryByID 根据商品ID得到相关商品信息
func (r *router) GetCategoryByID(c *gin.Context) {
	UidString := c.Param("id")
	Commodity, err := Dao.GetCategoryByID(global.DBEngine, UidString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "服务器错误",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "获取成功",
			"obj":  Commodity,
		})
	}
}

// PutCategory 修改商品信息
func (r *router) PutCategory(c *gin.Context) {
	var category Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
			"msg":   "请求参数错误",
		})
		return
	}
	if err := Dao.PutCategory(global.DBEngine, &category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "服务器错误",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "修改成功",
		})
	}
}

// DeleteCategory 删除商品种类信息
func (r *router) DeleteCategory(c *gin.Context) {
	idString := c.Param("id")
	fmt.Printf("is:%v\n", idString)
	if flag, _ := Dao.CheckCategoryID(global.DBEngine, idString); flag {
		fmt.Printf("One:%v,%v\n", flag, idString)
		if err := Dao.DeleteCategory(global.DBEngine, idString); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "服务器错误",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "删除成功",
			})
		}
	} else {
		// 该商品不存在
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400001,
			"msg":  "该种类不存在",
		})
	}
}

// PostCategory 添加商品信息
func (r *router) PostCategory(c *gin.Context) {
	var category Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
			"msg":   "请求参数错误",
		})
		return
	}
	if flag := Dao.CheckCategoryName(global.DBEngine, category.Name); flag {
		// 该商品种类已经存在
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "该商品已经存在",
		})
	} else {
		// 新增商品种类
		if err := Dao.PostCategory(global.DBEngine, &category); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "服务器错误",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
			})
		}
	}
}
