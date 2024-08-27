package main

import (
	"GinDemo/db"
	"GinDemo/router"
)

func main() {
	router := router.SetupRouter()
	db.SetUp()

	_ = router.Run()
}
