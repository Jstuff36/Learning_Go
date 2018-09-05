package main

import (
	"fmt"
	"math"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
	"strings"
)

// Implement own Sqrt funciton

func Sqrt(num float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - num) / (2 *  z)	
		fmt.Println(z)
	}
	return z
}

func SqrtSmartStop(num float64) float64 {
	curr := 1.0
	old := 0.0
	for math.Abs(old - curr) > 0.1 {
		old = curr
		curr -= (curr * curr - num) / (2 * curr)
		fmt.Println(curr)
	}
	return curr
}

func main() {
	fmt.Println(SqrtSmartStop(1424))
}

// Slices exercise
func Pic(dx, dy int) [][]uint8 {
	picValues := make([][]uint8, dy)
	for i := range picValues {
		picValues[i] = make([]uint8, dx)
		for j := range picValues[i] {
			picValues[i][j] = uint8((i+j)/2)
		}
	}
	return picValues
}


func main() {
	pic.Show(Pic)
}

// Count number of times each word appears in the sentence

func WordCount(s string) map[string]int {
	dic := make(map[string]int)
	words := strings.Fields(s);
	for _, word := range words {
		if _, ok := dic[word]; ok {
			dic[word] += 1;
		} else {
			dic[word] = 1;
		}
	}
	return dic
}

func main() {
	wc.Test(WordCount)
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	old := 1
	older := 0
	return func() int {
		nextNum := old + older
		older = old
		old = nextNum
		return nextNum
		
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
