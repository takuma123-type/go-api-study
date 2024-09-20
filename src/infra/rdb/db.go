package rdb

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDBFromContext(ctx *gin.Context) (*gorm.DB, error) {
	db, exists := ctx.Get("db")
	if !exists {
		return nil, errors.New("no database connection found in context")
	}
	return db.(*gorm.DB), nil
}
