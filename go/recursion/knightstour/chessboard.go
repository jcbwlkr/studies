package main

type Chessboard struct {
	size    int
	move    int
	history [][]int
}

func NewChessboard(size int) *Chessboard {
	history := make([][]int, size)
	for i := range history {
		history[i] = make([]int, size)
	}
	return &Chessboard{
		size:    size,
		move:    0,
		history: history,
	}
}

func (b *Chessboard) spotOpen(x, y int) bool {
	// Make sure they don't go off the board.
	if x < 0 || x >= b.size || y < 0 || y >= b.size {
		return false
	}

	// Return true if we have no move recorded for that spot.
	return b.history[x][y] == 0
}

func (b *Chessboard) recordMove(x, y int) {
	b.move++
	b.history[x][y] = b.move
}

func (b *Chessboard) undoMove(x, y int) {
	b.move--
	b.history[x][y] = 0
}

func (b *Chessboard) done() bool {
	return b.move == b.size*b.size
}
