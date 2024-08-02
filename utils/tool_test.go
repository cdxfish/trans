package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestDebounce(t *testing.T) {
	f := Debounce(func() {
		fmt.Println("testDebounce")
	}, 5)
	for i := 0; i < 5; i++ {
		f()
		time.Sleep(3 * time.Second)
	}
	time.Sleep(100 * time.Second)
}
func TestThrottled(t *testing.T) {
	f := Throttled(func() {
		fmt.Println("testThrottled")
	}, 5)
	for i := 0; i < 5; i++ {
		f()
		time.Sleep(3 * time.Second)
	}
	time.Sleep(100 * time.Second)
}
