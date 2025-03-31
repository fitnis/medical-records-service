package services

import (
	"github.com/fitnis/record-service/models"
)

var chartNotes []models.ChartNote
var examRequests []models.ExamRequest
var examResults []models.ExamResult

func RecordChartNote(note models.ChartNote) models.GenericResponse {
	chartNotes = append(chartNotes, note)
	return models.GenericResponse{Message: "Chart updated"}
}

func GetChartNotes() []models.ChartNote {
	return chartNotes
}

func PerformExam(req models.ExamRequest) models.GenericResponse {
	examRequests = append(examRequests, req)
	return models.GenericResponse{Message: "Exam recorded"}
}

func RecordExamResult(res models.ExamResult) models.GenericResponse {
	examResults = append(examResults, res)
	return models.GenericResponse{Message: "Result recorded"}
}
