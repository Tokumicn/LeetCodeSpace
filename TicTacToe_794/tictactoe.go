package main

import (
	"fmt"
	"strings"
)

func main() {

	// "XOX","OX ","OXO"
	// "XOX","O O","XOX"
	fmt.Println(validTicTacToe([]string{"XOX", "OX ", "OXO"}))
}

func win(board []string, p byte) bool {
	// 三行、三列
	for i := 0; i < 3; i++ {
		if board[i][0] == p && board[i][1] == p && board[i][2] == p || // 行
			board[0][i] == p && board[1][i] == p && board[2][i] == p { // 列
			return true
		}
	}

	// 对角线 2种
	return board[0][0] == p && board[1][1] == p && board[2][2] == p ||
		board[0][2] == p && board[1][1] == p && board[2][0] == p
}

func validTicTacToe(board []string) bool {
	oCount, xCount := 0, 0
	for _, row := range board {
		oCount += strings.Count(row, "O")
		xCount += strings.Count(row, "X")
	}

	if oCount != xCount && oCount != xCount-1 {
		return false
	}

	if oCount != xCount && win(board, 'O') {
		return false
	}

	if oCount != xCount-1 && win(board, 'X') {
		return false
	}

	return true
	//return !(oCount != xCount && oCount != xCount-1 ||
	//	oCount != xCount && win(board, 'O') ||
	//	oCount != xCount-1 && win(board, 'X'))
}

/**
总结4个非法条件：
(1) 棋盘中任何时候，X的个数必须等于O的个数，或只比O多一个；
(2) 两个玩家不可能同时赢，一个赢游戏就停了，另一个不允许再放棋子；
(3) 若X赢，因为是X、O轮流下，那么棋盘中的X只能比O多一个；
(4) 若O赢，同理，棋盘中的X和O个数必须相同；

判断是否赢，共有三行、三列、2个对角线共8个情况需要判断。
**/
