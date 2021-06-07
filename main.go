package main

import (
	"github.com/gin-contrib/cors"
	"shares/db"
	"shares/router"
)

func main() {
	db.Init()
	r := router.NewRouter()
	r.Use(cors.Default())
	r.Run(":3000")
}
