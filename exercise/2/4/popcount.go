package main

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	count := 0
	mask := uint64(1)

	for i := 0; i < 64; i++ {
		if x&mask > 0 {
			count++
		}

		x >>= 1
	}

	return count
}

func main() {
	println(PopCount(250))
}
