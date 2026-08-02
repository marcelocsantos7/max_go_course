package main

import "fmt"

type str string

func (text str) log() {
	fmt.Println(text)
}
