package physics

import "math"

type Point struct {
	X, Y float64
}

func P(x, y float64) Point {
	return Point{X: x, Y: y}
}

var Zero = Point{0, 0}

func (p Point) Add(q Point) Point {
	return P(p.X+q.X, p.Y+q.Y)
}

func (p Point) Sub(q Point) Point {
	return P(p.X-q.X, p.Y-q.Y)
}

func (p Point) Length() float64 {
	return math.Sqrt(math.Pow(p.X, 2) + math.Pow(p.Y, 2))
}

func (p Point) Normalize() Point {
	if p.X == 0 && p.Y == 0 {
		return p
	}
	return p.Scale(1 / p.Length())
}

func (p Point) Scale(k float64) Point {
	return P(p.X*k, p.Y*k)
}

func (p Point) DistanceTo(q Point) float64 {
	return math.Sqrt(math.Pow(math.Abs(p.X-q.X), 2) + math.Pow(math.Abs(p.Y-q.Y), 2))
}

func (p Point) Rotate(radians float64) Point {
	return Point{
		X: p.X*math.Cos(radians) - p.Y*math.Sin(radians),
		Y: p.X*math.Sin(radians) + p.Y*math.Cos(radians),
	}
}
