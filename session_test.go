package main

import (
	"bufio"
	"bytes"
	"testing"

	"zombies/game"
)

func TestSession_handleShot(t *testing.T) {
	buff := bytes.Buffer{}

	s := session{
		w: bufio.NewWriter(&buff),
		p: &player{
			name:   "bob",
			points: 12,
		},
		g: &NilPlayable{},
	}

	if err := s.handleShot([]string{"1", "2"}); err != nil {
		t.Fatalf("handleShot return unexpected errorf: %v", err)
	}

	response, err := buff.ReadString('\n')
	if err != nil {
		t.Fatalf("fail to read from buffer %v", err)
	}
	if response != "BOOM bob 12 \n" {
		t.Errorf("response message invalid: %s ", response)
	}
}

// NilPlayable is an empty implementation of game.Playable interface
// for test purpose
type NilPlayable struct{}

func (*NilPlayable) Start()                          {}
func (*NilPlayable) ResolveShot(*game.Position) bool { return false }
func (*NilPlayable) AttackerName() string            { return "" }
func (*NilPlayable) Reset()                          {}
