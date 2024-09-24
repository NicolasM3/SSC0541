package ui

type Board struct {
	Size   int
	Matrix [][]int
}

func NewBoard(size int) *Board {
	board := Board{
		Size:   size,
		Matrix: make([][]int, size),
	}

	board.Init()

	return &board
}

func (b *Board) Init() {
	for i := 0; i < b.Size; i++ {
		b.Matrix[i] = make([]int, b.Size)
		for j := 0; j < b.Size; j++ {
			b.Matrix[i][j] = 0
		}
	}
}

func (b *Board) Print() {
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			if b.Matrix[i][j] == 1 {
				print("*", " ")
			} else {
				print(" ", " ")
			}
		}
		println()
	}
}

func (b *Board) CountNeighbours(x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
				count++
			}
			if x+i >= 0 && x+i < b.Size && y+j >= 0 && y+j < b.Size {
				count += b.Matrix[x+i][y+j]
			}
		}
	}
	return count
}
