package main

import "fmt"

func main() {
	a, b := 1, 2
	add(&a, &b)
	fmt.Println(a)

	c := &a // c的类型是 *int，c指向a的地址，*c就可以拿到a里边的东西
	d := &c // d的类型是 **int，a指向c的地址，d本身是个指针，它存的东西也是指针
	fmt.Println("d=", d, "*d=", *d, "**d=", **d)
	m := map[string]string{}
	mp1 := &m
	fmt.Println(mp1) // mp1的类型是 *map[string]string{}

	put(m)
	fmt.Println("*mp1: ", *mp1)

	f1 := add // f1 = func(int,int)
	f1(&a, &b)
	fmt.Println("f1,add = ", a)

	f1p1 := &f1
	(*f1p1)(&a, &b)
	fmt.Println("f1p1,add = ", a)

	{
		var nothing *int
		// *nothing = 1 // 注意：这里没有指向任何东西的int指针
		fmt.Println("nothing: ", nothing)
	}
	{
		var nothingMap map[string]string
		//nothingMap["a"] = "888"
		fmt.Println("nothingMap: ", nothingMap)
	}

	{
		// 切片是个特殊的可以给一个空的变量进行append的
		var nothingSlice []int
		//nothingSlice[0] = 100 //报错
		nothingSlice = append(nothingSlice, 100)
		fmt.Println(nothingSlice)
	}
}

func put(m map[string]string) {
	m["a"] = "AAA"
}

func add(a, b *int) {
	*a = *a + *b
}
