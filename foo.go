package foo

import (
	"fmt"
	"time"
)

type Foo interface {
	Do(int) int
	Do2(*time.Time) []int
}

func Bar(f Foo) {
	result := f.Do(123)
	fmt.Println(result)
	// t := time.Now()
	// f.Do2(&t)
	intSlice := f.Do2(nil)
	if len(intSlice) == 5 {
		fmt.Println("yes")
	}
}
