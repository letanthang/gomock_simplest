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
}
