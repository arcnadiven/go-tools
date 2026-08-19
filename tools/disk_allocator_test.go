package tools

import (
	"fmt"
	"testing"
)

func TestAllocateDisk(t *testing.T) {
	for i := 0; i < 20000; i++ {
		fmt.Println(i, AllocateDisk(i))
	}
}
