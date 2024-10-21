package main

import (
	"fmt"
	"sort"
)

/*
Дан массив целых чисел nums и целое число k. Нужно написать функцию, которая вынимает из массива k наиболее часто встречающихся элементов.

Пример
# ввод
nums = [5,1,1,1,2,2,3]
k = 2
# вывод (в любом порядке)
[1, 2]
*/

func getMostCommon(nums []int, countMostCommon int) (res []int) {
	resMap := make(map[int]int)

	for _, num := range nums {
		resMap[num] += 1
	}

	resList := make([][2]int, 0, countMostCommon)
	for k, v := range resMap {
		resList = append(resList, [2]int{k, v})
	}

	sort.SliceStable(resList, func(i, j int) bool {
		return resList[i][0] < resList[j][0]
	})

	for i, v := range resList {
		if i >= countMostCommon {
			return
		}

		res = append(res, v[0])
	}

	return res
}

func main() {
	res := getMostCommon([]int{5, 1, 1, 1, 2, 2, 3}, 2)
	fmt.Println(res)
}
