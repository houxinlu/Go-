package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)             //理解map,string:key  int:value
	input := bufio.NewScanner(os.Stdin)        //bufio.NewScanner(),input.Scan()

	for input.Scan() {                    
		counts[input.Text()]++             // line := input.Text(); counts[line] = counts[line] + 1
//               fmt.Println(counts,counts[input.Text()])
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
