package main

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var sum int

	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>i)])
	}

	return sum
}

func main() {
	println(PopCount(22))
}
