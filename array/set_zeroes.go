package array

func setZeroes(matrix [][]int) {
	rIdx := make(map[int]struct{})
	cIdx := make(map[int]struct{})
	for ridx, row := range matrix {
		for cidx, val := range row {
			if val == 0 {
				rIdx[ridx] = struct{}{}
				cIdx[cidx] = struct{}{}
			}
		}
	}
	for ridx, row := range matrix {
		_, okr := rIdx[ridx]
		for cidx := range row {
			_, okc := cIdx[cidx]
			if okr || okc {
				matrix[ridx][cidx] = 0
			}

		}
	}
}
