package lab1

import (
	"fmt"
	"math/rand"
)

func GenerateRandomMinefield(row int, col int, mineCount int) [][]int {
	field := make([][]int, row)
	for i := range field {
		field[i] = make([]int, col)
	}

	for i := 0; i < mineCount; i++ {
		row := rand.Intn(row)
		col := rand.Intn(col)
		field[row][col] = -1
	}

	return field

}
func CountMines(field [][]int) [][]int {
	rows := len(field)
	cols := len(field[0])
	result := make([][]int, rows)
	for i := range result {
		result[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if field[i][j] == -1 {
				result[i][j] = -1
				continue
			}
			count := 0
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					nx, ny := i+dx, j+dy
					if nx >= 0 && nx < rows && ny >= 0 && ny < cols && field[nx][ny] == -1 {
						count++
					}
				}
			}
			result[i][j] = count
		}
	}
	return result
}

func GenerateLab4_4() {
	field := GenerateRandomMinefield(25, 25, 99)

	countField := CountMines(field)
	for _, row := range countField {
		for _, col := range row {
			if col == -1 {
				fmt.Print("*", " ")
			} else {
				fmt.Print(col, " ")
			}
		}
		fmt.Println()
	}
}

/*
func main() {
	GenerateLab4_4()
}
*/
