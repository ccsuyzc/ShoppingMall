package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"gopkg.in/gomail.v2"
	"image/color"
	"math/rand"
	"net/http"
	"time"
)

var stores = base64Captcha.DefaultMemStore

// CaptMake 创建验证码
func CaptMake() (id, b64s string, code string, err error) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString
	// 配置验证码信息
	captchaConfig := base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, stores)
	lid, lb64s, _, lerr := captcha.Generate()
	code = stores.Get(lid, false)
	fmt.Println(code)
	return lid, lb64s, code, lerr
}

// CaptVerify 解析验证码
func CaptVerify(id string, capt string) bool {
	if stores.Verify(id, capt, true) {
		return true
	} else {
		return false
	}
}

// 生成指定长度的随机验证码
func generateRandomCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	charset := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	code := make([]byte, length)
	for i := 0; i < length; i++ {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}

// 发送验证码邮件
func sendVerificationCode(email string) (string, error) {
	// 生成6位验证码
	code := generateRandomCode(6)

	// 设置 QQ 邮箱 SMTP 服务器信息
	d := gomail.NewDialer("smtp.qq.com", 587, "your@qq.com", "your-smtp-authorization-code")

	// 创建邮件
	m := gomail.NewMessage()
	m.SetHeader("From", "your@qq.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "验证码")
	m.SetBody("text/plain", fmt.Sprintf("您的验证码是：%s", code))

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		return "", err
	}

	return code, nil
}

func Dev() {
	email := "recipient@example.com" // 替换为实际的收件人邮箱
	code, err := sendVerificationCode(email)
	if err != nil {
		fmt.Println("Failed to send verification code:", err)
		return
	}

	fmt.Println("Verification code sent successfully. Code:", code)
}

func DevIsOk(c *gin.Context) {
	email := c.PostForm("email")
	userCode := c.PostForm("code") // 用户输入的验证码

	// 在这里调用你的验证码验证函数
	if verifyCode(email, userCode) {
		c.JSON(http.StatusOK, gin.H{"message": "验证码验证成功"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "验证码验证失败"})
	}
}

func verifyCode(email, userCode string) bool {
	// 这里应该调用你的验证码验证逻辑
	// 比较 userCode 是否与之前发送的验证码匹配
	return userCode == "123456" // 替换为实际的验证码验证逻辑
}
