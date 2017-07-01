package main

import (
	"github.com/gin-gonic/gin"
	"github.com/catman/go_api/comments"
	"github.com/catman/go_api/users"
)


func RegistersRoutes() {
	r := gin.Default()

	//commentHandler
	r.GET("/comments", comments.Handler)

	//usersHandler
	r.GET("/users", users.Handler)

	r.Run(":30000")
}


//r.GET("/", func(c *gin.Context) {
//    var product mds.Product
//    product.Code = "L1212"
//    samsam := db.First(&product)
//    fmt.Println(samsam)
//    c.String(200, "Hello world!")
//})