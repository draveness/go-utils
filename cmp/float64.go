package cmp

import (
	"math"

	gocmp "github.com/google/go-cmp/cmp"
)

var float64Comparer = gocmp.Comparer(func(x, y float64) bool {
	delta := math.Abs(x - y)
	mean := math.Abs(x+y) / 2.0
	return delta/mean < 0.00001
})

func LessThan(x, y float64) bool { return x < y }

func LessThanOrEqual(x, y float64) bool {
	return x <= y || gocmp.Equal(x, y, float64Comparer)
}

func GreaterThan(x, y float64) bool { return x > y }

func GreaterThanOrEqual(x, y float64) bool {
	return x >= y || gocmp.Equal(x, y, float64Comparer)
}

func Equal(x, y float64) bool {
	return x == y || gocmp.Equal(x, y, float64Comparer)
}
