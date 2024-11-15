package main

import "fmt"

func main() {

	// Initialize our pieces.
	knight := NewKnight(0, 0)
	board := NewChessboard(8)

	// Record the starting position.
	board.recordMove(knight.x, knight.y)

	// Start recursing until we find a solution or don't.
	if tour(knight, board) {
		printResult("Solution found! :)", board.history)
	} else {
		fmt.Println("No solution found :(")
	}
}

func tour(knight Knight, board *Chessboard) bool {

	// Look at all of the possible moves from this spot.
	for _, possible := range knight.possibleMoves(board) {

		// Move to that new spot
		board.recordMove(possible.x, possible.y)

		// If that completes our tour then we found a solution!
		if board.done() {
			return true
		}

		// If this wasn't our last move, try more from this point.
		// If it returns true then it found our last move.
		if tour(possible, board) {
			return true
		}

		// If we get here then we hit a dead end. Undo that and
		// try the next possiblity.
		board.undoMove(possible.x, possible.y)
	}

	// If we tried all of our possible moves and none worked
	// then we're done.
	return false
}
