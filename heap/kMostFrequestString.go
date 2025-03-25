package heap

import "container/heap"

type stringItem struct {
	value string
	freq  int
}

type Items []*stringItem

func (items Items) Len() int {
	return len(items)
}
func (items Items) Less(i, j int) bool {
	return items[i].freq > items[j].freq
}

func (items Items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func (items *Items) Push(i any) {
	*items = append(*items, i.(*stringItem))
}
func (items *Items) Pop() any {
	item := (*items)[len(*items)-1]
	*items = (*items)[0 : len(*items)-1]
	return item
}

func calculateFreq(strs []string) *Items {
	freq := make(map[string]int)
	for _, str := range strs {
		if count, found := freq[str]; found {
			freq[str] = count + 1
		} else {
			freq[str] = 1
		}
	}
	items := Items{}
	for key, item := range freq {
		st := &stringItem{key, item}
		items.Push(st)
	}

	return &items
}

func kMostFrequentString(strs []string, k int) []string {
	items := calculateFreq(strs)
	heap.Init(items)
	result := []string{}
	for i := range k {
		result[i] = heap.Pop(items).(*stringItem).value
	}
	return result
}
