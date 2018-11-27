package ex

import (
	"fmt"
	"testing"
)

// # 基礎寫入測試

func BenchmarkSliceInsert1(b *testing.B) {
	var sl []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sl = append(sl, i)
	}
}

// 測試數據:
// BenchmarkSliceInsert1-4   	100000000	        21.4 ns/op	      49 B/op	       0 allocs/op
// PASS
// ok  	github.com/king156156/golang-test/ex	2.302s

func BenchmarkSliceInsert2(b *testing.B) {
	var sl []int
	b.ResetTimer()
	var count = 100000000
	for i := 0; i < b.N; i++ {
		for p := 0; p < count; p++ {
			sl = append(sl, p)
		}
	}
}

// 測試數據:
// BenchmarkSliceInsert2-4   	       1	1434447434 ns/op	4935058688 B/op	      61 allocs/op
// PASS
// ok  	github.com/king156156/golang-test/ex	1.550s

func BenchmarkMapInsert1(b *testing.B) {
	sl := map[int]struct{}{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sl[i] = struct{}{} // 無消耗內存
	}
}

// 測試數據:
// BenchmarkMapInsert1-4   	10000000	       179 ns/op	      40 B/op	       0 allocs/op
// PASS
// ok  	github.com/king156156/golang-test/ex	1.973s

func BenchmarkMapInsert2(b *testing.B) {
	sl := map[int]struct{}{}
	b.ResetTimer()
	var count = 100000000
	for i := 0; i < b.N; i++ {
		for p := 0; p < count; p++ {
			sl[p] = struct{}{} // 無消耗內存
		}
	}
}

// 測試:
// BenchmarkMapInsert2-4   	       1	21037082912 ns/op	3418690040 B/op	 3904068 allocs/op
// PASS
// ok  	github.com/king156156/golang-test/ex	21.257s

// # map取值測試
func Test_MapIntIncr1(t *testing.T) {
	ddd := map[string]int{}
	ddd["aaa"]++
	fmt.Println(ddd["aaa"])  // output: 1
	fmt.Println(ddd["aaa1"]) // output: 0
	//Benchmark			100000000	        14.8 ns/op
}

func Test_MapIntIncr2(t *testing.T) {
	ddd := map[string]int{}
	v := ddd["aaa"]
	v++
	fmt.Println(ddd["aaa"])  // output:0
	fmt.Println(ddd["aaa1"]) // output: 0
	// 沒有寫入成功
	//Benchmark 		1000000000	         2.30 ns/op
}

func Test_MapIntIncr3(t *testing.T) {
	ddd := map[string]int{}
	v, ok := ddd["aaa"]
	if !ok {
		ddd["aaa"] = 0
		v = ddd["aaa"]
	}
	v++
	fmt.Println(ddd["aaa"])  // output:0
	fmt.Println(ddd["aaa1"]) // output: 0
	// 沒有寫入成功
	//Benchmark 		300000000	         4.40 ns/op
}

// MapIntIncr4
type MapIncr4 int

func (p *MapIncr4) Incr() {
	*p++
}

func Test_MapIntIncr4(t *testing.T) {
	ddd := map[string]*MapIncr4{}
	v, ok := ddd["aaa"]
	if !ok {
		var k MapIncr4
		v, ddd["aaa"] = &k, &k
	}
	v.Incr()
	fmt.Println(*ddd["aaa"]) // output:1
	//300000000	         4.76 ns/op	       0 B/op	       0 allocs/op
}

func Benchmark_MapIntIncr1(b *testing.B) {
	ddd := map[string]int{}
	b.ResetTimer()
	for i := 0; i <= b.N; i++ {
		ddd["aaa"]++
	}
}

// 測試數據
// Benchmark_MapIntIncr1-4   	100000000	        14.8 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	github.com/king156156/golang-test/ex	1.515s

func Benchmark_MapIntIncr2(b *testing.B) {
	ddd := map[string]int{}
	b.ResetTimer()
	for i := 0; i <= b.N; i++ {
		v := ddd["aaa"]
		v++
	}
	// 沒有寫入成功
}

// 測試數據
// Benchmark_MapIntIncr2-4   	1000000000	         2.30 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	github.com/king156156/golang-test/ex	2.908s

func Benchmark_MapIntIncr3(b *testing.B) {
	ddd := map[string]int{}
	b.ResetTimer()
	for i := 0; i <= b.N; i++ {
		v, ok := ddd["aaa"]
		if !ok {
			ddd["aaa"] = 0
			continue
		}
		v++
	}
	// 沒有寫入成功
}

// 測試數據
// Benchmark_MapIntIncr3-4   	300000000	         4.40 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	github.com/king156156/golang-test/ex	1.790s

func Benchmark_MapIntIncr4(b *testing.B) {
	ddd := map[string]*MapIncr4{}
	b.ResetTimer()
	for i := 0; i <= b.N; i++ {
		v, ok := ddd["aaa"]
		if !ok {
			var k MapIncr4
			v, ddd["aaa"] = &k, &k
		}
		v.Incr()
	}
	// fmt.Println("次數:", *ddd["aaa"])
}

// 測試數據
// Benchmark_MapIntIncr4-4   	300000000	         4.63 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	github.com/king156156/golang-test/ex	1.879s

// # slice 測試
func Test_SliceIntIncr1(t *testing.T) {
	n := [10]int{}
	for _, v := range n {
		v++
	}

	fmt.Println(n)

	// 無法寫入
}

func Test_SliceIntIncr2(t *testing.T) {
	n := [10]int{}
	for i := range n {
		n[i]++
	}
	// 寫入
	fmt.Println(n)
}

func Test_SliceIntIncr3(t *testing.T) {
	n := [10]int{}
	for i := range n {
		v := n[i]
		v++
	}
	// 無法寫入
	fmt.Println(n)
}

func Test_SliceIntIncr4(t *testing.T) {
	n := [10]int{}
	for i := range n {
		v := &n[i]
		*v++
	}
	// 寫入
	fmt.Println(n)
}

func Benchmark_SliceIntIncr1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i <= b.N; i++ {
		n := [10]int{}
		for _, v := range n {
			v++
		}
	}
}

// Benchmark_SliceIntIncr1-4   	200000000	         6.34 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	github.com/king156156/golang-test/ex	1.942s

func Benchmark_SliceIntIncr2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i <= b.N; i++ {
		n := [10]int{}
		for i := range n {
			n[i]++
		}
	}
}

// Benchmark_SliceIntIncr2-4   	200000000	         7.42 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	github.com/king156156/golang-test/ex	2.277s

func Benchmark_SliceIntIncr3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i <= b.N; i++ {
		n := [10]int{}
		for i := range n {
			v := n[i]
			v++
		}
	}
}

// Benchmark_SliceIntIncr3-4   	300000000	         4.75 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	github.com/king156156/golang-test/ex	1.943s

func Benchmark_SliceIntIncr4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i <= b.N; i++ {
		n := [10]int{}
		for i := range n {
			v := &n[i]
			*v++
		}
		// 寫入
	}
}

// Benchmark_SliceIntIncr4-4   	200000000	         7.17 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	github.com/king156156/golang-test/ex	2.179s
