package intervals

import "sort"

type IntervalWithPoint struct {
	value     int
	pointType string
}

type IntervalWithPoints []*IntervalWithPoint

func (items IntervalWithPoints) Len() int {
	return len(items)
}
func (items IntervalWithPoints) Less(i, j int) bool {
	if items[i].value == items[j].value {
		if items[i].pointType == "E" {
			return true
		}
	}
	return items[i].value < items[j].value
}

func (items IntervalWithPoints) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func (items *IntervalWithPoints) Push(i any) {
	*items = append(*items, i.(*IntervalWithPoint))
}
func (items *IntervalWithPoints) Pop() any {
	item := (*items)[len(*items)-1]
	*items = (*items)[0 : len(*items)-1]
	return item
}

func LargestOverlapOfIntervals(intervals [][]int) int {
	maxOverlap := 0
	intervalWithPoints := IntervalWithPoints{}
	for _, item := range intervals {
		intervalWithPoints = append(intervalWithPoints, &IntervalWithPoint{value: item[0], pointType: "S"})
		intervalWithPoints = append(intervalWithPoints, &IntervalWithPoint{value: item[1], pointType: "E"})
	}
	sort.Sort(intervalWithPoints)
	counter := 0
	for _, item := range intervalWithPoints {
		if item.pointType == "S" {
			counter++
		} else {
			counter--
		}
		maxOverlap = max(maxOverlap, counter)
	}
	return maxOverlap
}
