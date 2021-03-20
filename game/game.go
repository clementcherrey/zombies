package game

import (
	"time"
)

type EndResult string

const (
	Win  EndResult = "win"
	Lose EndResult = "lose"
)

type Playable interface {
	// launch a new attack with a fresh zombie
	Start()

	// check if a shot is successful
	ResolveShot(position *Position) bool

	// gets the attacker's (zombie) name
	AttackerName() string

	// resets the game before starting a new round
	Reset()
}

type Game struct {
	// the zombie attacking the wall
	attacker Attacker

	// how often the attacker move
	moveFrequency time.Duration

	// limits of the board
	xMax, yMax int

	// broadcast latest position of the zombie
	PositionCh chan<- Position

	// broadcast the end result of the game
	EndCh chan<- EndResult

	// used to reset the game
	resetCh chan struct{}
}

func New(positionCh chan<- Position, endCh chan<- EndResult) *Game {
	return &Game{
		attacker:      nil,
		moveFrequency: 1 * time.Second,
		xMax:          20,
		yMax:          20,
		PositionCh:    positionCh,
		EndCh:         endCh,
	}
}

func (g *Game) Start() {
	g.attacker = newZombie("white-king")
	g.resetCh = make(chan struct{})

	ticker := time.NewTicker(g.moveFrequency)
	for {
		select {
		case <-ticker.C:
			lastPosition := g.attacker.Move()
			g.PositionCh <- lastPosition
			if lastPosition.Y == g.yMax {
				g.EndCh <- Lose
			}
		case <-g.resetCh:
			return
		}
	}
}

func (g *Game) Reset() {
	close(g.resetCh)
	g.attacker = nil
}

func (g *Game) ResolveShot(shot *Position) bool {
	pos := g.attacker.Position()
	return pos.X == shot.X && pos.Y == shot.Y
}

func (g *Game) AttackerName() string {
	return g.attacker.Name()
}
