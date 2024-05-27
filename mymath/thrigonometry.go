package mymath

import (
	"math"
)

func Sin(x float64) float64 {
	var result float64 = 0
	var reverse float64 = 1
	if math.IsInf(x, -1) {
		return math.Inf(-1)
	}

	if math.IsInf(x, 1) {
		return math.Inf(1)
	}
	if x < 0 {
		reverse = -1
		x *= -1
	}

	if x == 0 {
		return 0
	}

	for x > math.Pi {
		x -= math.Pi
		if reverse < 0 {
			reverse = 1
		} else {
			reverse = -1
		}
	}

	sign := 1
	unit := float64(1)
	i := 0
	for i < 50 {
		f := Factorial(uint64(2*i + 1))

		p := Pow(x, 2*i+1)

		unit = float64(p / f)
		result += float64(sign) * unit

		sign = -sign
		i += 1
	}

	return RoundFloat(result*reverse, 9)
}

func Cos(x float64) float64 {
	var result float64 = 0
	var reverse float64 = 1
	if math.IsInf(x, -1) {
		return math.Inf(-1)
	}

	if math.IsInf(x, 1) {
		return math.Inf(1)
	}

	if x < 0 {
		x *= -1
	}
	if x == 0 {
		return 1
	}
	for x > (2 * math.Pi) {
		x -= math.Pi
	}

	sign := 1
	i := 0
	unit := float64(1)
	for i < 50 {
		f := Factorial(uint64(2 * i))

		p := Pow(x, 2*i)

		unit = float64(p / f)
		result += float64(sign) * unit

		sign = -sign
		i++
	}

	return RoundFloat(result*reverse, 9)
}

func Ln(x float64) float64 {
	// Returns natural logarithm of number
	// Calculates it with Taylor series
	var result float64 = 0

	if x <= 0 {
		return math.Inf(-1)
	}

	if math.IsInf(x, 1) {
		return math.Inf(1)
	}

	if math.IsNaN(x) {
		return math.NaN()
	}

	if x > 1 {
		return Ln(x/2) - Ln(0.5)
	}

	var unit float64 = 1
	var sign float64 = 1
	var i int = 1
	for i < 1000 {
		p := Pow((x - 1), i)
		unit = p / float64(i)
		result += sign * unit

		sign = -sign
		i++
	}

	return RoundFloat(result, 9)
}
