package game

type Position struct {
	X int
	Y int
}

func NewPosition(x, y int) *Position {
	return &Position{
		X: x,
		Y: y,
	}
}

// way of walking
type WalkStrategy func(*Position)

func MoveInStraightLine(p *Position) {
	p.Y += 1
}

// TODO add more WalkStrategy
