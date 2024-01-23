package main

import "fmt"

type Piece struct {
	move      int
	kill      int
	pieceType string // "pawn", "rook", "knight", "bishop", "queen", "king"
	color     string
}

func createBoard(lin int, col int) [][]string {
	board := make([][]string, lin)
	for i := range board {
		board[i] = make([]string, col)
	}
	for i := 0; i < lin; i++ {
		for j := 0; j < col; j++ {
			board[i][j] = "-"
		}
	}
	return board
}

func main() {

	chessBoard := createBoard(8, 8)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(chessBoard[i][j], " ")
		}
		fmt.Println()
	}
}
