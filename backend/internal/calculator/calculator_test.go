package calculator

import (
	"errors"
	"math"
	"testing"
)

func f(v float64) *float64 { return &v }

func TestCalculate(t *testing.T) {
	tests := []struct {
		name    string
		op      Operation
		a       float64
		b       *float64
		want    float64
		wantErr error
	}{
		{"addition", Add, 2, f(3), 5, nil},
		{"subtraction", Subtract, 10, f(4), 6, nil},
		{"multiplication", Multiply, 3, f(3), 9, nil},
		{"division", Divide, 10, f(2), 5, nil},
		{"division by zero", Divide, 10, f(0), 0, ErrDivisionByZero},
		{"power", Power, 2, f(10), 1024, nil},
		{"sqrt", SquareRoot, 9, nil, 3, nil},
		{"sqrt negative", SquareRoot, -4, nil, 0, ErrNegativeSqrt},
		{"percentage", Percentage, 50, nil, 0.5, nil},
		{"missing operand", Add, 1, nil, 0, ErrMissingOperandB},
		{"unknown op", Operation("modulo"), 1, f(1), 0, ErrUnknownOperation},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.op, tt.a, tt.b)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if math.Abs(got-tt.want) > 1e-9 {
				t.Fatalf("expected %v, got %v", tt.want, got)
			}
		})
	}
}
