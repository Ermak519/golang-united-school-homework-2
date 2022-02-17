package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type NumberOFSides int

const SidesTriangle NumberOFSides = 3
const SidesSquare NumberOFSides = 4
const SidesCircle NumberOFSides = 0

func CalcSquare(sideLen float64, sidesNum NumberOFSides) float64 {

	pi := math.Pi

	switch {
	case sidesNum == SidesTriangle:
		return math.Pow(sideLen, 2) * (math.Sqrt(3) / 4)
	case sidesNum == SidesSquare:
		return math.Pow(sideLen, 2)
	case sidesNum == SidesCircle:
		return pi * math.Pow(sideLen, 2)
	}
	return 0
}
