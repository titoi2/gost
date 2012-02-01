package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

var stack Stack
var dict WordDict

func main() {
	stack.Init(StackSize)
	dict.Init()

	setWords()

	var sc TokenScanner
	for {
		fmt.Print("GOST>")
		s := readln()
		if string(s) == "\n" {
			break
		}
		sc.Init(s)
		for {
			token := sc.GetToken()
			if token == "" {
				break
			}

			f, err := strconv.Atof32(token)
			if err == nil {
				stack.Push(f)
			} else {
				fn, ok := dict.Search(token)
				if ok {
					fn()
				}
			}
		}

	}
}

func pop() (f float32, err string) {
	err = ""
	x := stack.Pop()
	switch i := x.(type) {
	case nil:
		err = "stack under flow"
	case int:
		f = float32(i)
	case float32:
		f = i
	case float64:
		f = float32(i)
	default:
		err = "stack pop type unknowm"
	}
	return
}

func push(f float32) string {
	err := ""
	x := stack.Push(f)
	if x != 0 {
		err = "stack over flow"
	}
	return err
}

func setWords() {
	fn := func() {
		n, err := pop()
		if err == "" {
			fmt.Println(n)
		} else {
			fmt.Println(err)
		}
	}
	dict.SetWord(".", fn)

	fn = func() {
		n1, err := pop()
		if err != "" {
			fmt.Println(err)
			return
		}
		n2, err := pop()
		if err != "" {
			fmt.Println(err)
			return
		}
		push(n2 + n1)
	}
	dict.SetWord("+", fn)

	fn = func() {
		n1, err := pop()
		if err != "" {
			fmt.Println(err)
			return
		}
		n2, err := pop()
		if err != "" {
			fmt.Println(err)
			return
		}
		push(n2 - n1)
	}
	dict.SetWord("-", fn)

	fn = func() {
		n1, err := pop()
		if err != "" {
			fmt.Println(err)
			return
		}
		n2, err := pop()
		if err != "" {
			fmt.Println(err)
			return
		}
		push(n2 * n1)
	}
	dict.SetWord("*", fn)

	fn = func() {
		n1, err := pop()
		if err != "" {
			fmt.Println(err)
			return
		}
		n2, err := pop()
		if err != "" {
			fmt.Println(err)
			return
		}
		push(n2 / n1)
		return
	}
	dict.SetWord("/", fn)
}

func readln() []byte {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadBytes('\n')
	if err != nil {
		return []byte{'\n'}
	}
	return input
}
