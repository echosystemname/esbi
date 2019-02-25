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
		fmt.Println(i, string(v))
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
			if t.Tape[t.Cursor] == 0 {
				i = strings.Index(t.Source[i:], string(']')) + 1
			}
		case ']':
			if t.Tape[t.Cursor] != 0 {
				i = strings.LastIndex(t.Source[:i], string('['))
			}
		}

		if i+1 == len(src) {
			break
		}
		i++
		v = src[i]

	}
}

func main() {
	dat, _ := ioutil.ReadFile(os.Args[1])
	source := string(dat)

	tape := tape{}
	tape.Source = source
	tape.process()
}
