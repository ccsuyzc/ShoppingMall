package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

// FileUpload 上传图片
func FileUpload(c *gin.Context) {
	file, err := c.FormFile("image")
	if err == nil {
		//2、获取后缀名 判断类型是否正确 .jpg .png .jpeg
		extName := path.Ext(file.Filename)
		allowExtMap := map[string]bool{
			".jpg":  true,
			".png":  true,
			".jpeg": true,
		}
		if _, ok := allowExtMap[extName]; !ok {
			c.JSON(http.StatusOK, gin.H{
				"code": 50001,
				"msg":  "文件类型不合法",
			})
			return
		}

		// 判断是否存在存储图片的目录，没有就创建
		if _, err := os.Stat("D:/GoProject/ShoppingMall/images"); os.IsNotExist(err) {
			os.Mkdir("D:/GoProject/ShoppingMall/images", os.ModePerm)
		}

		// 生成文件名
		filename := strconv.FormatInt(time.Now().UnixNano(), 10) + extName

		// 绝对地址
		filename2 := "D:/GoProject/ShoppingMall/images/" + filename
		if err := c.SaveUploadedFile(file, filename2); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}
		// 返回图片链接
		c.JSON(http.StatusOK, gin.H{"code": 200, "url": filename})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

// ImageAddress 返回图片地址
func ImageAddress(c *gin.Context) {
	filename := c.Param("filename")
	c.File("D:/GoProject/ShoppingMall/images/" + filename)
}

// DeleteImage 删除图片
func DeleteImage(c *gin.Context) {
	filename := c.Param("filename")
	err := os.Remove("D:/GoProject/ShoppingMall/images/" + filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200})
}
