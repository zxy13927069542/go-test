package leetcode

//给你一个大小为 m x n 的整数矩阵 grid ，表示一个网格。另给你三个整数 row、col 和 color 。网格中的每个值表示该位置处的网格块的颜色
//。
//
// 两个网格块属于同一 连通分量 需满足下述全部条件：
//
//
// 两个网格块颜色相同
// 在上、下、左、右任意一个方向上相邻
//
//
// 连通分量的边界 是指连通分量中满足下述条件之一的所有网格块：
//
//
// 在上、下、左、右四个方向上与不属于同一连通分量的网格块相邻
// 在网格的边界上（第一行/列或最后一行/列）
//
//
// 请你使用指定颜色 color 为所有包含网格块 grid[row][col] 的 连通分量的边界 进行着色，并返回最终的网格 grid 。
//
//
//
// 示例 1：
//
//
//输入：grid = [[1,1],[1,2]], row = 0, col = 0, color = 3
//输出：[[3,3],[3,2]]
//
// 示例 2：
//
//
//输入：grid = [[1,2,2],[2,3,2]], row = 0, col = 1, color = 3
//输出：[[1,3,3],[2,3,3]]
//
// 示例 3：
//
//
//输入：grid = [[1,1,1],[1,1,1],[1,1,1]], row = 1, col = 1, color = 2
//输出：[[2,2,2],[2,1,2],[2,2,2]]
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 50
// 1 <= grid[i][j], color <= 1000
// 0 <= row < m
// 0 <= col < n
//
//
//
// Related Topics 深度优先搜索 广度优先搜索 数组 矩阵
// 👍 109 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
type Node struct {
	row int
	col int
	//isBorder bool
}

//存储同一连通分量的边界结点
var nodes []*Node

//原始颜色
var source int

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	return dfsSolution(grid, row, col, color)
}

//
//  dfsSolution
//  @Description: 深度遍历
//
func dfsSolution(grid [][]int, row int, col int, color int) [][]int {
	//代表上下左右四个方向，与原坐标相加后即可得到上下左右四个坐标
	directions := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	//表示连通分量的原始颜色
	originColor := grid[row][col]

	//表示二维数组中的每一个结点
	type node struct {
		x int
		y int
	}

	//标志当前坐标是否被访问
	vis := make([][]bool, len(grid))
	for i := range vis {
		vis[i] = make([]bool, len(grid[i]))
	}

	//存储边界结点
	var borders []node

	//point := node{
	//	row,
	//	col,
	//}

	var dfs func(int, int)
	dfs = func(x, y int) {
		isBorder := false
		vis[x][y] = true

		//对四个方向的结点进行处理
		for _, dir := range directions {
			nx, ny := x+dir.x, y+dir.y

			//判断是否是边界,逆向判断，先用 不是边界的条件，取反就是 边界的条件
			if !(nx >= 0 && nx < len(grid) && ny >= 0 && ny < (len(grid[nx])) && grid[nx][ny] == originColor) {
				isBorder = true
			} else if !vis[nx][ny] {
				dfs(nx, ny)
			}

			if isBorder {
				borders = append(borders, node{x, y})
			}
		}
	}

	dfs(row, col)

	//着色
	for _, point := range borders {
		grid[point.x][point.y] = color
	}

	return grid

}

func colorBorder2(grid [][]int, row int, col int, color int) [][]int {
	source = 0
	nodes = []*Node{}
	if grid[row][col] == color {
		return grid
	}

	//二层切片拷贝
	tmp := make([][]int, len(grid))
	for i := 0; i < len(tmp); i++ {
		son := make([]int, len(grid[i]))
		copy(son, grid[i])
		tmp[i] = son
	}

	//二层切片拷贝
	tmp1 := make([][]int, len(grid))
	for i := 0; i < len(tmp1); i++ {
		son := make([]int, len(grid[i]))
		copy(son, grid[i])
		tmp1[i] = son
	}

	colorBorder1(tmp, row, col, color)
	for _, node := range nodes {
		//是边界
		if node.row == 0 || node.row == len(grid)-1 || node.col == 0 || node.col == len(grid[node.row])-1 || exist(grid, node.row, node.col, source) {
			//node.isBorder = true
			tmp1[node.row][node.col] = color
		}
		//if node.isBorder {
		//	grid[node.row][node.col] = color
		//}
	}
	return tmp1
}

func colorBorder1(grid [][]int, row int, col int, color int) {
	//初始化原始颜色
	if source == 0 {
		source = grid[row][col]
	}

	//表示当前节点
	node := Node{
		row: row,
		col: col,
		//isBorder: false,
	}

	//表示当前节点已遍历
	grid[row][col] = -1
	nodes = append(nodes, &node)

	//向上遍历
	if row > 0 && grid[row-1][col] == source {
		colorBorder1(grid, row-1, col, color)
	}

	//向下遍历
	if row < len(grid)-1 && grid[row+1][col] == source {
		colorBorder1(grid, row+1, col, color)
	}

	//向左遍历
	if col > 0 && grid[row][col-1] == source {
		colorBorder1(grid, row, col-1, color)
	}

	//向右遍历
	if col < len(grid[row])-1 && grid[row][col+1] == source {
		colorBorder1(grid, row, col+1, color)
	}

	return

}

//是否存在异色
func exist(grid [][]int, row, col, color int) bool {

	//上
	if row > 0 && grid[row-1][col] != color {
		return true
	}

	//下
	if row < len(grid)-1 && grid[row+1][col] != color {
		return true
	}

	//左
	if col > 0 && grid[row][col-1] != color {
		return true
	}

	//右
	if col < len(grid[row])-1 && grid[row][col+1] != color {
		return true
	}

	return false

}

//leetcode submit region end(Prohibit modification and deletion)
