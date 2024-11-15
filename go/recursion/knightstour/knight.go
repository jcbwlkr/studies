package main

type Knight struct {
	x, y int
}

func NewKnight(x, y int) Knight {
	return Knight{
		x: x,
		y: y,
	}
}

func (knight Knight) possibleMoves(board *Chessboard) []Knight {
	xMoves := []int{2, 2, 1, 1, -1, -1, -2, -2}
	yMoves := []int{-1, 1, -2, 2, -2, 2, -1, 1}

	possible := []Knight{}
	for i := 0; i < 8; i++ {
		newX := knight.x + xMoves[i]
		newY := knight.y + yMoves[i]

		if board.spotOpen(newX, newY) {
			possible = append(possible, NewKnight(newX, newY))
		}
	}

	return possible
}
