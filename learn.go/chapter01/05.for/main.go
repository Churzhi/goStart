package main

import "fmt"

func main() {
	fmt.Println("round 1")
	for i := 0; i < 5; i++ {
		fmt.Println("您好，golang")
	}

	fmt.Println("round 2")
	j := 0
	for ; j < 5; j++ {
		fmt.Println("round 2, 您好，golang")
	}

	fmt.Println("round 3,for循环简写")
	k := 0
	for k < 5 {
		fmt.Println("round 3, 您好，golang")
		k++
	}
	fmt.Println("round 4")
	l := 0
	for {
		fmt.Println("round 4, hello golang")
		l++
		if l >= 3 {
			break
		}

	}
	fmt.Println("round 5")
	m := 0
	for {
		fmt.Println("round 5, hello golang")
		m++
		if m >= 10 {
			break
		}
		if m%2 == 0 {
			continue
		}
		fmt.Println("round 5, 练习跳过", m)
	}
}
