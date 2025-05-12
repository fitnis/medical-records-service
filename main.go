package main

import (
	"github.com/fitnis/record-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	rec := router.Group("/records")
	{
		handlers.RegisterRecordRoutes(rec)
	}

	router.Run(":8084")
}
