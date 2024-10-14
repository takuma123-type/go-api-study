package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/takuma123-type/go-api-study/src/infra/middleware"
	"github.com/takuma123-type/go-api-study/src/infra/router"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:password@tcp(db:3306)/golang-db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	g := gin.Default()
	g.Use(middleware.DBTransactionMiddleware(db))
	g.Use(middleware.HandleErrorMiddleware())
	router.NewUserRouter(g)
	router.NewMentorRouter(g)
	router.NewPlanRouter(g)
	router.NewContractRequestRouter(g)
	log.Fatal(g.Run(":9090"))
}
