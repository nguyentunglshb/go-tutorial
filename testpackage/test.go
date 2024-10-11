package testpackage

import "fmt"

func AnyFunc(step int) {
	for i := 0; i < step; i++ {
		fmt.Printf("Step %v\n", i+1)
	}
}
