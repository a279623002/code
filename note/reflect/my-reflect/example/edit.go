package example

import (
	"fmt"
	"reflect"
)

func edit1() {
	//声明整形变量a并赋初值
	var a int = 1024

	//获取变量a的反射值对象
	rValue := reflect.ValueOf(a)

	//尝试将a修改为1(此处会崩溃)
	rValue.SetInt(1)
}

func edit2() {
	//声明整形变量a并赋初值
	var a int = 1024

	//获取变量a的反射值对象
	rValue := reflect.ValueOf(&a)

	//取出a地址的元素(a的值)
	rValue = rValue.Elem()
	fmt.Println(rValue.Int())

	//尝试将a修改为1
	rValue.SetInt(1)

	//打印a的值
	fmt.Println(rValue.Int())
}


func edit3() {
	type dog struct {
		legCount int
	}

	//获取dog实例的反射值对象
	valueOfDog := reflect.ValueOf(&dog{})

	valueOfDog = valueOfDog.Elem()

	//获取legCount字段的值
	vLegCount := valueOfDog.FieldByName("legCount")

	//尝试设置legCount的值(这里会发生崩溃)
	vLegCount.SetInt(4)
}

func edit4() {
	type dog struct {
		LegCount int
	}

	//获取dog实例的反射值对象
	valueOfDog := reflect.ValueOf(&dog{})

	valueOfDog = valueOfDog.Elem()

	//获取legCount字段的值
	vLegCount := valueOfDog.FieldByName("LegCount")

	//尝试设置legCount的值(这里会发生崩溃)
	vLegCount.SetInt(4)

	fmt.Println(vLegCount.Int())
}

func ExEdit() {
	// 值可修改条件之一：可被寻址
	//edit1() // panic: reflect: reflect.Value.SetInt using unaddressable value
	// 报错意思是:SetInt正在使用一个不能被寻址的值。从 reflect.ValueOf 传入的是 a 的值，而不是 a 的地址
	edit2()

	// 值可修改条件之一:被导出
	//edit3() // panic: reflect: reflect.Value.SetInt using value obtained using unexported field
	// 报错的意思是：SetInt() 使用的值来自于一个未导出的字段
	edit4()
}
