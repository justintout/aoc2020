package main

func day3(input []string) (int, int, error) {
	var part1, part2 int

	width := len(input[0]) - 1 // the terrain is equally wide
	slope := slope{
		rise: 1,
		run:  3,
	}
	pointer := point{
		row: 0,
		col: 0,
	}

	// fmt.Println(input[0])
	for i := 0; i < len(input)-1; i++ {
		// wrap around to the start if we hit the edge
		if pointer.col+slope.run > width {
			pointer.col = slope.run - (width - pointer.col + 1)
		} else {
			pointer.col = pointer.col + slope.run
		}
		pointer.row = i + slope.rise
		// row := input[pointer.row]
		if string(input[pointer.row][pointer.col]) == "#" {
			part1 = part1 + 1
			// row = row[:pointer.col] + "X" + row[pointer.col+1:]
		} else {
			// row = row[:pointer.col] + "O" + row[pointer.col+1:]
		}
		// fmt.Printf("%s\t%+v\n", row, pointer)
	}

	return part1, part2, nil
}

type slope struct {
	rise int
	run  int
}

type point struct {
	row int
	col int
}
