package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/go-api-study/src/support/smperr"
)

func HandleErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("HandleErrorMiddleware is called")

		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic: %+v", r)
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Internal Server Error",
				})
			}
		}()

		ctx.Next()

		if err := ctx.Errors.Last(); err != nil {
			switch e := err.Err.(type) {
			case *smperr.BadRequestErr:
				log.Printf("BadRequest ERROR: %+v", e.Trace())
				ctx.AbortWithStatusJSON(e.Code(), gin.H{
					"message": e.Msg(),
				})
			case *smperr.InternalErr:
				log.Printf("Internal ERROR: %+v", e.Trace())
				ctx.AbortWithStatusJSON(e.Code(), gin.H{
					"message": e.Msg(),
				})
			case *smperr.NotFoundErr:
				log.Printf("NotFound ERROR: %+v", e.Trace())
				ctx.AbortWithStatusJSON(e.Code(), gin.H{
					"message": e.Msg(),
				})
			case *smperr.UnauthorizedErr:
				log.Printf("Unauthorized ERROR: %+v", e.Trace())
				ctx.AbortWithStatusJSON(e.Code(), gin.H{
					"message": e.Msg(),
				})
			default:
				log.Printf("Unknown ERROR: %+v", e)
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Internal Server Error",
				})
			}
		}
	}
}
