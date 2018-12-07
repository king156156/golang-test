package eq

import "fmt"

// #####################################################################
// 請列出下面panic和defer的執行順序
func deferStart() {
	deferCall()
	fmt.Println("A")
}

func deferCall() {
	defer func() {
		fmt.Println("B")
	}()
	defer func() {
		fmt.Println("C")
	}()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from r : ", r)
		}
	}()

	defer func() {
		fmt.Println("D")
	}()

	fmt.Println("E")

	panic("Panic")

	fmt.Println("F")
}

// #####################################################################

// 請列出以下Code有什麼問題, 請說明原因和正確修改的方式
type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "AAA", Age: 10},
		{Name: "BBB", Age: 20},
		{Name: "CCC", Age: 30},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
}

// #####################################################################
