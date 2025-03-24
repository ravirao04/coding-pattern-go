package hashmap_and_set

func zeroStriping(matrix [][]int) [][]int {
	rowSet := make(map[int]struct{})
	columSet := make(map[int]struct{})

	for rowIndex, row := range matrix {
		for columnIndex, value := range row {
			if value == 0 {
				rowSet[rowIndex] = struct{}{}
				columSet[columnIndex] = struct{}{}
			}

		}
	}

	for rowIndex, row := range matrix {
		for columnIndex, _ := range row {
			if _, found := rowSet[rowIndex]; found {
				matrix[rowIndex][columnIndex] = 0
			}
			if _, found := columSet[columnIndex]; found {
				matrix[rowIndex][columnIndex] = 0
			}
		}

	}
	return matrix

}
