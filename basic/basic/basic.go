package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	bb = "kk"
	ss = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitalValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypedeDeduction() {
	var a, b, c, s = 3, 4, "abc", true
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, "abc", true
	fmt.Println(a, b, c, s)
}

func euler() {
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle()  {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	return c
}

func consts()  {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a * a + b * b))
	fmt.Println(filename, c)
	
}

func enums()  {
	const (
		cpp = iota
		_
		python
		golang
		java
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)	
}


func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	variableInitalValue()
	variableTypedeDeduction()
	variableShorter()
	fmt.Println(aa, bb, ss)
	euler()
	triangle()
	consts()
	enums()
}
