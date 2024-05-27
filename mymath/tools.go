package mymath

import "math"

func RoundFloat(number float64, digit int) float64 {
	p10 := Pow(10, digit)
	return (math.Round(number*p10) / p10)
}

func Factorial(num uint64) float64 {
	fact := float64(1)

	for i := float64(1); i <= float64(num); i++ {
		fact *= i
	}

	return float64(fact)
}

func Pow(x float64, y int) float64 {
	if y < 0 {
		x = float64(1.0 / x)
		y = -y
	}

	if y == 0 {
		return 1
	}

	if y == 1 {
		return x
	}

	if y%2 == 0 {
		return Pow(Pow2(x), y/2)
	} else {
		return x * Pow(Pow2(x), (y-1)/2)
	}
}

func Pow2(x float64) float64 {
	return x * x
}

func Sqrt(x float64) float64 {
	if x < 0 {
		return -1
	}

	if math.IsInf(x, 1) {
		return math.Inf(1)
	}

	if math.IsNaN(x) {
		return math.NaN()
	}
	mid := (x + 1) / 2
	eps := 1e-10
	for Abs(mid-x/mid) > eps {
		mid = (mid + x/mid) / 2
	}

	return mid
}

func Abs(x float64) float64 {
	if x < 0 {
		x *= -1
	}

	if math.IsNaN(x) {
		return math.NaN()
	}

	return x
}
