package mymath_test

import (
	"Testing/laba2/mymath"
	"math"
	"testing"
)

func TestFactorial(t *testing.T) {
	testCases := []struct {
		inp     uint64
		out     float64
		message string
	}{
		{5.0, 120.0, "Simple number test"},
		{10.0, 3628800.0, "Big number test"},
		{0.0, 1.0, "Zero test"},
	}

	for _, data := range testCases {
		res := mymath.Factorial(data.inp)
		t.Logf("test: %v", data.message)
		if data.out != res {
			t.Errorf("Uncorrect value: get %v, but expected %v", res, data.out)
		}
	}
}

func TestPow(t *testing.T) {
	testCases := []struct {
		x       float64
		y       int
		out     float64
		message string
	}{
		{2.0, 2.0, 4.0, "2 in square"},
		{2.0, 10.0, 1024.0, "2^10"},
		{2.0, 0.0, 1.0, "2^0"},
		{2.0, -10, float64(1.0 / 1024.0), "2^-10"},
	}

	for _, data := range testCases {
		res := mymath.Pow(data.x, data.y)
		t.Logf("test: %v", data.message)
		if res != data.out {
			t.Errorf("Wrong return: returned %v, expected %v", res, data.out)
		}
	}
}

func TestPow2(t *testing.T) {
	testCases := []struct {
		inp     float64
		out     float64
		message string
	}{
		{100, 10000, "100^2"},
	}

	for _, data := range testCases {
		res := mymath.Pow2(data.inp)
		t.Log(data.message)
		if res != data.out {
			t.Errorf("Wrong return: returned %v, expected %v", res, data.out)
		}
	}
}

func TestSqrt(t *testing.T) {
	testCases := []struct {
		inp     float64
		out     float64
		message string
	}{
		{4, 2, "sqrt(4)"},
		{5, mymath.RoundFloat(math.Sqrt(5), 9), "sqrt(5)"},
		{-2, -1, "negative number"},
		{math.Inf(-1), -1, "negative number"},
		{math.NaN(), math.NaN(), "NaN"},
		{math.Inf(1), math.Inf(1), "Infinite"},
	}

	for _, data := range testCases {
		res := mymath.Sqrt(data.inp)
		t.Log(data.message)
		if !math.IsNaN(data.out) {
			if data.out != mymath.RoundFloat(res, 9) {
				t.Errorf("Wrong Ansvwer out is %v", res)
			}
		} else {
			if !math.IsNaN(res) {
				t.Errorf("Wrong Ansvwer out is %v", res)
			}
		}
	}
}

func TestAbs(t *testing.T) {
	testCases := []struct {
		inp     float64
		out     float64
		message string
	}{
		{-2, 2, "negative number"},
		{2, 2, "positive number"},
		{math.Inf(-1), math.Inf(1), "negative infinite number"},
		{math.NaN(), math.NaN(), "nan"},
	}

	for _, data := range testCases {
		t.Log(data.message)
		res := mymath.Abs(data.inp)
		if math.IsNaN(data.out) {
			if !math.IsNaN(res) {
				t.Errorf("Wrong Ansvwer out is %v", res)
			}
		} else {
			if data.out != res {
				t.Errorf("Wrong Ansvwer out is %v", res)
			}
		}
	}
}
