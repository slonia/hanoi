package main

import (
	"fmt"
	"time"
)

var board *Board

func main() {
	var disks int
	fmt.Println("Enter number of disks:")
	fmt.Scanf("%d", &disks)
	board = &Board{new(Tower), new(Tower), new(Tower), disks}
	board.Init()
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
	time.Sleep(500 * time.Millisecond)
}
