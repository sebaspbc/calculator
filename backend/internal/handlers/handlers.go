package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"Calculator-backend/internal/calculator"
	"Calculator-backend/internal/service"
)

type CalculatorHandler struct {
	service service.CalculatorService
}

func NewCalculatorHandler(service service.CalculatorService) *CalculatorHandler {
	return &CalculatorHandler{service: service}
}

type calculateRequest struct {
	Operation string   `json:"operation"`
	A         *float64 `json:"a"`
	B         *float64 `json:"b"`
}

type calculateResponse struct {
	Result float64 `json:"result"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func (h *CalculatorHandler) Calculate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req calculateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON boy")
		return
	}

	if req.Operation == "" {
		writeError(w, http.StatusBadRequest, "Operation is required")
		return
	}

	if req.A == nil {
		writeError(w, http.StatusBadRequest, "Operand A is required and must be a number")
		return
	}

	result, err := h.service.Execute(calculator.Operation(req.Operation), *req.A, req.B)
	if err != nil {
		status := http.StatusBadRequest
		if errors.Is(err, calculator.ErrDivisionByZero) {
			status = http.StatusBadRequest
		}
		writeError(w, status, err.Error())
		return
	}

	json.NewEncoder(w).Encode(calculateResponse{Result: result})
}

func writeError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorResponse{Error: message})
}
