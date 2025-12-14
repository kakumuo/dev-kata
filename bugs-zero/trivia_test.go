package main

import (
	"fmt"
	"testing"
)

func TestMinPlayerCount(t *testing.T) {
	game := NewGame()

	game.Add("Player 1")

	if game.IsPlayable() {
		t.Fail()
	}
}

func TestMaxPlayerCount(t *testing.T) {
	game := NewGame()

	for i := range 7 {
		game.Add(fmt.Sprintf("Player %d", i+1))
	}

	if game.IsPlayable() {
		t.Fail()
	}
}

func TestRightAnswer(t *testing.T) {
	game := NewGame()

	game.Add("p1")
	game.Add("p2")

	game.Roll(1)
	game.WasCorrectlyAnswered()

	if game.penaltyStates[game.currentPlayer] != PSNone {
		t.Fail()
	}
}

func TestPenalty(t *testing.T) {
	game := NewGame()

	game.Add("p1")
	game.Add("p2")

	game.Roll(1)
	game.WrongAnswer()

	if game.penaltyStates[game.currentPlayer] != PSPenalty {
		t.Error("Expected to be in penalty state")
	}

	game.nextPlayer()
	game.nextPlayer()

	game.Roll(2)

	if game.penaltyStates[game.currentPlayer] != PSPenalty {
		t.Error("Expected to remain in penalty state after even die roll")
	}

	game.nextPlayer()
	game.nextPlayer()

	game.Roll(1)
	if game.penaltyStates[game.currentPlayer] != PSLeavingPenalty {
		t.Error("Expected to be in leaving penalty state after odd die roll")
	}

	game.WrongAnswer()
	if game.penaltyStates[game.currentPlayer] != PSPenalty {
		t.Error("Expected to remain in penalty after failed leaving penalty attempt")
	}

	game.nextPlayer()
	game.nextPlayer()

	game.Roll(1)
	game.WasCorrectlyAnswered()
	if game.penaltyStates[game.currentPlayer] != PSNone {
		t.Error("Expected to be in no penalty state after correct answer")
	}
}
