package main

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	count := 0

	for x != 0 {
		x &= x - 1
		count++
	}

	return count
}

func main() {
	println(PopCount(250))
}
