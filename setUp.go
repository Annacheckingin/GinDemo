package main

import (
	"GinDemo/db"
	"GinDemo/user"
)

func SetUp() {
	user.SetUp()
	db.Setup()
}
