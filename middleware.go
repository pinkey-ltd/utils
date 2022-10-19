package utils

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Cors() gin.HandlerFunc {
	// return func(c *gin.Context) {
	// 	method := c.Request.Method

	// 	c.Header("Access-Control-Allow-Origin", "*")
	// 	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	// 	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	// 	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	// 	c.Header("Access-Control-Allow-Credentials", "true")

	// 	//放行所有OPTIONS方法
	// 	if method == "OPTIONS" {
	// 		c.AbortWithStatus(http.StatusNoContent)
	// 	}
	// 	// 处理请求
	// 	c.Next()
	// }
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "AccessToken", "X-CSRF-Token", "Authorization", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

// seeeions middleware with go-orm
func SessionsGoRM(db *gorm.DB) gin.HandlerFunc {
	store := gormsessions.NewStore(db, true, []byte("secret"))
	return sessions.Sessions("mysession", store)
}

// gzip middleware
func Gzip() gin.HandlerFunc {
	return gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedExtensions([]string{".pdf", ".mp4"}))
}

// pprof debug
func Pprof(r *gin.Engine, paths ...string) {
	path := "dev/pprof"
	if len(paths) > 0 {
		path = paths[0]
	}
	pprof.Register(r, path)
}
