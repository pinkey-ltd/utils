package utils

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// hash defaultSal
const (
	defaultSal = "abcdefghijklmnopqrstuvwxyz"
)

// GetToken make Token
func GetToken(identity string) (*string, error) {
	claims := make(jwt.MapClaims)
	claims["Identity"] = identity
	// exp default 18 hour config from system setting
	exp := 18
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(exp)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString([]byte(defaultSal))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &tokenString, nil
}

// CheckToken auth token
func CheckToken(key string) bool {
	kv := strings.Split(key, " ")
	// Key type check. the Bearer token.
	if len(kv) != 2 || kv[0] != "Bearer" {
		log.Println("Token invalid:", key)
		return false
	}
	tokenString := kv[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(defaultSal), nil
	})
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	} else {
		log.Println("Token invalid:", err)
	}
	return false
}

// JWTAuthMiddleware MiddleWare
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"data": "",
				"msg":  "Authorization failed",
			})
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"data": "",
				"msg":  "Authorization failed",
			})
			c.Abort()
			return
		}

		if !CheckToken(authHeader) {
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"data": "",
				"msg":  "Authorization failed",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
