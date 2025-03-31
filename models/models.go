package models

type ChartNote struct {
	PatientID string `json:"patientId"`
	Note      string `json:"note"`
}

type ExamRequest struct {
	PatientID string `json:"patientId"`
	ExamType  string `json:"examType"`
}

type ExamResult struct {
	PatientID string `json:"patientId"`
	Result    string `json:"result"`
}

type GenericResponse struct {
	Message string `json:"message"`
}
