package foo

import "fmt"

type Foo interface {
	Do(int) int
}

func Bar(f Foo) {
	result := f.Do(123)
	fmt.Println(result)
}
