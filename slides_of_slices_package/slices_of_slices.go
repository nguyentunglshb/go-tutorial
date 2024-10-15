package slides_of_slices_package

import "fmt"

type board [][]string

func createDefaultBoard() board {
	board := board{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	return board
}

func (board board) updateBoard( row int, col int, value string) {
	board[row][col] = value
}

func (board board) printBoard() {
	for _, row := range board {
		for _, value := range row {
			fmt.Printf("%v ", value)
		}
		fmt.Println()
	}
}




func (board board) appendRow(value []string) {
    board = append(board, value)
}

func (board board) appendColumn( value []string) {
	for i := 0; i < len(board); i++ {
		board[i] = append(board[i], value[i])
	}
}

func SlicesOfSlices() {
	defaultBoard := createDefaultBoard()

	defaultBoard.updateBoard(0, 0, "X")
	defaultBoard.updateBoard(2, 2, "O")
	defaultBoard.updateBoard(1, 1, "X")
	defaultBoard.updateBoard(1, 2, "O")
	defaultBoard.appendColumn([]string{"X", "X", "X", "X"})
	defaultBoard.appendRow([]string{"_", "_", "_", "_"})

	defaultBoard.printBoard()

}