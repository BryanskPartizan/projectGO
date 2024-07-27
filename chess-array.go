package main

import "fmt"

func boardDisplay(board [8][8]rune) {
	for _, row := range board {
		for _, column := range row {
			if column == 0 {
				fmt.Print(" ")
			} else {
				fmt.Printf("%c ", column)
			}
		}
		fmt.Println()
	}
}

func main() {
	var chessBoard [8][8]rune
	blackFigures := [6]rune{'k', 'q', 'r', 'b', 'n', 'p'}
	whiteFigures := [6]rune{'K', 'Q', 'R', 'B', 'N', 'P'}

	for column := range chessBoard[1] {
		chessBoard[1][column] = 'p'
		chessBoard[6][column] = 'P'
	}

	// TODO убрать кринж
	chessBoard[0][0] = blackFigures[2]
	chessBoard[0][7] = blackFigures[2]
	chessBoard[0][1] = blackFigures[4]
	chessBoard[0][6] = blackFigures[4]
	chessBoard[0][2] = blackFigures[3]
	chessBoard[0][5] = blackFigures[3]
	chessBoard[0][3] = blackFigures[1]
	chessBoard[0][4] = blackFigures[0]

	chessBoard[7][0] = whiteFigures[2]
	chessBoard[7][7] = whiteFigures[2]
	chessBoard[7][1] = whiteFigures[4]
	chessBoard[7][6] = whiteFigures[4]
	chessBoard[7][2] = whiteFigures[3]
	chessBoard[7][5] = whiteFigures[3]
	chessBoard[7][3] = whiteFigures[1]
	chessBoard[7][4] = whiteFigures[0]

	boardDisplay(chessBoard)

}
