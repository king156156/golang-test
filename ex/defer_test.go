package ex_test

import (
	"fmt"
	"testing"
)

func Test_defer(t *testing.T) {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		// t.Point() //输出 c c c
		defer t.Value() //输出 c b a
	}
}

type Test struct {
	name string
}

func (this *Test) Point() {
	fmt.Println(this.name)
}

func (this Test) Value() {
	fmt.Println(this.name)
}
