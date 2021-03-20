package game

import (
	"sync"
)

type Attacker interface {
	Name() string
	Move() Position
	Position() Position
}

type zombie struct {
	name        string
	walkForward WalkStrategy
	position    *Position
	mu          sync.Mutex
}

func newZombie(name string) *zombie {
	return &zombie{
		name:        name,
		walkForward: MoveInStraightLine,
		// TODO make the position random on the first row
		position: NewPosition(0, 0),
	}
}

func (z *zombie) Name() string {
	z.mu.Lock()
	defer z.mu.Unlock()
	return z.name
}

func (z *zombie) Move() Position {
	z.mu.Lock()
	defer z.mu.Unlock()
	z.walkForward(z.position)
	return *z.position
}

func (z *zombie) Position() Position {
	z.mu.Lock()
	defer z.mu.Unlock()
	return *z.position
}
