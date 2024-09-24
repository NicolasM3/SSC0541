package main

import (
	logic "cgol/internal/game"
	"cgol/internal/ui"
)

func main() {
	boardSize := 30

	board := ui.NewBoard(boardSize)

	var game = logic.Game{
		Board: board,
	}

	game.NewGame([]logic.ModelPosition{
		{
			X:     0,
			Y:     0,
			Model: logic.Glider(),
		},
		{
			X:     10,
			Y:     10,
			Model: logic.Blinker(),
		},
		{
			X:     25,
			Y:     25,
			Model: logic.ReverseGlinde(),
		},
		{
			X:     24,
			Y:     10,
			Model: logic.Spaceship(),
		},
	})
}
