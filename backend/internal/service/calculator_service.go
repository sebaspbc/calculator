package service

import "Calculator-backend/internal/calculator"

type CalculatorService interface {
	Execute(op calculator.Operation, a float64, b *float64) (float64, error)
}

type calculatorService struct{}

func NewCalculatorService() CalculatorService {
	return &calculatorService{}
}

func (s *calculatorService) Execute(op calculator.Operation, a float64, b *float64) (float64, error) {
	return calculator.Calculate(op, a, b)
}
