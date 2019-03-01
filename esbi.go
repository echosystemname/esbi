package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

type tape struct {
	Source string
	Tape   [30000]byte
	Cursor int
	OpenLoops map[int]int
	CloseLoops map[int]int
}

func (t *tape) incCursor() {
	t.Cursor++
}

func (t *tape) decCursor() {
	t.Cursor--
}

func (t *tape) incVal() {
	t.Tape[t.Cursor]++
}

func (t *tape) decVal() {
	t.Tape[t.Cursor]--
}

func (t *tape) printVal() {
	fmt.Print(string(t.Tape[t.Cursor]))
}

func (t *tape) scanVal() {
	reader := bufio.NewReader(os.Stdin)
	char, _, _ := reader.ReadRune()
	t.Tape[t.Cursor] = byte(char)
}

func (t *tape) process() {
	src := []rune(t.Source)
	i := 0
	v := src[0]

	for {
		v = src[i]
		switch rune(v) {
		case '>':
			t.incCursor()
		case '<':
			t.decCursor()
		case '+':
			t.incVal()
		case '-':
			t.decVal()
		case '.':
			t.printVal()
		case ',':
			t.scanVal()
		case '[':
			if t.Tape[t.Cursor] == byte(0) {
				i = t.OpenLoops[i] + 1
				continue
			}
		case ']':
			i = t.CloseLoops[i]
			continue
		}
		if i+1 == len(src) {
			break
		}
		i++

	}
}

func main() {
	dat, _ := ioutil.ReadFile(os.Args[1])
	source := string(dat)

	tape := tape{}
	tape.Source = source

	stack := []int{}

	openLoops := make(map[int]int)
	closeLoops := make(map[int]int)
	for i, v := range source {
		if v == '[' {
			stack = append([]int{i}, stack...)
		} else if v == ']' {

			openLoops[stack[0]] = i
			closeLoops[i] = stack[0]

			stack = stack[1:]
		}
	}

	tape.OpenLoops = openLoops
	tape.CloseLoops = closeLoops

	tape.process()
}
