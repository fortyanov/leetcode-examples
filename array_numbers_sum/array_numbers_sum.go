package main

/*

Мы хотим складывать очень большие числа, которые превышают емкость базовых типов, поэтому мы храним их в виде массива неотрицательных чисел.

Нужно написать функцию, которая примет на вход два таких массива, вычислит сумму чисел, представленных массивами, и вернет результат в виде такого же массива.

Пример
# Пример 1
# ввод
arr1 = [1, 2, 3] # число 123
arr2 = [4, 5, 6] # число 456
# вывод
res = [5, 7, 9] # число 579. Допустим ответ с первым незначимым нулем [0, 5, 7, 9]
# Пример 2
# ввод
# ввод
arr1 = [5, 4, 4] # число 544
arr2 = [4, 5, 6] # число 456
# вывод
res = [1, 0, 0, 0] # число 1000
*/

import (
	"fmt"
	"leetcode-examples/utils/data_types"
)

func sumLists(arr1, arr2 []uint64) []uint64 {
	var stk1, stk2, stkRes data_types.Stack
	for _, v := range arr1 {
		stk1.Push(v)
	}
	for _, v := range arr2 {
		stk2.Push(v)
	}

	var (
		vsum       uint64
		isOverflow bool
	)

	for {
		v1, pos1 := stk1.Pop()
		v2, pos2 := stk2.Pop()

		if pos1 == -1 && pos2 == -1 {
			if isOverflow {
				stkRes.Push(1)
			}
			break
		}

		vsum = v1 + v2

		if isOverflow {
			vsum += 1
			isOverflow = false
		}

		if vsum > 9 {
			vsum -= 10
			isOverflow = true
		}

		stkRes.Push(vsum)
	}

	return stkRes.ToArray()
}

func main() {
	res1 := sumLists([]uint64{5, 4, 4}, []uint64{4, 5, 6})
	fmt.Println(res1)

	res2 := sumLists([]uint64{1, 2, 3}, []uint64{4, 5, 6})
	fmt.Println(res2)
}
