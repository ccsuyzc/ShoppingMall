package routers

import (
	"ShoppingMall/global"
	. "ShoppingMall/internal/model"
	middleware "ShoppingMall/pkg/JWT"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type LoginClass struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login 登录
func (r *router) Login(c *gin.Context) {
	//if()
	var userA LoginClass
	if err := c.ShouldBindJSON(&userA); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "请求参数错误",
		})
		return
	}

	if userA.Password == "" || userA.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "邮箱和密码不能为空",
		})
		return
	}

	userB, err7 := Dao.GetUserByEmail(global.DBEngine, userA.Email)
	fmt.Printf("userB: %+v, error: %v\n", userB, err7)
	if errors.Is(err7, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"msg":  "用户不存在",
		})
		return
	} else if err7 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "服务器错误",
		})
		return
	}

	fmt.Printf("B:%+v\nA:%+v\n", userB, userA)
	// Verify Password
	if flag := userB.Password == userA.Password; !flag {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "密码错误",
		})
		return
	}

	// Generate Token
	token, err4 := middleware.GenToken(userB)
	if err4 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "服务器错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "登录成功",
		"token": token,
		"obj":   userB,
	})
}

// LoginRoot 登录
func (r *router) LoginRoot(c *gin.Context) {
	//if()
	var userA LoginClass
	if err := c.ShouldBindJSON(&userA); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "请求参数错误",
		})
		return
	}

	if userA.Password == "" || userA.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "邮箱和密码不能为空",
		})
		return
	}

	userB, err7 := Dao.GetUserByEmailRoot(global.DBEngine, userA.Email)
	fmt.Printf("userB: %+v, error: %v\n", userB, err7)
	if errors.Is(err7, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"msg":  "用户不存在",
		})
		return
	} else if err7 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "服务器错误",
		})
		return
	}

	fmt.Printf("B:%+v\nA:%+v\n", userB, userA)
	// Verify Password
	if flag := userB.Password == userA.Password; !flag {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "密码错误",
		})
		return
	}

	// Generate Token
	token, err4 := middleware.GenToken(userB)
	if err4 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "服务器错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "登录成功",
		"token": token,
		"obj":   userB,
	})
}

// Register 注册
func (r *router) Register(c *gin.Context) {
	var UserA User
	if err := c.ShouldBindJSON(&UserA); err == nil {
		if UserA.Password == "" {
			c.JSON(400, gin.H{
				"code": 400001,
				"msg":  "密码不能为空",
			})
			return
		}
		if UserA.Username == "" {
			c.JSON(400, gin.H{
				"code": 400004,
				"msg":  "用户名不能为空",
			})
			return
		}

		// 检查用户名是否已经存在
		UserB, err := Dao.GetUser(global.DBEngine, UserA.Username)
		fmt.Printf("this is %+v\n", UserB)
		if errors.Is(err, gorm.ErrRecordNotFound) {

			// 插入数据
			if err := Dao.AddUser(global.DBEngine, &UserA); err != nil {
				c.JSON(400, gin.H{
					"code": 4000004,
					"msg":  "注册失败",
					"err":  err,
				})
			} else {
				//	 生成token
				token, _ := middleware.GenToken(UserA)
				c.JSON(200, gin.H{
					"code":  200000,
					"msg":   "注册成功",
					"token": token,
				})
			}
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  "服务器错误",
			})
			return
		}

	}
}

// TokenIsEffective token是否有效
func (r *router) TokenIsEffective(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "登录成功",
	})
}

// GetUsers 获取用户列表
func (r *router) GetUsers(c *gin.Context) {
	Users, err := Dao.GetUsers(global.DBEngine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "服务器错误",
			"err":  err,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "获取成功",
			"obj":  Users,
		})
	}
}

// GetUserByID 获取单个用户信息
func (r *router) GetUserByID(c *gin.Context) {
	UidString := c.Param("id")
	user, err := Dao.GetUserByID(global.DBEngine, UidString)
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
			"obj":  user,
		})
	}
}

// DeleteUser 删除单个用户，gorm的软删除
func (r *router) DeleteUser(c *gin.Context) {
	UidString := c.Param("id")
	if err := Dao.DeleteUser(global.DBEngine, UidString); err != nil {
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
}

// PutUser  修改单个用户信息
func (r *router) PutUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
			"msg":   "请求参数错误",
		})
		return
	}
	if err := Dao.PutUser(global.DBEngine, &user); err != nil {
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
