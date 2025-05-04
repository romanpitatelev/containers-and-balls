package reader

import (
	"errors"
	"io"
	"os"
	"strconv"
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

//nolint:funlen
func TestReadContainersBalls(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		n       int
		wantErr error
	}{
		{
			name:    "valid input",
			input:   "1 2\n3 4\n",
			n:       2,
			wantErr: nil,
		},
		{
			name:    "valid 3x3 input",
			input:   "1 0 0\n0 1 0\n0 0 1\n",
			n:       3,
			wantErr: nil,
		},
		{
			name:    "row with wrong length",
			input:   "1 2 3\n4 5\n",
			n:       3,
			wantErr: ErrWrongRowLength,
		},
		{
			name:    "negative number of balls",
			input:   "1 -1\n2 3\n",
			n:       2,
			wantErr: ErrInvalidNumberOfBalls,
		},
		{
			name:    "invalid number of balls (too large)",
			input:   "1 1000000001\n2 3\n",
			n:       2,
			wantErr: ErrInvalidNumberOfBalls,
		},
		{
			name:    "non-numeric input",
			input:   "1 abc\n2 3\n",
			n:       2,
			wantErr: strconv.ErrSyntax,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("failed to create pipe: %v", err)
			}
			defer r.Close() //nolint:errcheck

			go func() {
				defer w.Close() //nolint:errcheck

				_, err := io.WriteString(w, tt.input)
				if err != nil {
					t.Errorf("failed to write to pipe: %v", err)
				}
			}()

			oldStdin := os.Stdin

			defer func() {
				os.Stdin = oldStdin
			}()

			os.Stdin = r

			_, _, err = ReadContainersBalls(tt.n)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("ReadContainersBalls() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
		})
	}
}
