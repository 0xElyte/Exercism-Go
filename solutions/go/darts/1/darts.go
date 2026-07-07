package darts

func Score(x, y float64) int {
	d2 := x*x + y*y

	switch {
	case d2 > 100:
		return 0
	case d2 > 25:
		return 1
	case d2 > 1:
		return 5
	default:
		return 10
	}
}