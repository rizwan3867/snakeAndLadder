package service

import (
	"fmt"
	"math/rand"

	"snakeAndLadder/pkg/internal/model"
)

type game struct {
	board   *model.Board
	players []*model.Player
}

type GameInterface interface {
	MakeMove(playerNo int)
	IsWinOrLose(playerNo int) bool
}

func NewGame(noOfPlayers int, boardSize int) GameInterface {
	g := &game{}
	g.players = make([]*model.Player, noOfPlayers)
	g.init()
	g.board.Size = boardSize
	return g
}

func (g *game) init() {
	// initialise the board and number of players
	g.board = &model.Board{}
	g.board.Snakes = map[int]int{
		16: 6,
		47: 26,
		49: 11,
		56: 53,
		62: 19,
		64: 60,
		87: 24,
		93: 73,
		95: 75,
		98: 78,
	}
	g.board.Ladder = map[int]int{
		1:  38,
		4:  14,
		9:  31,
		21: 42,
		28: 84,
		36: 44,
		51: 67,
		71: 91,
		80: 100,
	}

	for i := 0; i < len(g.players); i++ {
		g.players[i] = &model.Player{
			Position: 0,
		}
	}

}

func (g *game) MakeMove(playerNo int) {
	roll := rand.Intn(6) + 1
	player := g.players[playerNo]
	fmt.Printf("Player %d rolled %d \n", playerNo+1, roll)
	if player.Position+roll > 100 {
		fmt.Printf("Player stays at position : %d \n", player.Position)
		return
	}
	player.Position = player.Position + roll
	// check snakes 
	if value, ok := g.board.Snakes[player.Position]; ok {
		fmt.Printf("Player %d hit by snake. Moves down to %d \n", playerNo+1, value)
		player.Position = value
	}
	// check ladder 
	if value, ok := g.board.Ladder[player.Position]; ok {
		fmt.Printf("Player %d lifted by ladder.  Moves up to %d \n", playerNo+1, value)
		player.Position = value
	}

}

func (g *game) IsWinOrLose(playerNo int) bool {
	player := g.players[playerNo]
	if player.Position == 100 {
		return true
	}
	return false
}
