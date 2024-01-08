package middleware

import (
	. "ShoppingMall/internal/model"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type MyClaims struct {
	User
	jwt.StandardClaims // jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 96 // 设置token过期时间
var MySecret = []byte("Yzc666")            // 设置加密字符串

// GenToken 生成JWT
func GenToken(user User) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		user, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "Yzc",                                      // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

//
//// ParseToken 解析JWT
//func ParseToken(tokenString string) (*MyClaims, error) {
//	// 解析token
//	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
//		return MySecret, nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
//		return claims, nil
//	}
//	return nil, errors.New("invalid token")
//}

// GenerateToken 登录成功后调用，传入User结构体,生成新的token
//func GenerateToken(userInfo User) (string, error) {
//	expirationTime := time.Now().Add(TokenExpireDuration) // 4天有效期
//	claims := &MyClaims{
//		User: userInfo,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: expirationTime.Unix(),
//			Issuer:    "Yzc",
//		},
//	}
//	// 生成Token，指定签名算法和claims
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	// 签名
//	if tokenString, err := token.SignedString(MySecret); err != nil {
//		return "", err
//	} else {
//		return tokenString, nil
//	}
//}

//// ParseToken 解析token
//func ParseToken(tokenString string) (*MyClaims, error) {
//	claims := &MyClaims{}
//	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
//		return MySecret, nil
//	})
//	// 若token只是过期claims是有数据的，若token无法解析claims无数据
//	return claims, err
//}

// ParseToken2 第二种方法通过jwt.ParseWithClaims返回的Token结构体取出Claims结构体
func ParseToken2(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token无法解析")
}

// JWTAuth 鉴权中间件
func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 从请求头中获取token
		auth := context.Request.Header.Get("token")

		if len(auth) == 0 {
			context.Abort()
			context.JSON(400, gin.H{
				"code":  400,
				"error": "token不能为空",
			})
			return
		}
		// 检查token是否过期
		if !CheckToken(auth) {
			context.Abort()
			context.JSON(400, gin.H{
				"code":  400,
				"error": "token过期",
			})
			return
		}
		// 校验token，只要出错直接拒绝请求，这边是解析token
		claims, err := ParseToken2(auth)
		if err != nil {
			context.Abort()
			message := err.Error()
			context.JSON(200, gin.H{
				"code":  400,
				"error": message,
			})
			return
		} else {
			log.Println(claims) // 打印出token中的信息 结构体
			println("token 正确")
			// 将当前请求的username信息保存到请求的上下文c上
			context.Set("token", claims)
		}
		context.Next()
	}
}

// CheckToken 检查token是否过期
func CheckToken(tokenString string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return false
	}
	if _, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return true
	}
	return false
}

func HaveToken(context *gin.Context) bool {
	// 从请求头中获取token
	auth := context.Request.Header.Get("token")
	if len(auth) == 0 {
		return false
	}
	// 检查token是否过期
	if !CheckToken(auth) {
		context.Abort()
		context.JSON(400, gin.H{
			"code":  400,
			"error": "token过期",
		})
	}
	// 校验token，只要出错直接拒绝请求，这边是解析token
	claims, err := ParseToken2(auth)
	if err != nil {
		context.Abort()
		message := err.Error()
		context.JSON(200, gin.H{
			"code":  400,
			"error": message,
		})
		return false
	} else {
		log.Println(claims) // 打印出token中的信息 结构体
		println("token 正确")
		// 将当前请求的username信息保存到请求的上下文c上
		context.Set("token", claims)
	}
	return true
}

// ImageVerificationCode 创建图形化验证码
func ImageVerificationCode(c *gin.Context) {
	id, b64s, _, err := CaptMake()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate captcha"})
		return
	}

	// 将验证码信息返回给前端
	response := map[string]string{
		"id":    id,
		"image": b64s,
	}
	c.JSON(200, response)
}

// VerifyGraphicVerificationCode 验证图形化验证码
func VerifyGraphicVerificationCode(c *gin.Context) {
	// 从前端获取验证码信息和用户输入
	var request map[string]string
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	fmt.Printf("this %v %[1]T\n", request)
	id := request["id"]
	userEnteredCode := request["userCode"]
	fmt.Printf("this1 %T %[1]v %[2]T  %[2]v \n", id, userEnteredCode)

	// 验证验证码
	if CaptVerify(id, userEnteredCode) {
		// 验证通过
		c.JSON(http.StatusOK, gin.H{"message": "验证通过", "code": 666})
	} else {
		// 验证失败
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Verification failed"})
	}
}

// SendMail 发电子邮件进行验证
func SendMail(c *gin.Context) {

}

// EmailVerificationCode 处理邮箱验证最终结果
func EmailVerificationCode(c *gin.Context) {
	// 从前端获取验证码信息和用户输入
	var request map[string]string
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	fmt.Printf("this %v %[1]T\n", request)
}
