package utils

import (
	"encoding/json"
	"net/http"
)

func SendSuccessResponse(w http.ResponseWriter, data interface{}) {
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "success", "data": data})
}

func SendErrorResponse(w http.ResponseWriter, code int, message string) {
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "error", "message": message})
}
