package eq

import (
	"testing"
)

func Test_deferStart(t *testing.T) {
	deferStart()
}

/**
考題: defer
印出:
E
D
Recover from r :  Panic
C
B
A
*/

func Test_pase_student(t *testing.T) {
	pase_student()
}

/**
考題: forech
印出: map[AAA:0xc42000a100 BBB:0xc42000a100 CCC:0xc42000a100]
解答: m[stu.Name]=&stu指針都是指向同樣的指針,他的值最後會是forech最後的struct的內容

正確寫法
func pase_student() {
    m := make(map[string]*student)
    stus := []student{
        {Name: "AAA", Age: 10},
        {Name: "BBB", Age: 20},
        {Name: "CCC", Age: 30},
    }
    for i:=0;i<len(stus);i++  {
        m[stus[i].Name] = &stus[i]
    }
}
*/
