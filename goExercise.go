package main

import (
	"fmt"
)

	type Direction int

	const (
		None Direction = iota
		North
		East
		South
		West
	)

func DirectionToString(d Direction) string {
	switch d{
	case 1:
		return "North"
	case 2:
		return "East"
	case 3:
		return "South"
	case 4:
		return "West"
	default:
		return "None"
	}
}

	func GetDirection(angle float64) Direction {
	switch {
	case angle >= 315:
		return 1
	case angle >=0 && angle < 45:
		return 1
	case angle >=45 && angle < 135:
		return 2
	case angle >=135 && angle < 225:
		return 3
	case angle >=225 && angle < 315:
		return 4
	default:
		return None
	}
}

func main() {
	fmt.Println(DirectionToString(GetDirection(38.3)))
	fmt.Println(DirectionToString(GetDirection(235.8)))
	fmt.Println(DirectionToString(GetDirection(94.2)))
	fmt.Println(DirectionToString(GetDirection(-30)))
}
