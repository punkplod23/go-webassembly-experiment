package main

import (
	"fmt"
	wordchecker "jtr/internal"
)

func main() {

	fmt.Println(wordchecker.BruteForce("MD5", "26253c50741faa9c2e2b836773c69fe6", 100000, 5))

}
