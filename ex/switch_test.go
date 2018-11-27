package ex

import (
	"fmt"
	"testing"
	"time"
)

// Test_Switch1 switch預設是有break, 可以透過fallthrough強制下一句case
func Test_Switch1(t *testing.T) {
	a := 4
	switch a {
	case 1:
		fmt.Print("123")
	case 2, 3, 4:
		fmt.Print("456")
		fallthrough
	default:
		fmt.Print("789")
	}
}

// Test_Switch2 判斷型別
func Test_Switch2(t *testing.T) {
	var a interface{}
	a = 0
	switch v := a.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	default:
		_ = v
	}
}

type A struct {
}

func (a *A) Draw() {

}

func Test_Switch3(t *testing.T) {
	type Ccc interface {
		Draw()
	}

	type Bbb interface {
		Stop()
	}

	var v interface{}
	v = new(A)
	switch t := v.(type) {
	case *Ccc:
		fmt.Println("CCC")
	case *Bbb:
		fmt.Println("BBB")
	case *A:
		fmt.Println("A")
		t.Draw()
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
}

func Test_Switch4(t *testing.T) {
	d := time.Now()
	switch {
	case d.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}

func Test_Switch5(t *testing.T) {
	fn := func(value interface{}) {
		switch v := value.(type) {
		case int:
			fmt.Printf("%T\n", v) // 此處的v會是int
		case string, bool:
			fmt.Printf("%T\n", v) // 此處的v為interface, 因為有可能是bool 有可能是string
		default:
			fmt.Printf("Unexpected type %T\n", t)
		}
	}
	fn(123)     // int
	fn(123 > 0) // bool
	fn("123")   // string
}
