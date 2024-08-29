package main

// import (
// 	"GinDemo/db"
// 	"GinDemo/middleware"
// 	"GinDemo/user"
// 	"github.com/gin-gonic/gin"
// )

// 	func main() {
// 		gin := gin.Default()
// 		db.Init(gin)
// 		user.Init(gin)
// 		middleware.Init(gin)
// 		gin.Run("127.0.0.1:8080")
// 	}

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name" custom:"custom_value"`
	Age  int    `json:"age"`
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	t := reflect.TypeOf(p)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := reflect.ValueOf(p).Field(i).Interface()
		jsonTag := field.Tag.Get("json")
		customTag := field.Tag.Get("custom")
		fmt.Printf("Field: %s, JSON Tag: %s, , Value: %v, Custom Tag: %s\n", field.Name, jsonTag, value, customTag)
	}
}
