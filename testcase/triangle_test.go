package main

import (
	"fmt"
	"testing"
)

func TestTriang(t *testing.T) {
	num := 6

	for i := 0; i < num; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i || i == num-1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
