package main

func day3(input []string) (int, int, error) {
	var part1, part2 int

	width := len(input[0]) - 1 // the terrain is equally wide
	slope1 := slope{
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
		if pointer.col+slope1.run > width {
			pointer.col = slope1.run - (width - pointer.col + 1)
		} else {
			pointer.col = pointer.col + slope1.run
		}
		pointer.row = i + slope1.rise
		// row := input[pointer.row]
		if string(input[pointer.row][pointer.col]) == "#" {
			part1 = part1 + 1
			// row = row[:pointer.col] + "X" + row[pointer.col+1:]
		} else {
			// row = row[:pointer.col] + "O" + row[pointer.col+1:]
		}
		// fmt.Printf("%s\t%+v\n", row, pointer)
	}

	slopes := []slope{
		{run: 1, rise: 1},
		{run: 3, rise: 1},
		{run: 5, rise: 1},
		{run: 7, rise: 1},
		{run: 1, rise: 2},
	}

	var trees []int
	for _, slope := range slopes {
		pointer = point{row: 0, col: 0}
		var t int
		for i := slope.rise; i < len(input); i = i + slope.rise {
			if pointer.col+slope.run > width {
				pointer.col = slope.run - (width - pointer.col + 1)
			} else {
				pointer.col = pointer.col + slope.run
			}
			pointer.row = i
			row := input[pointer.row]
			if string(input[pointer.row][pointer.col]) == "#" {
				t = t + 1
				row = row[:pointer.col] + "X" + row[pointer.col+1:]
			} else {
				row = row[:pointer.col] + "O" + row[pointer.col+1:]
			}
		}
		trees = append(trees, t)
	}
	part2 = trees[0]
	for i := 1; i < len(trees); i++ {
		part2 = part2 * trees[i]
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
