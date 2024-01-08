package routers

import (
	"ShoppingMall/global"
	. "ShoppingMall/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetCommodity 得到全部商品信息
func (r *router) GetCommodity(c *gin.Context) {
	Commodity, err := Dao.GetCommodity(global.DBEngine)
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

// GetCommodityByID 根据商品ID得到相关商品信息
func (r *router) GetCommodityByID(c *gin.Context) {
	UidString := c.Param("id")
	Commodity, err := Dao.GetCommodityByID(global.DBEngine, UidString)
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

// PutCommodity 修改商品信息
func (r *router) PutCommodity(c *gin.Context) {
	var Commo Commodity
	if err := c.ShouldBindJSON(&Commo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
			"msg":   "请求参数错误",
		})
		return
	}
	fmt.Printf("this :%+v\n", Commo)
	if err := Dao.PutCommodity(global.DBEngine, &Commo); err != nil {
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

type image struct {
	Id  string `json:"id"`
	Val int    `json:"val"`
	Img string `json:"image"`
}

// PutCommodityImage 修改商品的图片的信息
func (r *router) PutCommodityImage(c *gin.Context) {
	var Image image
	if err := c.ShouldBindJSON(&Image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
			"msg":   "请求参数错误",
		})
		return
	} else {
		if err := Dao.PutCommodityImage(global.DBEngine, Image.Id, Image.Img, Image.Val); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "服务器错误",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "修改成功",
		})
	}
}

// DeleteCommodity 删除商品信息
func (r *router) DeleteCommodity(c *gin.Context) {
	idString := c.Param("id")
	// 检查商品是否存在
	if flag := Dao.CheckCommodity(global.DBEngine, idString); flag {
		// 删除商品信息
		if err := Dao.DeleteCommodity(global.DBEngine, idString); err != nil {
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
			"code": http.StatusBadRequest,
			"msg":  "该商品不存在",
		})
	}
}

// PostCommodity 添加商品信息
func (r *router) PostCommodity(c *gin.Context) {
	var commodity Commodity
	err := c.ShouldBindJSON(&commodity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
			"msg":   "请求参数错误",
		})
		return
	}
	if flag := Dao.CheckCommodityName(global.DBEngine, commodity.Name); flag {
		// 该商品种类已经存在
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "该商品已经存在",
		})
	} else {
		// 新增商品种类
		if err := Dao.PostCommodity(global.DBEngine, &commodity); err != nil {
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

// GetCommodityByCategoryID 根据商品种类得到那一类商品信息
func (r *router) GetCommodityByCategoryID(c *gin.Context) {
	idString := c.Param("id")
	Commodity, err := Dao.GetCommodityByCategoryID(global.DBEngine, idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "服务器错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "获取成功",
		"obj":  Commodity,
	})
}

// GetHomePageOnRecommend 得到所有HomePageOnRecommend字段为1 并且没有删除的商品
func (r *router) GetHomePageOnRecommend(c *gin.Context) {
	Commodity, err := Dao.GetHomePageOnRecommend(global.DBEngine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "服务器错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "获取成功",
		"obj":  Commodity,
	})
}

// GetCommodityByBatch 查询传入的批次的商品的信息
func (r *router) GetCommodityByBatch(c *gin.Context) {

	vals := c.Param("val")
	val, _ := strconv.Atoi(vals)
	data, err := Dao.GetCommodityByBatch(global.DBEngine, val)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
			"msg":   "请求参数错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "获取成功",
		"obj":  data,
	})
}
