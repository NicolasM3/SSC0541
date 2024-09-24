package game

import (
	"cgol/internal/ui"
	"time"
)

type Game struct {
	Board *ui.Board
}

type ModelPosition struct {
	X     int
	Y     int
	Model GolType
}

func (g *Game) NewGame(positionList []ModelPosition) {
	for i := 0; i < len(positionList); i++ {
		g.addGol(positionList[i])
	}
	g.process()
}

func (g *Game) addGol(position ModelPosition) {
	for i := 0; i < position.Model.Size; i++ {
		for j := 0; j < position.Model.Size; j++ {
			g.Board.Matrix[position.X+i][position.Y+j] = position.Model.Model[i][j]
		}
	}
}

func (g *Game) process() {

	for {
		printGeneration(g.Board)
		calculateNextGeneration(g.Board)

		time.Sleep(500 * time.Millisecond)
	}

}

func calculateNextGeneration(Board *ui.Board) {
	boardSize := Board.Size
	newBoard := ui.NewBoard(boardSize)
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			countNeighbours := Board.CountNeighbours(i, j)
			if countNeighbours == 3 {
				newBoard.Matrix[i][j] = 1
			} else if countNeighbours == 2 && Board.Matrix[i][j] == 1 {
				newBoard.Matrix[i][j] = 1
			} else {
				newBoard.Matrix[i][j] = 0
			}
		}
	}
	Board.Matrix = newBoard.Matrix
}

func printGeneration(Board *ui.Board) {
	ui.ClearConsole()
	Board.Print()
}
