package main

import (
	"shares/db"
	"shares/router"
)

func main(){
	db.Init()
	r:=router.NewRouter()
	r.Run(":3000")
}