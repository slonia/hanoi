package main

import "fmt"

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

func (t *Tower) Draw() {
	for _, el := range t.items {
		fmt.Printf("%d, ", el+1)
	}
	fmt.Println("")
}

// code for GameBoard

type Board struct {
	a *Tower
	b *Tower
	c *Tower
}

func (b *Board) Draw() {
	fmt.Println("===================")
	fmt.Println("A")
	b.a.Draw()
	fmt.Println("B")
	b.b.Draw()
	fmt.Println("C")
	b.c.Draw()
	fmt.Println("===================")
}
