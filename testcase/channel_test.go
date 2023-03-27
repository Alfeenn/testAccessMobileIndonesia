package main

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelData(t *testing.T) {
	num := []int{1, 2, 3, 4, 5}
	channel := make(chan []int)

	var ganjil []int
	defer close(channel)
	go func() {
		time.Sleep(3 * time.Second)
		for _, v := range num {
			if v%2 != 0 {
				ganjil = append(ganjil, v)
			}
		}

		channel <- ganjil
	}()
	data := <-channel
	fmt.Println(data)

}
