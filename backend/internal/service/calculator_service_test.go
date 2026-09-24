package service

import (
	"testing"

	"Calculator-backend/internal/calculator"
)

func TestCalculatorService_Execute(t *testing.T) {
	service := NewCalculatorService()
	b := 4.0

	result, err := service.Execute(calculator.Add, 2, &b)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 6 {
		t.Errorf("expected 6, got %v", result)
	}
}
