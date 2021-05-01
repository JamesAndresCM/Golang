package main

import (
	"fmt"
	"strings"
)

type t interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew\n", int(l))
}

func main() {
	martian := martian{}
	fmt.Println(martian.talk())

	laser := laser(8)
	fmt.Println(laser.talk())
}
