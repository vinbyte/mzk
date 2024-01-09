package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vinbyte/mzk/internal/domain/student"
	"github.com/vinbyte/mzk/shared/failure"
	"github.com/vinbyte/mzk/transport/http/response"
)

// StudentHandler is the HTTP handler for student domain.
type StudentHandler struct {
	StudentService student.StudentService
}

// ProvideStudentHandler is the provider for this handler.
func ProvideStudentHandler(studentService student.StudentService) StudentHandler {
	return StudentHandler{
		StudentService: studentService,
	}
}

// Router sets up the router for this domain.
func (h *StudentHandler) Router(r chi.Router) {
	r.Post("/records", h.FetchRecordsByFilter)
}

// FetchRecordsByFilter resolves a records by filter.
// @Summary Resolve records by filter
// @Description This endpoint resolves a records by filter.
// @Tags records
// @Param request body student.RecordFilter true "Request Body"
// @Produce json
// @Success 200 {object} response.ExampleSuccessResponse{Records=[]student.Record}
// @Failure 400 {object} response.ExampleBadRequestResponse
// @Failure 500 {object} response.ExampleInternalErrorResponse
// @Router /v1/records [post]
func (h *StudentHandler) FetchRecordsByFilter(w http.ResponseWriter, r *http.Request) {
	reqBody := student.RecordFilter{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqBody)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	err = reqBody.Validate()
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	recordData, err := h.StudentService.GetRecords(r.Context(), reqBody)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, recordData)
}
