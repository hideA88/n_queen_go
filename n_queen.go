package main

import (
	"flag"
	"fmt"
)

/*Queen ( x int, y int)*/
type Queen struct {
	x int
	y int
}

/*Board (putQueens []Queen, putSize int, boardSize int)*/
type Board struct {
	putQueens []Queen
	putSize   int
	boardSize int
}

func (b Board) deepCopy() Board {
	queens := make([]Queen, len(b.putQueens))
	copy(queens, b.putQueens)
	newBoard := Board{putQueens: queens, putSize: b.putSize, boardSize: b.boardSize}
	return newBoard
}

func (b *Board) setQueen(queen Queen) {
	b.putQueens[b.putSize] = queen
	b.putSize++
}

func main() {
	boardSize := flag.Int("n", 8, "queen size")
	dispOpt := flag.Bool("d", false, "disp queen pattern")
	flag.Parse()
	board := Board{putQueens: make([]Queen, *boardSize), putSize: 0, boardSize: *boardSize}
	result := getPutPattern(board, 0, 0)
	printResult(result, *dispOpt)
}

func getPutPattern(board Board, newX int, newY int) []Board {
	if board.boardSize <= board.putSize {
		return []Board{board}
	}
	if board.boardSize <= newY {
		return nil
	}
	if !isPuttable(board, newX, newY) {
		return getPutPattern(board, newX, newY+1)
	}
	newBoard := board.deepCopy()
	board.setQueen(Queen{x: newX, y: newY})
	answer := getPutPattern(board, newX+1, 0)
	return append(answer, getPutPattern(newBoard, newX, newY+1)...)
}

func isPuttable(board Board, newX int, newY int) bool {
	for i := 0; i < board.putSize; i++ {
		queen := board.putQueens[i]
		slantPlusY := queen.y + (newX - queen.x)
		slantMinusY := queen.y - (newX - queen.x)
		if newY == queen.y || newY == slantPlusY || newY == slantMinusY {
			return false
		}
	}
	return true
}

func printResult(result []Board, dispOpt bool) {
	answer := len(result)
	if dispOpt {
		for _, board := range result {
			dispBoard(board.putQueens, board.boardSize)
		}
	}
	fmt.Println("pattern :", answer)
}

func dispBoard(queens []Queen, boardSize int) {
	for _, queen := range queens {
		y := queen.y
		for i := 0; i < boardSize; i++ {
			if i == y {
				fmt.Printf("%s", "●")
			} else {
				fmt.Printf("%s", "○")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
