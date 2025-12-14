package main

import (
	"fmt"
	"math/rand"
)

type CategoryType int

const (
	CTPop CategoryType = iota
	CTScience
	CTSports
	CTRock
)

type Question struct {
	question string
	answer   int // 1-4
}

type PenaltyState int

const (
	PSNone PenaltyState = iota
	PSLeavingPenalty
	PSPenalty
)

type Game struct {
	players       []string
	places        []int
	purses        []int
	penaltyStates []PenaltyState

	questionMap     map[CategoryType][]Question
	questionIndexes []int

	currentPlayer int
}

var CATEGORY_TYPE_MAP = map[CategoryType]string{CTPop: "Pop", CTRock: "Rock", CTScience: "Science", CTSports: "Sports"}

const MIN_PLAYERS, MAX_PLAYERS = 2, 6
const NUM_QUESTIONS = 50
const NUM_PLACES = 20
const WINNER_COIN_COUNT = 6

func NewGame() *Game {
	game := &Game{}
	game.places = make([]int, 0)

	game.questionMap = make(map[CategoryType][]Question)
	game.questionIndexes = make([]int, len(CATEGORY_TYPE_MAP))
	game.questionMap[CTPop] = make([]Question, 0)
	game.questionMap[CTScience] = make([]Question, 0)
	game.questionMap[CTSports] = make([]Question, 0)
	game.questionMap[CTRock] = make([]Question, 0)

	for i := 0; i <= NUM_QUESTIONS; i++ {
		game.questionMap[CTPop] = append(game.questionMap[CTPop], Question{fmt.Sprintf("Pop Question %d", i), i % 4})
		game.questionMap[CTScience] = append(game.questionMap[CTScience], Question{fmt.Sprintf("Science Question %d", i), i % 4})
		game.questionMap[CTSports] = append(game.questionMap[CTSports], Question{fmt.Sprintf("Sports Question %d", i), i % 4})
		game.questionMap[CTRock] = append(game.questionMap[CTRock], Question{fmt.Sprintf("Rock Question %d", i), i % 4})
	}

	return game
}

func (me *Game) playerCount() int {
	return len(me.players)
}

func (me *Game) IsPlayable() bool {
	return me.playerCount() >= MIN_PLAYERS && me.playerCount() <= MAX_PLAYERS
}

func (me *Game) nextPlayer() {
	me.currentPlayer = (me.currentPlayer + 1) % me.playerCount()
}

func (me *Game) Add(playerName string) bool {
	me.players = append(me.players, playerName)
	me.places = append(me.places, 0)
	me.purses = append(me.purses, 0)
	me.penaltyStates = append(me.penaltyStates, PSNone)

	fmt.Printf("Player #%d added: %s\n", me.playerCount(), playerName)
	return true
}

func (me *Game) movePlayer(roll int) {
	me.places[me.currentPlayer] = (me.places[me.currentPlayer] + roll) % NUM_PLACES
}

func (me *Game) addCoin(amount int) {
	me.purses[me.currentPlayer] += amount
	fmt.Printf("%s now has %d Gold Coins.\n", me.players[me.currentPlayer], me.purses[me.currentPlayer])
}

func (me *Game) Roll(roll int) {
	fmt.Printf("%s is the current player\n", me.players[me.currentPlayer])
	fmt.Printf("They have rolled a %d\n", roll)

	if (me.penaltyStates[me.currentPlayer] == PSPenalty && roll%2 != 0) || me.penaltyStates[me.currentPlayer] == PSNone {
		if me.penaltyStates[me.currentPlayer] == PSPenalty {
			me.penaltyStates[me.currentPlayer] = PSLeavingPenalty
			fmt.Printf("%s is getting out of the penalty box\n", me.players[me.currentPlayer])
		}

		me.movePlayer(roll)

		fmt.Printf("%s's new location is %d\n", me.players[me.currentPlayer], me.places[me.currentPlayer])
		fmt.Printf("The category is %s\n", CATEGORY_TYPE_MAP[me.currentCategory()])
		me.askQuestion()

	} else if me.penaltyStates[me.currentPlayer] == PSPenalty {
		fmt.Printf("%s is not getting out of the penalty box\n", me.players[me.currentPlayer])
	}
}

func (me *Game) askQuestion() {
	curCategory := me.currentCategory()
	q := me.questionMap[curCategory][me.questionIndexes[curCategory]]
	me.questionIndexes[curCategory] = (me.questionIndexes[curCategory] + 1) % NUM_QUESTIONS

	fmt.Printf("Asking question: %s\n", q.question)
}

func (me *Game) currentCategory() CategoryType {
	return CategoryType(me.places[me.currentPlayer] % len(CATEGORY_TYPE_MAP))
}

func (me *Game) WrongAnswer() bool {
	fmt.Println("Question was incorrectly answered")
	fmt.Printf("%s was sent to the penalty box\n", me.players[me.currentPlayer])
	me.penaltyStates[me.currentPlayer] = PSPenalty

	return false
}

func (me *Game) WasCorrectlyAnswered() bool {
	switch me.penaltyStates[me.currentPlayer] {
	case PSPenalty:
		return false
	case PSLeavingPenalty:
		me.penaltyStates[me.currentPlayer] = PSNone
		fallthrough
	default:
		fmt.Println("Answer was correct!!!!")
		me.addCoin(1)

		winner := me.didPlayerWin()
		return winner
	}
}

func (me *Game) didPlayerWin() bool {
	return me.purses[me.currentPlayer] >= WINNER_COIN_COUNT
}

func main() {
	winner := false

	game := NewGame()

	game.Add("Chet")
	game.Add("Pat")
	game.Add("Sue")

	for {
		game.Roll(rand.Intn(5) + 1)

		if rand.Intn(9) == 7 {
			winner = game.WrongAnswer()
		} else {
			winner = game.WasCorrectlyAnswered()
		}

		if winner {
			fmt.Printf("%s has won!!!", game.players[game.currentPlayer])
			break
		}

		game.nextPlayer()
		fmt.Println()
	}
}
