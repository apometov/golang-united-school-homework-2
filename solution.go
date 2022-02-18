package square

import "math"

type sides uint16

const (
	SidesCircle   sides = 0
	SidesTriangle sides = 3
	SidesSquare   sides = 4
)

func CalcSquare(sideLen float64, sidesNum sides) float64 {
	switch sidesNum {
	case SidesCircle:
		return math.Pi * math.Pow(sideLen, 2)
	case SidesTriangle:
		return math.Sqrt(3) * math.Pow(sideLen, 2) / 4
	case SidesSquare:
		return math.Pow(sideLen, 2)
	default:
		return 0
	}
}
