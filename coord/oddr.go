package coord

// OddR - Horisontal hex-coord representation
// Type Odd-R
type OddR struct {
	Q int
	R int
}

func (obj OddR) ToCube() Cube {
	x := obj.Q
	z := obj.R - (obj.Q - (obj.Q&1)) / 2

	return Cube{
		X: x,
		Y: -x-z,
		Z: z,
	}
}