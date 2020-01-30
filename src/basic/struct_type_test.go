package basic

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	ID   string
	NAME string
	AGE  int
}

func TestInitEmplyeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{NAME: "Mike", AGE: 30}
	e2 := new(Employee)
	e2.ID = "2"
	e2.NAME = "Rose"
	e2.AGE = 40
	t.Log(e)
	t.Log(e1)
	t.Log(e1.ID)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)
}

//第一种定义方式在实例对应方法被调用时,实例的成员会进行值复制
func (e Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.NAME))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.ID, e.NAME, e.AGE)
}

//(推荐)第二种定义方式可以避免实例成员的内存复制
// func (e *Employee) String() string {
// 	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.NAME))
// 	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.ID, e.NAME, e.AGE)
//}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.NAME))
	t.Log(e.String())
}
