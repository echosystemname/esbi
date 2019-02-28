package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type tape struct {
	Source string
	Tape   [30000]byte
	Cursor int
	Loops  []loopPair
}

type loopPair struct {
	Open  int
	Close int
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
		fmt.Println(i+1, t.Cursor, t.Tape[:10])
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
				i = i + strings.Index(t.Source[i:], string(']')) + 1
				continue
			}
		case ']':
			//if t.Tape[t.Cursor] != 0 {
			i = strings.LastIndex(t.Source[:i], string('['))
			continue
			//}
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
	result := []loopPair{}

	for i, v := range source {
		if v == '[' {
			stack = append([]int{i}, stack...)
		} else if v == ']' {
			result = append(result, loopPair{stack[0], i})
			stack = stack[1:]
		}
	}

	tape.Loops = result

	tape.process()
}
