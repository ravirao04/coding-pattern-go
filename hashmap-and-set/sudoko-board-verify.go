package hashmap_and_set

func verifySudokoBoard(board [][]int) bool {
	rowSets := make([]map[int]struct{}, 9)    // 9 rows
	columnSets := make([]map[int]struct{}, 9) // 9 columns
	// Initialize subBoardSets as a slice of slices of maps (3x3 subgrids)
	subBoardSets := make([][]map[int]struct{}, 3) // 3 rows of subgrids

	for i := 0; i < 3; i++ {
		subBoardSets[i] = make([]map[int]struct{}, 3) // 3 columns of subgrids
	}
	// Initialize each set (map) inside the slices
	for i := 0; i < 9; i++ {
		rowSets[i] = make(map[int]struct{})
		columnSets[i] = make(map[int]struct{})
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			subBoardSets[i][j] = make(map[int]struct{})
		}
	}

	for rowIndex, row := range board {
		for columnIndex, value := range row {
			if value == 0 { // Skip empty cells
				continue
			}
			if _, found := rowSets[rowIndex][value]; found {
				return false
			} else {
				rowSets[rowIndex][value] = struct{}{}
			}

			if _, found := columnSets[columnIndex][value]; found {
				return false
			} else {
				columnSets[columnIndex][value] = struct{}{}
			}

			if _, found := subBoardSets[rowIndex/3][columnIndex/3][value]; found {
				return false
			} else {
				subBoardSets[rowIndex/3][columnIndex/3][value] = struct{}{}
			}

		}
	}
	return true
}
