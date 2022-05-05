package example

import (
	"fmt"
	"reflect"
)

func ExNil () {
	// IsNil()和IsValid() -- 判断反射值的空和有效性

	// IsNil()
	// 返回值是否为 nil。如果值类型不是通道（channel）、函数、接口、map、指针或 切片时发生 panic，类似于语言层的v== nil操作

	// IsValid()
	// 判断值是否有效。 当值本身非法时，返回 false，例如 reflect Value不包含任何值，值为 nil 等

	//*int的空指针
	var a *int
	fmt.Println("var a *int:", reflect.ValueOf(a).IsNil()) // true

	//nil值
	fmt.Println("nil:", reflect.ValueOf(nil).IsValid()) // false

	//*int类型的空指针
	fmt.Println("(*int)(nil):", reflect.ValueOf((*int)(nil)).Elem().IsValid()) // false

	//实例化一个结构体
	s := struct {}{}

	//尝试从结构体中查找一个不存在的字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(s).FieldByName("").IsValid()) // false

	//尝试从结构体中查找一个不存在的方法
	fmt.Println("不存在的方法:", reflect.ValueOf(s).MethodByName("").IsValid()) // false

	//实例化一个map
	m := map[int]int{}

	//尝试从map中查找一个不存在的键
	fmt.Println("不存在的键:", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid()) // false
}
