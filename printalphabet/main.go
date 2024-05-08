// 1 first make the skelton of package maina nd function main
package main

import (
	"github.com/01-edu/z01"
)

func main() {
	// 2 make a for loop with i from 'a' to 'z'
	for i := 'a'; i <= 'z'; i++ {
		// 3 print each i with PrintRune
		z01.PrintRune(i)
	}
	// 4 print a new line also
	z01.PrintRune('\n')
}
