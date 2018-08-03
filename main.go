package main

import (
	"flag"
	"fmt"
)

func main() {
	inputString := flag.String("inputString", "",
		"char chain to analyse")
	flag.Parse()
	fmt.Printf("Given the instructions \"%s\" the delivery man visit %v diferent house(s)",
		*inputString, calculateHouses(*inputString))
}

// This was my approach. Codded it in 2m 44s. I finished in top10! ✌️
func calculateHouses(s string) int {
	type pos struct {
		x, y int
	}

	v := make(map[pos]bool)
	p := pos{0, 0}
	v[p] = true

	for _, n := range []rune(s) {
		switch n {
		case '>':
			p.x++
		case '<':
			p.x--
		case '^':
			p.y++
		case 'v':
			p.y--
		}
		v[p] = true
	}

	return len(v)
}
