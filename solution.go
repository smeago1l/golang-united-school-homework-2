package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type kek int

const (
	SidesTriangle kek = 3
	SidesSquare   kek = 4
	SidesCircle   kek = 0
)

func CalcSquare(sideLen float64, sidesNum kek) float64 {
	if sidesNum == SidesTriangle {
		return math.Sqrt(3) / 4 * math.Pow(sideLen, 2)
	}
	if sidesNum == SidesSquare {
		return math.Pow(sideLen, 2)
	}
	if sidesNum == SidesCircle {
		return math.Pi * math.Pow(sideLen, 2)
	}
	return 0
}
