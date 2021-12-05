package main

import "fmt"

func main() {
	a := "hello"
	fmt.Println(a)
	aBytes := []byte(a)
	fmt.Println(aBytes)
	fmt.Println("修改切片内的内容")
	aBytes[0] = 'H'
	a = string(aBytes)
	fmt.Println(a)

	//UTF-8字符转换：rune()
	b := "您好"
	fmt.Println(b)
	fmt.Println("byte a:", []byte(a), len([]byte(a)))
	fmt.Println("rune a:", []rune(a), len([]rune(a)))
	bBytes := []rune(b)
	fmt.Println(bBytes)
	fmt.Println("修改切片内的内容")
	bBytes[0] = 'H'
	b = string(bBytes)
	fmt.Println(b)

}
