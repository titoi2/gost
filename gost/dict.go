package main

import (
	"fmt"
)

type WordParam struct {
	function func()
}

type WordDict struct {
	dict map[string]*WordParam
}

func (d *WordDict) Init() {
	d.dict = make(map[string]*WordParam)
}

func (d *WordDict) SetWord(name string, fn func()) {
	p := new(WordParam)
	p.function = fn
	d.dict[name] = p
}

func (d *WordDict) Search(name string) (fn func(), ok bool) {
	wp, ok := d.dict[name]
	if ok {
		fn = wp.function
	} else {
		fmt.Println("not found")
	}
	return
}
