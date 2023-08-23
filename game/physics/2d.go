package physics

import "math"

type Point struct {
	X, Y float64
}

func P(x, y float64) Point {
	return Point{X: x, Y: y}
}

var Zero = Point{0, 0}

func (p Point) DistanceTo(q Point) float64 {
	return math.Sqrt(math.Pow(math.Abs(p.X-q.X), 2) + math.Pow(math.Abs(p.Y-q.Y), 2))
}

func (p Point) Rotate(degrees float64) Point {
	radians := degrees / (2 * math.Pi)
	return Point{
		X: p.X*math.Cos(radians) - p.Y*math.Sin(radians),
		Y: p.X*math.Sin(radians) + p.Y*math.Cos(radians),
	}
}
