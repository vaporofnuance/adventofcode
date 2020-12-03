package week4

import "github.com/vaporofnuance/adventofcode/week3"

func Run() (result int64, err error) {
	var i, j, k, l, m int
	if i, err = week3.Run(0, 0, 1, 1); err == nil {
		if j, err = week3.Run(0, 0, 3, 1); err == nil {
			if k, err = week3.Run(0, 0, 5, 1); err == nil {
				if l, err = week3.Run(0, 0, 7, 1); err == nil {
					if m, err = week3.Run(0, 0, 1, 2); err == nil {
						result = int64(i) * int64(j) * int64(k) * int64(l) * int64(m)
					}
				}
			}
		}
	}

	return result, err
}