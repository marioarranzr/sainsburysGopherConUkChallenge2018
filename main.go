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

// This was my approach. Coded it in 2m 44s. I finished in top10! ✌️
func calculateHouses(input string) int {
	type pos struct {
		x, y int
	}

	v := make(map[pos]bool)
	p := pos{0, 0}
	v[p] = true

	for _, n := range []rune(input) {
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
// This was the winner's approach. Coded it in 28s. 👏️
func calculateHousesWinner(input string) int {
	var c[2]int
	g := map[[2]int]int{c:0}
	for _, f := range input {
		c[f/64]+=int(f+10)&8/4-1
		g[c]++
	}

	return len(g)
}
