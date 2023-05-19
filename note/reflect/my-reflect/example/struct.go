package example

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
	Type int `json:"type" form:"type"`
}

func (s *Student) SetAge(age, t int) int {
	s.Age = age
	s.Type = t
	fmt.Println("set done")
	return s.Age
}

func ExStruct() {
	// 获取类型信息
	var stu Student
	typeOfStu := reflect.TypeOf(stu)
	fmt.Println(typeOfStu.Name(), typeOfStu.Kind(), typeOfStu.String()) // Student struct example.Student

	// 获取指针指向的元素类型
	stuPtr := &Student{}
	typeOfStuPrt := reflect.TypeOf(stuPtr)
	fmt.Println(typeOfStuPrt.Name(), typeOfStuPrt.Kind(), typeOfStuPrt.String()) //  ptr *example.Student

	// 获取指针类型的元素类型
	elem := typeOfStuPrt.Elem()
	fmt.Println(elem.Name(), elem.Kind(), elem.String()) // Student struct example.Student

	// 获取成员反射信息
	newStu := &Student{Name: "shiro", Age: 233, Type: 1}
	typeOfNewStuPtr := reflect.TypeOf(newStu)
	newStuElem := typeOfNewStuPtr.Elem()
	for i := 0; i < newStuElem.NumField(); i++ {
		field := newStuElem.Field(i)
		// ...
		// name: Type, tag: 'json:"type" form:"type"'
		fmt.Printf("name: %v, tag: '%v'\n", field.Name, field.Tag)
	}
	// 通过字段名, 找到字段类型信息
	if stuType, ok := newStuElem.FieldByName("Type"); ok {
		// {Type  int json:"type" form:"type" 24 [2] false} type
		// false -- 是否为一个嵌入式字段
		fmt.Println(stuType, stuType.Tag.Get("json"))
	}

	// 获取值信息
	rValue := reflect.ValueOf(newStu)
	rValue = rValue.Elem()
	fmt.Println(rValue.NumField())
	for i := 0; i < rValue.NumField(); i++ {
		field := rValue.Field(i)
		fmt.Println(field.Type(), field.Interface())
	}

	// 调用方法
	rValueNew := reflect.ValueOf(newStu)

	//获取到该结构体有多少个方法
	fmt.Println(rValueNew.NumMethod()) // 1

	//构造函数参数，传入两个整形值
	params := []reflect.Value{reflect.ValueOf(20), reflect.ValueOf(2)}

	//调用结构体的第一个方法Method(0)
	//注意:在反射值对象中方法索引的顺序并不是结构体方法定义的先后顺序
	//而是根据方法的ASCII码值来从小到大排序，所以Dec排在第一个，也就是Method(0)
	res := rValueNew.Method(0).Call(params)
	fmt.Println(res[0].Int())

	rValue = rValueNew.Elem()
	for i := 0; i < rValue.NumField(); i++ {
		//注:经过测试发现Field(i)的参数索引是从0开始的，
		//并且是按照定义的结构体的顺序来的，而不是按照字段名字的ASCii码值来的
		field := rValue.Field(i)
		fmt.Println(field.Type(), field.Interface())
	}
}
