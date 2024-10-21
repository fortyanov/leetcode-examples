package data_types

type Stack []uint64

func (s *Stack) Push(v uint64) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (uint64, int) {
	l := len(*s)

	if l == 0 {
		return 0, -1
	}

	res := (*s)[l-1]
	*s = (*s)[:l-1]

	return res, l
}

func (s Stack) ToArray() []uint64 {
	var (
		res []uint64
		v   uint64
		pos int
	)

	for {
		v, pos = s.Pop()
		if pos == -1 {
			break
		}
		res = append(res, v)
	}

	return res
}
