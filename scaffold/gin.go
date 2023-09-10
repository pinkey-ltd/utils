package scaffold

import (
	"gorm.io/gorm"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// List .
func List[T any](db *gorm.DB, ob []*T) gin.HandlerFunc {

	return func(c *gin.Context) {
		resp, err := IList(db, ob)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 404,
				"data": "",
				"msg":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}

// ListByOrder .
func ListByOrder[T any](db *gorm.DB, ob []*T) gin.HandlerFunc {

	return func(c *gin.Context) {
		resp, err := IListByOrder(db, ob)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 404,
				"data": "",
				"msg":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}

// ListWhitChildren .
func ListWhitChildren[T any](db *gorm.DB, ob []*T, level int) gin.HandlerFunc {

	return func(c *gin.Context) {
		resp, err := IListWithChildren(db, ob, level)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 404,
				"data": "",
				"msg":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}

// FindByID .
func FindByID[T any](db *gorm.DB, ob *T) gin.HandlerFunc {

	return func(c *gin.Context) {
		id := c.Param("id")
		resp, err := IFindByID(db, ob, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 404,
				"data": "",
				"msg":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}

// Update .
func Update[T any](db *gorm.DB, ob *T) gin.HandlerFunc {

	return func(c *gin.Context) {
		if err := c.ShouldBindBodyWith(ob, binding.JSON); err == nil {
			if s, errDb := IUpdate(db, ob); errDb == nil {
				c.JSON(http.StatusOK, s)
				return
			} else {
				log.Println("Update error: ", err)
				c.JSON(http.StatusBadRequest, gin.H{
					"code": 404,
					"data": "",
					"msg":  errDb.Error(),
				})
				return
			}
		} else {
			log.Println("Read request error: ", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 404,
				"data": "",
				"msg":  err.Error(),
			})
		}
	}
}

// Create .
func Create[T any](db *gorm.DB, ob *T) gin.HandlerFunc {

	return func(c *gin.Context) {
		if err := c.ShouldBindBodyWith(ob, binding.JSON); err == nil {
			if s, errDb := ICrerate(db, ob); errDb == nil {
				c.JSON(http.StatusOK, s)
				return
			} else {
				log.Println("Create error: ", errDb)
				c.JSON(http.StatusBadRequest, gin.H{
					"code": 404,
					"data": "",
					"msg":  errDb.Error(),
				})
				return
			}
		} else {
			log.Println("Read request error: ", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 404,
				"data": "",
				"msg":  err.Error(),
			})
		}
	}
}

// Remove .
func Remove[T any](db *gorm.DB, ob *T) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		res, err := IRemove(db, ob, id)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, res)
			return
		}
		c.JSON(http.StatusOK, res)
	}
}
