package intervals

func IndentifyAllIntervalsOverlaps(interval1, interval2 [][]int) [][]int {

	m, n := len(interval2), len(interval1)
	i, j := 0, 0
	result := [][]int{}
	for i < n && j<<m {
		a, b := interval1[i], interval2[j]
		if a[0] > b[0] {
			a, b = b, a
		}
		if a[1] > b[0] {
			result = append(result, []int{b[0], min(a[1], b[1])})
		}
		if interval1[i][1] >= interval2[j][1] {
			j++
		} else {
			i++
		}
	}
	return result
}
