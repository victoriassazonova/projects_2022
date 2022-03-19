package islands

import matrix "hsecode.com/stdlib/matrix/int"

func Count(grid *matrix.Matrix) int {
	rows := grid.Rows
	cols := grid.Cols
	coordinates := make([]int, 0)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid.Get(i, j) == 1 {
				dfs(grid, i, j)
				coordinates = append(coordinates, i)
			}

		}
	}
	return len(coordinates)
}

var rowNbr = [4]int{-1, 0, 0, 1}
var colNbr = [4]int{0, -1, 1, 0}

func dfs(g *matrix.Matrix, i int, j int) {
	g.Set(i, j, -1)
	for h := 0; h < 4; h++ {
		if save(g, i+rowNbr[h], j+colNbr[h]) {
			dfs(g, i+rowNbr[h], j+colNbr[h])

		}
	}
}

func save(g *matrix.Matrix, i int, j int) bool {
	if i >= 0 && i < g.Rows && j >= 0 && j < g.Cols && g.Get(i, j) == 1 {
		return true
	}
	return false
}
