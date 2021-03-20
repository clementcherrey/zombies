package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"zombies/game"
)

type player struct {
	name   string
	points int
}

type session struct {
	r            *bufio.Reader
	w            *bufio.Writer
	p            *player
	g            game.Playable
	lastPosition chan game.Position
	gameResult   chan game.EndResult
	done         chan struct{}
}

func NewSession(conn net.Conn) *session {
	lastPosition := make(chan game.Position)
	gameResult := make(chan game.EndResult)

	return &session{
		r:            bufio.NewReader(conn),
		w:            bufio.NewWriter(conn),
		p:            &player{},
		g:            game.New(lastPosition, gameResult),
		lastPosition: lastPosition,
		gameResult:   gameResult,
		done:         make(chan struct{}),
	}
}

func (s *session) readMessages() {
	defer close(s.done)

	for {
		msg, err := s.r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		fields := strings.Fields(msg)
		// TODO check the message format before processing it

		switch fields[0] {
		case "START":
			s.p.name = fields[1]
			go s.g.Start()
		case "SHOOT":
			if err := s.handleShot(fields[1:]); err != nil {
				log.Println(err)
				return
			}
		default:
			log.Printf("invalid action %q", fields[0])
			return
		}
	}
}

func (s *session) writeUpdates() {
	defer close(s.done)

	for {
		select {
		case pos := <-s.lastPosition:
			msg := fmt.Sprintf("WALK %s %d %d\n", s.g.AttackerName(),
				pos.X, pos.Y)
			if err := send(s.w, msg); err != nil {
				log.Println(err)
				return
			}
		case gameResult := <-s.gameResult:
			msg := fmt.Sprintf("YOU %s\n", gameResult)
			if err := send(s.w, msg); err != nil {
				log.Println(err)
				return
			}
			s.g.Reset()
		}
	}
}

func (s *session) handleShot(fields []string) error {
	shotHit := s.g.ResolveShot(newShot(fields))
	if shotHit {
		s.p.points++
	}

	// write response message
	var endOfMsg string
	if shotHit {
		endOfMsg = s.g.AttackerName()
	}
	msg := fmt.Sprintf("BOOM %s %d %s\n", s.p.name, s.p.points, endOfMsg)
	if err := send(s.w, msg); err != nil {
		return err
	}

	if shotHit {
		s.gameResult <- game.Win
	}
	return nil
}

func send(w *bufio.Writer, msg string) error {
	if _, err := w.WriteString(msg); err != nil {
		return err
	}
	return w.Flush()
}

func newShot(fields []string) *game.Position {
	x, _ := strconv.Atoi(fields[0])
	y, _ := strconv.Atoi(fields[1])
	return &game.Position{
		X: x,
		Y: y,
	}
}
