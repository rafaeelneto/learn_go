package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	dySlice := make([][]uint8, dy)
	
	for y := range dySlice {
		dxSlice := make([]uint8, dx)
		for x := range dxSlice {
			value := uint8(x*x+y*y)
			dxSlice[x] = value
		}
		dySlice[y] = dxSlice
	}
	
	return dySlice;
}

func main() {
	pic.Show(Pic)
}
