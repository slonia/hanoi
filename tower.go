package main

import (
	"fmt"
	"os"
	"os/exec"
)

// code for one Tower

type Tower struct {
	items []int
}

func (t *Tower) AddDisk(val int) {
	t.items = append(t.items, val)
}

func (t *Tower) GetDisk() int {
	lastIndex := len(t.items) - 1
	el := t.items[lastIndex]
	t.items = t.items[:lastIndex]
	return el
}

func (t *Tower) Height() int {
	return len(t.items)
}

func (t *Tower) DrawLevel(level int, max int) string {
	str := ""
	// fmt.Printf("level: %d, items: %d, max: %d\n", level, len(t.items), max)
	if len(t.items) > 0 && len(t.items) > level {
		for i := 0; i < t.items[level]+1; i++ {
			str += "*"
		}
		for i := 0; i < max-t.items[level]-1; i++ {
			str += " "
		}
	} else {
		for i := 0; i < max; i++ {
			str += " "
		}
	}
	return str
}

// code for GameBoard

type Board struct {
	a      *Tower
	b      *Tower
	c      *Tower
	height int
}

func (b *Board) Draw(moves int) {
	clear()
	height := b.height
	for i := 0; i < height; i++ {
		s := b.a.DrawLevel(height-i-1, height)
		s += "    "
		s += b.b.DrawLevel(height-i-1, height)
		s += "    "
		s += b.c.DrawLevel(height-i-1, height)
		fmt.Println(s)
	}
	line := ""
	for i := 0; i < 3*height+8; i++ {
		line += "="
	}
	fmt.Println(line)
	fmt.Printf("Moves: %d\n", moves)
}

func (b *Board) Init() {
	for i := b.height - 1; i > -1; i-- {
		b.a.AddDisk(i)
	}
}

// general

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
