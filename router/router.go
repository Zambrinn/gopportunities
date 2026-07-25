package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default() // intialize router

	initializeRoutes(r)

	r.Run() // run the server
}
