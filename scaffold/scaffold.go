// Package scaffold .
package scaffold

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Container[T any] struct {
	path  string
	model T
}

func (s Container[T]) Fac(db *gorm.DB, svr *gin.Engine) *gin.RouterGroup {
	r := svr.Group(s.path)
	{
		r.GET(s.path, FindByID(db, &s.model))

	}
	return r
}

// ProviderSet .
//var ProviderSet = wire.NewSet(NewScaffold)
