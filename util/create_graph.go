package util

import "fmt"

const default_size = 9 // Default size of the sudoku board

type Cell struct {
	x int
	y int
}

func CreateGrid() map[Cell][]Cell {
	Grid := make(map[Cell][]Cell, default_size*default_size)

	for i := 0; i < default_size; i++ {
		for j := 0; j < default_size; j++ {
			cell := Cell{x: i, y: j}
			var adjacents []Cell

			// Adds the row to the cell's adjacency list
			for n := 0; n < default_size; n++ {
				if n != j {
					adjacents = append(adjacents, Cell{x: i, y: n})
				}
			}

			// Adds the column to the cell's adjacency list
			for m := 0; m < default_size; m++ {
				if m != i {
					adjacents = append(adjacents, Cell{x: m, y: j})
				}
			}

			// Adds the subgrid to the cell's adjacency list
			adjacents = append(adjacents,
				Cell{x: i + (i+1)%3 - i%3, y: j + (j+1)%3 - j%3},
				Cell{x: i + (i+1)%3 - i%3, y: j + (j+2)%3 - j%3},
				Cell{x: i + (i+2)%3 - i%3, y: j + (j+1)%3 - j%3},
				Cell{x: i + (i+2)%3 - i%3, y: j + (j+2)%3 - j%3})

			Grid[cell] = adjacents

		}
	}

	for i := 0; i < default_size; i++ {
		for j := 0; j < default_size; j++ {
			fmt.Print("{", i, ",", j, "}: ")
			fmt.Println(Grid[Cell{i, j}])
			fmt.Println("---------------------------------")
		}
	}
	return Grid
}
