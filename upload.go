package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// UploadHandler POST Upload images files 单文件上传至upload文件夹中
func UploadHandler(c *gin.Context) {
	file, _ := c.FormFile("file")
	// the path of document files
	fileName := uuid.Must(uuid.NewV4()).String() + "." + strings.Split(file.Filename, ".")[1]

	dst := "./dist/upload/" + fileName
	// Upload the file to specific dst.
	c.SaveUploadedFile(file, dst)

	c.JSON(http.StatusOK, map[string]interface{}{
		"file_name": fileName,
	})
}

// UploadsHandler UploadHandler POST Upload temp files 多文件上传至tmp文件目录中
func UploadsHandler(c *gin.Context) {
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		dst := "./upload/" + file.Filename
		// 上传文件至指定目录
		c.SaveUploadedFile(file, dst)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"file_count": len(files),
	})
}
