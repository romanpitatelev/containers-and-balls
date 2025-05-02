package scannerservice

import "errors"

var (
	ErrInvalidN             = errors.New("n is out of acceptable range")
	ErrWrongRowLength       = errors.New("invalid row length")
	ErrInvalidNumberOfBalls = errors.New("number of balls is out of acceptable range")
)
