package main

/*
Программа выводящая позиции соответствующих пар скобок в арифметическом выражении. Порядок выдачи пар не важен.
Примеры:
# Пример 1
# ввод
a*(b+c)
# вывод
(2,6)
# Пример 2
# ввод
((a+b)-c)*d*(e-f)
# вывод
(0,8),(1,5),(12,16)
*/

import (
	"fmt"
	"leetcode-examples/utils/data_types"
)

func bracketPositions(inp string) [][2]uint64 {
	var (
		res      [][2]uint64
		stk      data_types.Stack
		charPos  uint64
		stackPos int
	)

	for i, char := range inp {
		if char == '(' {
			stk.Push(uint64(i))
		}
		if char == ')' {
			charPos, stackPos = stk.Pop()
			if stackPos == -1 {
				continue
			}
			res = append(res, [2]uint64{charPos, uint64(i)})
		}
	}
	return res
}

func main() {
	res1 := bracketPositions("a*(b+c)")
	fmt.Println(res1)

	res2 := bracketPositions("((a+b)-c)*d*(e-f)")
	fmt.Println(res2)

	res3 := bracketPositions("")
	fmt.Println(res3)

	res4 := bracketPositions(")()((")
	fmt.Println(res4)
}
