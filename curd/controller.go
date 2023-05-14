package curd

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

// ScaffoldList .
func ScaffoldList[T any](db *gorm.DB, obs []*T) gin.HandlerFunc {

	return func(c *gin.Context) {
		err := db.Order("id").Find(&obs).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 404,
				"data": "",
				"msg":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, obs)
	}
}

// ScaffoldListWhitChildren .
func ScaffoldListWhitChildren[T any](db *gorm.DB, obs []*T, level int) gin.HandlerFunc {

	return func(c *gin.Context) {
		// resp, err := IListWithChildren(ob, level)
		if level < 1 {
			level = 1
		}
		preload := "Children"
		for l := 1; l < level; l++ {
			preload += ".Children"
		}
		err := db.Preload(preload).Find(&obs).Error

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 404,
				"data": "",
				"msg":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, obs)
	}
}

// ScaffoldFindByID .
func ScaffoldFindByID[T any](db *gorm.DB, ob *T) gin.HandlerFunc {

	return func(c *gin.Context) {
		id := c.Param("id")
		err := db.Where("id = ?", id).First(&ob).Error

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 404,
				"data": "",
				"msg":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, ob)
	}
}

// ScaffoldUpdate .
func ScaffoldUpdate[T any](db *gorm.DB, ob *T) gin.HandlerFunc {

	return func(c *gin.Context) {
		if err := c.ShouldBindBodyWith(ob, binding.JSON); err == nil {
			if errDb := db.Save(&ob).Error; errDb == nil {
				c.JSON(http.StatusOK, ob)
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

// ScaffoldCreate .
func ScaffoldCreate[T any](db *gorm.DB, ob *T) gin.HandlerFunc {

	return func(c *gin.Context) {
		if err := c.ShouldBindBodyWith(ob, binding.JSON); err == nil {
			if errDb := db.Omit("id").Create(&ob).Error; errDb == nil {
				c.JSON(http.StatusOK, ob)
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

// ScaffoldRemove .
func ScaffoldRemove[T any](db *gorm.DB, ob *T) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		// res, err := IRemove(ob, id)
		err := db.Delete(ob, id).Error
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": "Deleted successfully!"})
	}
}
