package mymath_test

import (
	"Testing/laba2/mymath"
	"math"
	"testing"
)

func TestSin(t *testing.T) {
	testCases := []struct {
		input   float64
		out     float64
		message string
	}{
		{math.Pi, mymath.RoundFloat(math.Sin(math.Pi), 9), "Pi number"},
		{0, mymath.RoundFloat(math.Sin(0), 9), "Pi number"},
		{(math.Pi * 2 / 3), mymath.RoundFloat(math.Sin(math.Pi*2/3), 9), " 2 / 3 Pi number"},
		{(math.Pi / 2), mymath.RoundFloat(math.Sin(math.Pi/2), 9), " 0.5 Pi number"},
		{(math.Pi * 3 / 2), mymath.RoundFloat(math.Sin(math.Pi*3/2), 9), " 270 degrees number"},
		{(-(math.Pi * 3 / 2)), mymath.RoundFloat(math.Sin(-(math.Pi * 3 / 2)), 9), " -270 degrees number"},
	}

	for _, data := range testCases {
		res := mymath.Sin(data.input)
		if res != data.out {
			t.Errorf("Wrong answer: returned %v, but expected %v", res, data.out)
		}
		t.Logf("input %v, out %v, message %v", data.input, res, data.message)
	}
}

func TestCos(t *testing.T) {
	testCases := []struct {
		input   float64
		out     float64
		message string
	}{
		{math.Pi, mymath.RoundFloat(math.Cos(math.Pi), 9), "Pi number"},
		{0, mymath.RoundFloat(math.Cos(0), 9), "Pi number"},
		{(math.Pi * 2 / 3), mymath.RoundFloat(math.Cos(math.Pi*2/3), 9), " 2 / 3 Pi number"},
		{(math.Pi / 2), mymath.RoundFloat(math.Cos(math.Pi/2), 9), " 0.5 Pi number"},
		{(math.Pi * 3 / 2), mymath.RoundFloat(math.Cos(math.Pi*3/2), 9), " 270 degrees number"},
		{(-(math.Pi * 3 / 2)), mymath.RoundFloat(math.Cos(-(math.Pi * 3 / 2)), 9), " -270 degrees number"},
	}

	for _, data := range testCases {
		res := mymath.Cos(data.input)
		if res != data.out {
			t.Errorf("Wrong answer: returned %v, but expected %v", res, data.out)
		}
		t.Logf("input %v, out %v, message %v", data.input, data.out, data.message)
	}
}

func TestLn(t *testing.T) {
	testCases := []struct {
		inp     float64
		out     float64
		message string
	}{
		{1, 0, "1 test"},
		{0, math.Inf(-1), "0 test"},
		{5, mymath.RoundFloat(math.Log(5), 9), "simple number test"},
		{0.35022, mymath.RoundFloat(math.Log(0.35022), 9), "simple number test"},
		{math.Inf(1), math.Inf(1), "Inf test"},
	}

	for _, data := range testCases {
		res := mymath.Ln(data.inp)

		t.Log(data.message)

		if math.IsInf(res, -1) && data.inp != 0 {
			t.Errorf("something was wrong")
		} else if math.IsInf(res, 1) && !math.IsInf(data.out, 1) {
			t.Errorf("something was wrong")
		} else if math.IsNaN(res) && !math.IsNaN(data.inp) {
			t.Errorf("something was wrong")
		} else if mymath.RoundFloat(res, 3) != mymath.RoundFloat(data.out, 3) {
			t.Errorf("wrong answer expected %v, but returned %v", data.out, res)
		}
	}
}
