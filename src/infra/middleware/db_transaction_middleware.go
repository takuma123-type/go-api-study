package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DBTransactionMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodPost ||
			ctx.Request.Method == http.MethodPut ||
			ctx.Request.Method == http.MethodDelete ||
			ctx.Request.Method == http.MethodPatch {
			tx := db.Begin()
			ctx.Set("db", tx)

			defer func() {
				if len(ctx.Errors) > 0 {
					log.Println("Rolling back transaction due to errors")
					if err := tx.Rollback().Error; err != nil {
						log.Printf("Failed to rollback transaction: %v", err)
					}
				} else {
					if err := tx.Commit().Error; err != nil {
						ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
							"message": "Failed to commit transaction",
						})
					}
				}
			}()
		} else {
			ctx.Set("db", db)
		}

		ctx.Next()
	}
}
