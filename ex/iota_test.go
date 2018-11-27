package ex_test

import (
	"fmt"
	"testing"
)

// const (
// 	_, _ = iota, iota * 10
// 	a, b
// 	c, d
// )

// const (
// 	_ = iota // 0
// 	a        // 1
// 	b = 100  // 100
// 	c        // 100
// 	d = iota // 4
// )

const (
	_ = iota               // 0 int
	a                      // 1 int
	b float64 = iota * 0.1 // 2 * 0.1 = 0.2  float
	c                      // 3 * 0.1 = 0.3  float
	d                      // 4 * 0.1 = 0.4  float   or  Ex: d = iota >>> 4 int
)

// 自訂類型
type color byte

const (
	black color = iota
	red
	bule
)

func Test_runColor(t *testing.T) {
	var p = struct {
		Ddd int
	}{
		2,
	}
	testcolor(red)
	testcolor(100)
	// testcolor(p.Ddd) // 此格式無法輸入
	_ = p
}

// 可以限制傳入格式
func testcolor(c color) {
	fmt.Println(c)
}
