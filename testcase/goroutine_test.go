package main

import (
	goroutinedata "github.com/Alfeenn/testcase/goroutineData"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {

	go goroutinedata.DummyData1()
	go goroutinedata.DummyData2()
	time.Sleep(1 * time.Second)
}
