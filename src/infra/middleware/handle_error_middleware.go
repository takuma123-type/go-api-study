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
				smperr.HandleError(ctx, smperr.Internal("Internal Server Error"), http.StatusInternalServerError)
			}
		}()

		ctx.Next()

		if err := ctx.Errors.Last(); err != nil {
			switch e := err.Err.(type) {
			case *smperr.BadRequestErr:
				log.Printf("BadRequest ERROR: %+v", e.Trace())
				smperr.HandleError(ctx, e, e.Code())
			case *smperr.InternalErr:
				log.Printf("Internal ERROR: %+v", e.Trace())
				smperr.HandleError(ctx, e, e.Code())
			case *smperr.NotFoundErr:
				log.Printf("NotFound ERROR: %+v", e.Trace())
				smperr.HandleError(ctx, e, e.Code())
			default:
				log.Printf("Unknown ERROR: %+v", e)
				smperr.HandleError(ctx, smperr.Internal("Internal Server Error"), http.StatusInternalServerError)
			}
		}
	}
}
