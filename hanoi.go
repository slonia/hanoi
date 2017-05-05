package main

import (
	"fmt"
)

var board *Board

func main() {
	var disks int
	board = &Board{new(Tower), new(Tower), new(Tower)}
	fmt.Println("Enter number of disks:")
	fmt.Scanf("%d", &disks)
	for i := disks - 1; i > -1; i-- {
		board.a.AddDisk(i)
	}
	board.Draw()
	move(board.a.Height(), board.a, board.b, board.c)
}

func move(level int, a *Tower, b *Tower, c *Tower) {
	if level == 1 {
		c.AddDisk(a.GetDisk())
	} else {
		// b.AddDisk(a.GetDisk())
		move(level-1, a, c, b)
		move(1, a, b, c)
		move(level-1, b, a, c)
		// board.Draw()
	}
	board.Draw()
}
