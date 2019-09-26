package main

import (
	"fmt"
)

func swap(a int, b int) (int, int) {
	return b, a
}

func partition(date []int, begin int, end int) int {
	pvalue := date[begin]
	i := begin
	j := begin + 1
	for j < end {
		if date[j] < pvalue {
			i++
			date[i], date[j] = swap(date[i], date[j])
		}
		j++
	}
	date[i], date[begin] = swap(date[i], date[begin])
	return i
}

func quickSort(date []int, begin int, end int) {
	if begin+1 < end {
		mid := partition(date, begin, end)
		quickSort(date, begin, mid)
		quickSort(date, mid+1, end)
	}
}

func putArray(date []int) {
	l := len(date)
	for i := 0; i < l; i++ {
		fmt.Scanf("%d", &date[i])
	}
}

func main() {
	intas := make([]int, 10)
	putArray(intas)
	quickSort(intas, 0, 10)
	fmt.Println(intas)
}
