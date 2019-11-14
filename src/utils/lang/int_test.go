package lang

import (
	"fmt"
	"testing"
)

func TestRandomChan(t *testing.T) {
	ch := make(chan int)
	go RandomChan(10, ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
