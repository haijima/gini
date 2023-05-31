// Package gini is a simple implementation of calculating Gini coefficient.
package gini

import (
	"errors"
	"sort"
)

type number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

// NegativeValueError is returned when some elements of a slice are negative.
var NegativeValueError = errors.New("all values cannot be negative")

// Gini returns Gini coefficient.
func Gini[T number](values []T) (float64, error) {
	n := len(values)
	if n == 0 {
		return 0, nil
	}

	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })

	if values[0] < 0 {
		return 0, NegativeValueError
	}

	var g, s T
	for i, v := range values {
		g += T(i) * v
		s += v
	}
	g *= 2
	g -= T(n-1) * s
	return float64(g) / float64(T(n)*s), nil
}
