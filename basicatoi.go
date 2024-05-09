package piscine

func BasicAtoi(s string) int {
	o_number := 0
	c := 0
	for _, word := range s {
		for i := '0'; i < word; i++ {
			c++
		}
		o_number = o_number*10 + c
		c = 0
	}
	return o_number
}
