package main

import (
	"fmt"
	"os"
	"strings"
)

func echo1() {
	var s, seq string
	for _, args := range os.Args[1:] {
		s += seq + args
		seq = " "
	}
	fmt.Println(s)
}

func echo2() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func ch1_prac1() {
	fmt.Println(os.Args[0])
}

func ch1_prac2() {
	for idx, args := range os.Args[1:] {
		fmt.Println("%d, %s", idx, args)
	}
}

func ch1_prac3() {
	//pass
}
func main() {
	echo1()
	echo2()
}
