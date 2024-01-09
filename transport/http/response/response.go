package response

import (
	"encoding/json"
	"net/http"

	"github.com/vinbyte/mzk/shared/logger"
)

// Base is the base object of all responses
type Base struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// WithMessage sends a response with a simple text message
func WithMessage(w http.ResponseWriter, code int, message string) {
	respond(w, code, Base{Msg: message})
}

// WithPreparingShutdown sends a default response for when the server is preparing to shut down
func WithPreparingShutdown(w http.ResponseWriter) {
	WithMessage(w, http.StatusServiceUnavailable, "SERVER PREPARING TO SHUT DOWN")
}

// WithUnhealthy sends a default response for when the server is unhealthy
func WithUnhealthy(w http.ResponseWriter) {
	WithMessage(w, http.StatusServiceUnavailable, "SERVER UNHEALTHY")
}

func respond(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		logger.ErrorWithStack(err)
	}
}
