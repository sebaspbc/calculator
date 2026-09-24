package calculator

import (
	"errors"
	"math"
)

var (
	ErrDivisionByZero   = errors.New("division by zero is not allowed")
	ErrNegativeSqrt     = errors.New("cannot calculate square root of a negative number")
	ErrUnknownOperation = errors.New("unknown operation")
	ErrMissingOperandB  = errors.New("operand 'b' is required for this operation")
)

type Operation string

const (
	Add        Operation = "add"
	Subtract   Operation = "subtract"
	Multiply   Operation = "multiply"
	Divide     Operation = "divide"
	Power      Operation = "power"
	SquareRoot Operation = "sqrt"
	Percentage Operation = "percentage"
)

var binaryOps = map[Operation]bool{
	Add: true, Subtract: true, Multiply: true, Divide: true, Power: true,
}

func Calculate(op Operation, a float64, b *float64) (float64, error) {
	if binaryOps[op] && b == nil {
		return 0, ErrMissingOperandB
	}

	switch op {
	case Add:
		return a + *b, nil
	case Subtract:
		return a - *b, nil
	case Multiply:
		return a * *b, nil
	case Divide:
		if *b == 0 {
			return 0, ErrDivisionByZero
		}
		return a / *b, nil
	case Power:
		return math.Pow(a, *b), nil
	case SquareRoot:
		if a < 0 {
			return 0, ErrNegativeSqrt
		}
		return math.Sqrt(a), nil
	case Percentage:
		return a / 100, nil
	default:
		return 0, ErrUnknownOperation
	}
}
