package main

func main() {
	var names []string

	persons := []string{}

}

type Person struct {
	name   string
	sex    string
	tall   float64
	age    float64
	weight float64
}

type calculator struct {
	a, b      int
	sum       int
	product   int
	quotient  int
	remainder int
}

func (c calculator) Add() int {
	return c.a + c.b
} {
type calculator struct {
	a, b      int
	sum       int
	product   int
	quotient  int
	remainder int
}

func (c *calculator) Add() int {
	c.sum = c.a + c.b
	return c.sum
}
