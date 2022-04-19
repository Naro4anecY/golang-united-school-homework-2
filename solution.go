package square

import (
	"math"
)

type Side uint8

const (
	SidesTriangle Side = 3
	SidesSquare   Side = 4
	SidesCircle   Side = 0
)

func CalcSquare(sideLen float64, sidesNum Side) float64 {
	var result float64
	switch sidesNum {
	case SidesSquare:
		result = math.Pow(sideLen, 2)
	case SidesTriangle:
		result = math.Pow(sideLen, 2) * math.Sqrt(3) / 4
	case SidesCircle:
		result = math.Pi * math.Pow(sideLen, 2)
	}
	return result
}
