package game

type GolType struct {
	Size  int
	Model [][]int
}

func Glider() GolType {
	model := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
	}

	return GolType{3, model}
}

func Blinker() GolType {
	model := [][]int{
		{0, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	}

	return GolType{3, model}
}

func ReverseGlinde() GolType {
	model := [][]int{
		{1, 1, 1},
		{1, 0, 0},
		{0, 1, 0},
	}

	return GolType{3, model}
}

func Spaceship() GolType {
	model := [][]int{
		{0, 1, 1, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
	}

	return GolType{5, model}
}
