package main

import "fmt"

func PrintPrimes(n int) int {

	s := make([]bool, n+1)

	count := 0
	for i := 2; i <= n; i++ {
		if !s[i] {
			count++

			fmt.Printf("%3d ", i)
			for k := i + i; k <= n; k += i {

				s[k] = true
			}
			if count%10 == 0 {
				fmt.Println()
			}
		}
	}
	return count // return
}

func main() {
	count := PrintPrimes(300)
	fmt.Printf("\nКоличество простых чисел: %d\n", count)
}
