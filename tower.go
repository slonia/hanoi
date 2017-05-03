package main

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
