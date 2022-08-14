package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// fmt.Println(input.Text())
		counts[input.Text()]++
	}
	fmt.Println(counts)
	for line, n := range counts {
		fmt.Println(n)
		if n > 1 {
			fmt.Println("%d\t%s\n", n, line)
		}
	}
}
