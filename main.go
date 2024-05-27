package laba2

import (
	"Testing/laba2/mymath"
)

func Function1(x float64) float64 {
	res := mymath.Pow2(mymath.Sin(x)) + mymath.Pow2(mymath.Cos(x)) + mymath.Ln(mymath.Sin(x))
	return res
}

func Function2(x float64) float64 {
	sin := mymath.Sin(mymath.Pow2(x))
	cos := mymath.Cos(mymath.Ln(mymath.Pow2(x)))
	sq := mymath.Sqrt(sin + cos)

	return sq
}

func F(x float64) float64 {
	if x < 0 {
		return Function1(x)
	} else {
		return Function2(x)
	}
}
