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
	starsCount := 0
	if t.Height() > 0 && t.Height() > level {
		starsCount = t.items[level] + 1
	}
	spaceCount := max - starsCount
	for i := 0; i < starsCount; i++ {
		str += "*"
	}
	for i := 0; i < spaceCount; i++ {
		str += " "
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
		level := height - i - 1
		s := b.a.DrawLevel(level, height)
		s += "    "
		s += b.b.DrawLevel(level, height)
		s += "    "
		s += b.c.DrawLevel(level, height)
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
