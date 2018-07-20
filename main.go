package main

import (
	"github.com/gin-gonic/gin"
	"web/router"
)


func main() {
	r := gin.Default()



	private := r.Group("/api")
	private = router.SetRouter(private)

	r.Run()
}
