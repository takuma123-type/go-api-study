package main

import (
	"github.com/takuma123-type/go-api-study/src/infra/router"
)

func main() {
	r := router.InitRouter()
	r.Run(":9090")
}
