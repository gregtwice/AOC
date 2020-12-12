package helper

type Point struct {
	X int
	Y int
}

func (p Point) ManDist(point Point) int {
	return (Abs(p.X-point.X) + Abs(p.Y-point.Y))
}

func (p *Point) Rotate(angle int) {
	copy := Point{
		X: p.X,
		Y: p.Y,
	}
	switch angle {
	case 90:
		p.X = copy.Y
		p.Y = -copy.X
	case 180:
		p.X = -copy.X
		p.Y = -copy.Y
	case 270:
		p.X = -copy.Y
		p.Y = copy.X
	case -90:
		p.X = -copy.Y
		p.Y = copy.X
	case -180:
		p.X = -copy.X
		p.Y = -copy.Y
	case -270:
		p.X = copy.Y
		p.Y = -copy.X
	}
}
