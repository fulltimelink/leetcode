package array

// --  @# 给你一个大小为 m x n 的矩阵 mat ，请以对角线遍历的顺序，用一个数组返回这个矩阵中的所有元素。
func findDiagonalOrder(mat [][]int) []int {
	rows, cols := len(mat), len(mat[0])

	// --  @# 处理极端情况
	if rows == 0 || cols == 0 {
		return []int{}
	}
	if rows == 1 {
		return mat[0]
	}

	// --  @# 遍历方向定义
	const (
		LoopLeft = iota
		LoopRight
	)

	direction := LoopRight
	i, j := 0, 0
	var r []int
	for i != rows-1 || j != cols-1 {
		tempDirection := direction
		if i < 0 || j > cols-1 {
			tempDirection = LoopLeft
		} else if j < 0 || i > rows-1 {
			tempDirection = LoopRight
		} else {
			r = append(r, mat[i][j])
		}

		if direction != tempDirection {
			// --  @# 方向变换时处理
			if direction == LoopRight {
				i += 1
			} else {
				j += 1
			}
			direction = tempDirection
		} else {
			// --  @# 按原方向再走一步
			if direction == LoopRight {
				i -= 1
				j += 1
			} else {
				i += 1
				j -= 1
			}
		}
	}
	r = append(r, mat[rows-1][cols-1])
	return r
}
