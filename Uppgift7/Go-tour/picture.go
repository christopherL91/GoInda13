package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	//the whole slice is now allocated!
	i := make([][]uint8, dy)
	for j := range i {
		i[j] = make([]uint8, dx)
	}

	for y_pos, row := range i {
		for x_pos := range row {
			row[x_pos] = uint8(x_pos ^ y_pos)
		}
	}
	return i
}

func main() {
	pic.Show(Pic)
}
