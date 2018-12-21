package code_38

import "fmt"

// 8皇后问题

func FindPlace(n int) {
	place := make([]int, n)
	Core(place, 0, n)
}

func Core(places []int, row int, n int) {
	if row == n {
		print(places)
		return
	}
	for col := 0; col < n; col++ {
		confolict := checkConflict(places, row, col)
		if !confolict {
			continue
		}
		places[row] = col
		Core(places, row+1, n)
		places[row] = 0
	}
}

func checkConflict(places []int, row, col int) bool {
	for i := 0; i < row; i++ {
		row_ := i
		col_ := places[i]
		if col == col_ {
			return false
		}
		diffX := row_ - row
		diffY := col_ - col
		if diffY == diffX || diffY == -diffX {
			return false
		}
	}
	return true
}

func print(places []int) {
	fmt.Println("找到可行解")
	for i := 0; i < len(places); i++ {
		for j := 0; j < places[i]; j++ {
			fmt.Print("* ")
		}
		fmt.Print("H ")
		for j := places[i] + 1; j < len(places); j++ {
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}
}
