package main

import (
	"fmt"
)

func main() {
	var disks int
	a := new(Tower)
	// b := new(Tower)
	// c := new(Tower)
	fmt.Println("Enter number of disks:")
	fmt.Scanf("%d", &disks)
	for i := 0; i < disks; i++ {
		a.AddDisk(i)
	}
	res := a.GetDisk()
	fmt.Printf("Last disk is: %d\n", res)
}
