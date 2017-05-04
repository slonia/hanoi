package main

import (
	"fmt"
)

func main() {
	var disks int
	a := new(Tower)
	b := new(Tower)
	c := new(Tower)
	fmt.Println("Enter number of disks:")
	fmt.Scanf("%d", &disks)
	for i := 0; i < disks; i++ {
		a.AddDisk(i)
	}
	DrawTowers(a, b, c)
	move(a.Height(), a, b, c)
}

func move(level int, a *Tower, b *Tower, c *Tower) {
	if level == 0 {
		bHeight := b.Height()
		if bHeight != 0 {
			move(bHeight, b, a, c)
		}
	} else if level == 1 {
		c.AddDisk(a.GetDisk())
		move(0, a, b, c)
	} else if level%2 == 0 {
		// b.AddDisk(a.GetDisk())
		move(level-1, a, c, b)
	} else {
		// c.AddDisk(a.GetDisk())
		move(level-1, a, b, c)
	}
	DrawTowers(a, b, c)
}

func DrawTowers(a *Tower, b *Tower, c *Tower) {
	fmt.Println("A")
	a.Draw()
	fmt.Println("B")
	b.Draw()
	fmt.Println("C")
	c.Draw()
}
