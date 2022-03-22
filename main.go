package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"

	"lemonadeFactory/factory"
)

func main() {
	bottleSize := 3
	lf := factory.NewLemonadeFactory(bottleSize, factory.PLAIN)

	screen.Clear()
	for {
		screen.MoveTopLeft()
		lf.NewLemonade()

		fmt.Printf("Lemonades Produced: %d", lf.ProductionCount)
		time.Sleep(time.Second)
	}
}
