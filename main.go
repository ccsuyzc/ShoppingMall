package main

import (
	"ShoppingMall/global"
	"ShoppingMall/internal/model"
	"ShoppingMall/internal/routers"
	"ShoppingMall/pkg/JWT"
	"ShoppingMall/pkg/logger"
	"ShoppingMall/pkg/setting"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"time"
)

// init
func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

func main() {
	rfunc := routers.NewRouter()
	r := gin.Default()

	// 使用CORS中间件,解决跨域问题
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true // 允许所有跨域

	// 设置单一站点
	//config.AllowOrigins = []string{"http://127.0.0.1:4000"} // 替换为你应用的实际源地址

	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Content-Type", "token"} // 添加 token 到允许的头部列表
	corsMiddleware := cors.New(config)
	r.Use(corsMiddleware)

	r.POST("/login", rfunc.Login)              // 登录
	r.POST("/login/rootuser", rfunc.LoginRoot) // 登录

	r.POST("/register", rfunc.Register)                                 // 注册
	r.GET("/login/token", middleware.JWTAuth(), rfunc.TokenIsEffective) // 验证token是否有效
	// 创建图形化验证码
	r.GET("/verification", middleware.ImageVerificationCode)

	// 处理前端提交的图像验证码和用户输入，进行验证
	r.POST("/verify", middleware.VerifyGraphicVerificationCode)

	// 处理前端提交的邮件验证码和用户邮箱进行登录
	r.POST("/login/email", middleware.EmailVerificationCode)

	// 向指定的邮箱发送邮件验证码
	r.POST("/sendmail", middleware.SendMail)

	rapi := r.Group("/api")

	// 用户相关操作
	rapi.GET("/users", rfunc.GetUsers)          //得到所有用户的信息
	rapi.GET("/users/:id", rfunc.GetUserByID)   // 得到指定用户的信息
	rapi.PUT("/users/:id", rfunc.PutUser)       // 修改用户权限
	rapi.DELETE("/users/:id", rfunc.DeleteUser) // 删除用户

	// 商品相关操作
	rapi.GET("/commoditys", rfunc.GetCommodity)                            // 得到所有商品的信息
	rapi.GET("/commoditys/:id", rfunc.GetCommodityByID)                    // 得到指定商品的信息
	rapi.PUT("/commoditys/:id", rfunc.PutCommodity)                        // 修改商品信息
	rapi.DELETE("/commoditys/:id", rfunc.DeleteCommodity)                  // 删除商品信息
	rapi.POST("/commoditys", rfunc.PostCommodity)                          // 新增商品信息
	rapi.GET("/commoditys/categoryID/:id", rfunc.GetCommodityByCategoryID) // 根据分类ID得到商品信息
	rapi.GET("/commoditys/isshow", rfunc.GetHomePageOnRecommend)           // 得到所有被推荐到首页的商品
	rapi.GET("/commoditys/value/:val", rfunc.GetCommodityByBatch)          // 分页器组件

	// 分类相关操作
	rapi.GET("/categories", rfunc.GetCategory)           // 得到所有分类的信息
	rapi.GET("/categories/:id", rfunc.GetCategoryByID)   // 得到指定分类的信息
	rapi.PUT("/categories/:id", rfunc.PutCategory)       // 修改分类信息
	rapi.DELETE("/categories/:id", rfunc.DeleteCategory) // 删除分类信息
	rapi.POST("/categories", rfunc.PostCategory)         // 添加分类

	// 购物车相关操作
	rapi.GET("/carts/:id", rfunc.GetCartByID)                    // 得到购物车的信息
	rapi.POST("/carts", rfunc.PostCart)                          // 添加购物车
	rapi.DELETE("/carts/:id", rfunc.DeleteCart)                  // 删除购物车
	rapi.PUT("/carts", rfunc.PutCart)                            // 修改购物车
	rapi.DELETE("/carts/user/:userid/:id", rfunc.DeleteCartByID) // 删除指定用户的购物车中指定的商品

	// 订单相关操作
	rapi.GET("/orders", rfunc.GetOrder)                                                   // 得到所有订单的信息
	rapi.GET("/orders/:id", rfunc.GetOrderByID)                                           // 根据订单id得到指定订单的信息
	rapi.GET("/orders/user/:userid", rfunc.GetOrderByUserID)                              // 根据用户id指定订单的信息
	rapi.GET("/orders/order/:orderid", rfunc.GetOrderByOrderID)                           // 根据订单id指定订单的信息
	rapi.GET("/orders/username/:username", rfunc.GetOrderByUserIDaaa)                     // 根据用户名查找指定订单的信息
	rapi.GET("/orders/phone/:phone/username/:username", rfunc.GetOrderByPhoneAndUsername) // 根据手机号和用户名查找指定订单的信息
	rapi.GET("/orders/phone/:phone", rfunc.GetOrderByPhone)                               // 根据手机号查找指定订单的信息
	rapi.POST("/orders", rfunc.PostOrder)                                                 // 加入订单
	rapi.DELETE("/orders/:id", rfunc.DeleteOrder)                                         // 删除订单
	rapi.PUT("/orders/:id", rfunc.PutOrder)                                               // 修改订单

	rapi.POST("/upload-image", routers.FileUpload)        // 上传图片
	rapi.GET("/images/:filename", routers.ImageAddress)   // 返回图片地址
	rapi.DELETE("/images/:filename", routers.DeleteImage) // 删除图片
	rapi.PUT("/commoditys", rfunc.PutCommodityImage)      // 修改商品图片

	r.Run(":9090")
}

// 初始化配置
func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

// 初始化数据库连接
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

// 初始化日志
func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
