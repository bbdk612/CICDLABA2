package laba2_test

import (
	"Testing/laba2"
	"Testing/laba2/mymath"
	"math"
	"testing"
)

func TestFunction1(t *testing.T) {
	testCases := []struct {
		inp     float64
		out     float64
		message string
	}{
		{0.37672, 0, "inp: 2.7649, out: 0"},
		{0, math.Inf(-1), "inp: 0, out: -inf"},
		{0.3807, 0.01, "inp: 0.3807, out: 0.01"},
	}

	for _, data := range testCases {
		res := laba2.Function1(data.inp)
		t.Log(data.message)
		if mymath.RoundFloat(res, 4) != mymath.RoundFloat(data.out, 4) {
			t.Errorf("wrong out: %v", mymath.RoundFloat(res, 4))
		}
	}
}

func TestFunction2(t *testing.T) {
	testCases := []struct {
		inp     float64
		out     float64
		message string
	}{
		{-0.417831, 0, "inp: -0.417831, out: 0"},
		{0, -1, "inp: 0, out: -1"},
		{-0.0419097, 1, "inp: -0.04068, out: 1"},
	}

	for _, data := range testCases {
		res := laba2.Function2(data.inp)
		t.Log(data.message)
		if mymath.RoundFloat(res, 2) != mymath.RoundFloat(data.out, 2) {
			t.Errorf("wrong out: %v", mymath.RoundFloat(res, 4))
		}
	}
}
