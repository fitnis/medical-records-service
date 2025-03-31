package main

import (
	"github.com/fitnis/record-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	rec := router.Group("/records")
	{
		rec.POST("/chart", handlers.RecordChartNote)
		rec.GET("/chart", handlers.GetChartNotes)
		rec.POST("/exam", handlers.PerformExam)
		rec.POST("/exam/results", handlers.RecordExamResult)
	}

	router.Run(":8084")
}
