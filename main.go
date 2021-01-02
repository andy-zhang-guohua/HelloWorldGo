package main // 要求这里报名必须是 main
import (
	"fmt"
	"reflect"
	"strconv"
)

type Element interface{}
type List [] Element

type Person struct {
	name string
	age int
}

// 定义了 String 方法，实现了 fmt.Stringer
func (p Person) String() string {
	return "(name: " + p.name + " - age: "+strconv.Itoa(p.age)+ " years)"
}

func main() {
	list := make(List, 3)
	list[0] = 1 // an int
	list[1] = "Hello" // a string
	list[2] = Person{"Dennis", 70}

	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}

	t := reflect.TypeOf(list[0])    // 得到类型的元数据,通过t我们能获取类型定义里面的所有元素
	v := reflect.ValueOf(list[0])

	fmt.Printf("type : %v\n", t)
	fmt.Printf("value : %v\n", v)
}