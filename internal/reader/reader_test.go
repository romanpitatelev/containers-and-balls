package reader

import (
	"errors"
	"testing"
)

func TestValidateNumber(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		wantErr error
	}{
		{
			name:    "valid minimum n",
			input:   minN,
			wantErr: nil,
		},
		{
			name:    "valid maximum n",
			input:   maxN,
			wantErr: nil,
		},
		{
			name:    "n below min",
			input:   minN - 1,
			wantErr: ErrInvalidN,
		},
		{
			name:    "n above max",
			input:   maxN + 1,
			wantErr: ErrInvalidN,
		},
		{
			name:    "negative n",
			input:   -10,
			wantErr: ErrInvalidN,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateNumber(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("validateNumber(%d) = %v, want %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestValidateContainersBalls(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		wantErr error
	}{
		{
			name:    "valid minimum balls",
			input:   minNumberOfBalls,
			wantErr: nil,
		},
		{
			name:    "valid maximum balls",
			input:   maxNumberOfBalls,
			wantErr: nil,
		},
		{
			name:    "balls below minimum",
			input:   minNumberOfBalls - 1,
			wantErr: ErrInvalidNumberOfBalls,
		},
		{
			name:    "balls above maximum",
			input:   maxNumberOfBalls + 1,
			wantErr: ErrInvalidNumberOfBalls,
		},
		{
			name:    "negative balls",
			input:   -100,
			wantErr: ErrInvalidNumberOfBalls,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateContainersBalls(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("validateContainersBalls(%d) = %v, want %v", tt.input, err, tt.wantErr)
			}
		})
	}
}
