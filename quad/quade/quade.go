package piscine

import "github.com/01-edu/z01"

func QuadE(x, y int) {
	if x < 0 || y < 0 {
		return
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 && j == 0 {
				z01.PrintRune('A')
			} else if i == 0 && j == x-1 {
				z01.PrintRune('C')
			} else if i == y-1 && j == 0 {
				z01.PrintRune('C')
			} else if i == y-1 && j == x-1 {
				z01.PrintRune('A')
			}
			if (i > 0 && i < y-1) && (j > 0 && j < x-1) {
				z01.PrintRune(' ')
			}
			if (i == 0 || i == y-1) && (j > 0 && j < x-1) {
				z01.PrintRune('B')
			}
			if (i > 0 && i < y-1) && (j == 0 || j == x-1) {
				z01.PrintRune('B')
			}
			if j == x-1 {
				z01.PrintRune('\n')
			}
		}
	}
}