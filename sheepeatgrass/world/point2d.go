package world

/**
 * Point2D
 */
type Point2D struct {
	X int
	Y int
}

func (p *Point2D) Move(dir Direction) {
	switch dir {
	case East:
		p.X = p.X + 1
	case West:
		p.Y = p.X - 1
	case North:
		p.Y = p.Y + 1
	case South:
		p.Y = p.Y - 1
	}
}
