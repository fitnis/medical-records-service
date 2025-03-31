package handlers

import (
	"net/http"

	"github.com/fitnis/record-service/models"
	"github.com/fitnis/record-service/services"
	"github.com/gin-gonic/gin"
)

// curl -X POST -H "Content-Type: application/json" -d '{"patientId":"123","note":"Patient is stable."}' http://localhost:8084/records/chart
func RecordChartNote(c *gin.Context) {
	var note models.ChartNote
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	c.JSON(http.StatusCreated, services.RecordChartNote(note))
}

// curl http://localhost:8084/records/chart
func GetChartNotes(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetChartNotes())
}

// curl -X POST -H "Content-Type: application/json" -d '{"patientId":"123","examType":"CT Scan"}' http://localhost:8084/records/exam
func PerformExam(c *gin.Context) {
	var req models.ExamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	c.JSON(http.StatusCreated, services.PerformExam(req))
}

// curl -X POST -H "Content-Type: application/json" -d '{"patientId":"123","result":"All clear"}' http://localhost:8084/records/exam/results
func RecordExamResult(c *gin.Context) {
	var res models.ExamResult
	if err := c.ShouldBindJSON(&res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	c.JSON(http.StatusCreated, services.RecordExamResult(res))
}
