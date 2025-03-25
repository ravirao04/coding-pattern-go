package intervals

import (
	"sort"
)

type Interval struct {
	start int
	stop  int
}

type Intervals []*Interval

func (items Intervals) Len() int {
	return len(items)
}
func (items Intervals) Less(i, j int) bool {
	return items[i].start < items[j].start
}

func (items Intervals) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func (items *Intervals) Push(i any) {
	*items = append(*items, i.(*Interval))
}
func (items *Intervals) Pop() any {
	item := (*items)[len(*items)-1]
	*items = (*items)[0 : len(*items)-1]
	return item
}

func MergeOverlappingIntervals(intervals [][]int) [][]int {
	sortableIntervals := Intervals{}
	for _, interval := range intervals {
		sortableIntervals.Push(&Interval{start: interval[0], stop: interval[1]})
	}
	sort.Sort(sortableIntervals)

	var result [][]int
	init := []int{sortableIntervals[0].start, sortableIntervals[0].stop}
	result = append(result, init)
	for i := 1; i < len(sortableIntervals); i++ {
		a := result[len(result)-1]
		b := []int{sortableIntervals[i].start, sortableIntervals[i].stop}
		if a[1] >= b[0] {
			result[len(result)-1] = []int{a[0], max(a[1], b[1])}
		} else {
			result = append(result, b)
		}
	}
	return result
}
