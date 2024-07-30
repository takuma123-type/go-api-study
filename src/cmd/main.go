package main

import (
	"github.com/takuma123-type/go-api-study/infra/router"
)

func main() {
	r := router.InitRouter()
	r.Run(":9090")
}
