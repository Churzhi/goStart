package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	var hello string = "hello golang"
	var world = "world"
	fmt.Println(hello, world)
	小数 := 1.234
	fmt.Println(小数)
	var int3, int4 uint = 33, 44
	fmt.Println(int3, int4)

	var ho, ver float64 = 3, 4.123

	var sc = ho * ver //变量计算需要类型相等，不同类型进行计算会报错
	fmt.Println(ho * ver)
	fmt.Println(sc)

	var newname int
	fmt.Println("newname = ", newname)

	a1 := "hello"
	fmt.Println(reflect.TypeOf(a1))
	a2 := 3
	fmt.Println(reflect.TypeOf(a2))
	a3 := 3.0
	fmt.Println(reflect.TypeOf(a3))
	a4 := &a3
	fmt.Println(reflect.TypeOf(a4))

	var int6Max uint = math.MaxUint64
	fmt.Println(int6Max)
	var int7 int = int(int6Max)
	fmt.Println(int7)
}
