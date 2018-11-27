package ex

import (
	"fmt"
	"testing"
)

func Test_Callts(t *testing.T) {
	type (
		user struct {
			id   int
			name string
		}

		Ddd func()
	)
	data := &user{}
	data.name = "Peter"
}

func Test_Closure(t *testing.T) {
	fn := func() string {
		return "123"
	}
	fmt.Println(fn()) // 123
}

// # 模塊方法測試

type funcA struct {
	Age int
}

func (fa funcA) Get() int {
	return fa.Age
}

// 指針
func (fa *funcA) AGet() int {
	return fa.Age
}

func (fa *funcA) AGet1() *int {
	return &fa.Age
}

// 複製參數
func (fa funcA) Set1(age int) {
	fa.Age = age
}

// 指針
func (fa *funcA) Set2(age int) {
	fa.Age = age
}

func Test_funcA(t *testing.T) {
	a := funcA{}
	a.Set1(18)
	fmt.Println(a.Get()) // output: 0

	a.Set2(18)
	fmt.Println(a.Get()) // output: 18

	n := a.AGet()
	n = 20
	fmt.Println(a.Get())  // output: 18
	fmt.Println(a.AGet()) // output: 18
	fmt.Println(n)        // output: 20(n)

	n1 := a.AGet1() // 取出指針位置
	*n1 = 25
	fmt.Println(a.Get())  // output: 25
	fmt.Println(a.AGet()) // output: 25
	fmt.Println(n1)       // output: 指針地址
}

func Benchmark_FuncA_Get(b *testing.B) {
	a := funcA{}
	for i := 0; i < b.N; i++ {
		a.Get()
	}
	// Benchmark_FuncA_Get-4   	2000000000	         0.60 ns/op	       0 B/op	       0 allocs/op
}

func Benchmark_FuncA_AGet(b *testing.B) {
	a := funcA{}
	for i := 0; i < b.N; i++ {
		a.AGet()
	}
	// Benchmark_FuncA_AGet-4   	2000000000	         0.29 ns/op	       0 B/op	       0 allocs/op
}
func Benchmark_FuncA_AGet1(b *testing.B) {
	a := funcA{}
	for i := 0; i < b.N; i++ {
		a.AGet1()
	}
	// Benchmark_FuncA_AGet1-4   	2000000000	         0.58 ns/op	       0 B/op	       0 allocs/op
}
