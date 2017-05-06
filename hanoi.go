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
	moves := 0
	board.Draw(moves)
	move(board.a.Height(), board.a, board.b, board.c, &moves)
}

func move(level int, a *Tower, b *Tower, c *Tower, moves *int) {
	if level == 1 {
		if a.Height() > 0 {
			c.AddDisk(a.GetDisk())
			*moves++
		}
	} else {
		move(level-1, a, c, b, moves)
		move(1, a, b, c, moves)
		move(level-1, b, a, c, moves)
	}
	board.Draw(*moves)
	time.Sleep(500 * time.Millisecond)
}
