package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"Calculator-backend/internal/calculator"
)

type mockService struct {
	result float64
	err    error
}

func (m *mockService) Execute(op calculator.Operation, a float64, b *float64) (float64, error) {
	return m.result, m.err
}

func TestCalculateHandler_ValidRequest(t *testing.T) {
	h := NewCalculatorHandler(&mockService{result: 8})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate",
		bytes.NewBufferString(`{"operation":"add","a":5,"b":3}`))
	rec := httptest.NewRecorder()

	h.Calculate(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var resp calculateResponse
	json.Unmarshal(rec.Body.Bytes(), &resp)
	if resp.Result != 8 {
		t.Fatalf("expected 8, got %v", resp.Result)
	}
}

func TestCalculateHandler_BackendError(t *testing.T) {
	h := NewCalculatorHandler(&mockService{err: calculator.ErrDivisionByZero})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate",
		bytes.NewBufferString(`{"operation":"divide","a":1,"b":0}`))
	rec := httptest.NewRecorder()

	h.Calculate(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rec.Code)
	}
}

func TestCalculateHandler_MissingOperand(t *testing.T) {
	h := NewCalculatorHandler(&mockService{})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate",
		bytes.NewBufferString(`{"operation":"add","b":3}`))
	rec := httptest.NewRecorder()

	h.Calculate(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rec.Code)
	}
}
