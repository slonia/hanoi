package main

import "fmt"

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
		fmt.Printf("%d ", el+1)
	}
	fmt.Println("")
}
