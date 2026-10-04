package main

// 算法思路：DFS 洪水填充。遍历网格，每遇到一个未访问的 '1' 就让岛屿数加一，
// 并从该点 DFS 把整座岛相连的 '1' 全部就地改写成 '0'（标记已访问），避免重复计数
func numIslands(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return
		}
		if grid[r][c] != '1' {
			return
		}
		grid[r][c] = '0'
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}
	return count

}

func main() {

}
