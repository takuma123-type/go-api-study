package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DBTransactionMiddleware is a middleware for managing GORM transactions based on HTTP methods.
func DBTransactionMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// POST, PUT, DELETEメソッドの場合はトランザクションを開始
		if ctx.Request.Method == http.MethodPost ||
			ctx.Request.Method == http.MethodPut ||
			ctx.Request.Method == http.MethodDelete ||
			ctx.Request.Method == http.MethodPatch {
			log.Println("GORM Transaction started")
			tx := db.Begin()
			if tx.Error != nil {
				log.Printf("Failed to start transaction: %+v", tx.Error)
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Failed to start transaction",
				})
				return
			}
			// トランザクションをコンテキストに追加
			ctx.Set("db", tx)

			// リクエスト処理後にコミットまたはロールバックを行う
			defer func() {
				if len(ctx.Errors) > 0 {
					log.Println("Rolling back transaction due to errors")
					tx.Rollback()
				} else {
					if err := tx.Commit().Error; err != nil {
						log.Printf("Failed to commit transaction: %+v", err)
						ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
							"message": "Failed to commit transaction",
						})
					} else {
						log.Println("Transaction committed")
					}
				}
			}()
		} else {
			// 通常のデータベース接続をコンテキストに設定
			ctx.Set("db", db)
		}

		// 次のミドルウェアまたはハンドラに処理を渡す
		ctx.Next()
	}
}
