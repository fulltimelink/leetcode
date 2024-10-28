package array

type RotateDirection int

const (
	RightRotate       RotateDirection = iota // --  @# 右旋
	LeftRotate                               // --  @# 左旋
	PrincipalDiagonal                        // --  @# 主对角线翻转
	AuxiliaryDiagonal                        // --  @# 副对角线翻转
)

// --  @# 旋转及翻转
func rotate(matrix [][]int, direction RotateDirection) {
	l := len(matrix)
	if l <= 1 {
		return
	}
	if RightRotate == direction {
		// --  @# 上下翻转
		for r := 0; r <= (l-1)/2; r++ {
			oindex := l - 1 - r
			for c := 0; c < len(matrix[0]); c++ {
				matrix[r][c], matrix[oindex][c] = matrix[oindex][c], matrix[r][c]
			}
		}
	} else if direction == LeftRotate {
		// --  @# 左右翻转
		for c := 0; c <= (l-1)/2; c++ {
			oindex := l - 1 - c
			for r := 0; r < l; r++ {
				matrix[r][c], matrix[r][oindex] = matrix[r][oindex], matrix[r][c]
			}
		}
	}
	if RightRotate == direction || LeftRotate == direction || PrincipalDiagonal == direction {
		// --  @# 主对角线翻转
		for r := 0; r < l; r++ {
			for c := r + 1; c < l; c++ {
				matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
			}
		}
	} else if AuxiliaryDiagonal == direction {
		// --  @# 副对角线翻转
		for r := 0; r < l; r++ {
			for c := 0; c < l-r-1; c++ {
				matrix[r][c], matrix[l-1-c][l-1-r] = matrix[l-1-c][l-1-r], matrix[r][c]
			}
		}
	}
}

// --  @# 右旋
func rightRotate(matrix [][]int) {
	rotate(matrix, RightRotate)
}

// --  @# 左旋
func leftRotate(matrix [][]int) {
	rotate(matrix, LeftRotate)
}

// --  @# 主对角线翻转
func principalDiagonal(matrix [][]int) {
	rotate(matrix, PrincipalDiagonal)
}

// --  @# 副对角线翻转
func auxiliaryDiagonal(matrix [][]int) {
	rotate(matrix, AuxiliaryDiagonal)
}
