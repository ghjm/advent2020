package main

import (
	"fmt"
	"github.com/ghjm/advent2020/pkg/utils"
)

type coord3d struct {
	x int
	y int
	z int
}

type coord4d struct {
	x int
	y int
	z int
	w int
}

func forAllNeighbors3d(pos coord3d, f func(c coord3d)) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				if dx == 0 && dy == 0 && dz == 0 {
					continue
				}
				f(coord3d{pos.x + dx, pos.y + dy, pos.z + dz})
			}
		}
	}
}

func forAllNeighbors4d(pos coord4d, f func(c coord4d)) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dw := -1; dw <= 1; dw++ {
					if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
						continue
					}
					f(coord4d{pos.x + dx, pos.y + dy, pos.z + dz, pos.w + dw})
				}
			}
		}
	}
}

func iterate3d(state map[coord3d]struct{}) map[coord3d]struct{} {
	result := make(map[coord3d]struct{})
	open := make(map[coord3d]struct{})
	closed := make(map[coord3d]struct{})
	for k, _ := range state {
		open[k] = struct{}{}
	}
	for len(open) > 0 {
		var pos coord3d
		for k := range open {
			pos = k
			break
		}
		delete(open, pos)
		_, ok := closed[pos]
		if ok {
			continue
		}
		closed[pos] = struct{}{}
		_, active := state[pos]
		neighbors := 0
		forAllNeighbors3d(pos, func(c coord3d) {
			if active {
				_, ok = closed[c]
				if !ok {
					open[c] = struct{}{}
				}
			}
			_, ok = state[c]
			if ok {
				neighbors++
			}
		})
		if active && (neighbors == 2 || neighbors == 3) {
			result[pos] = struct{}{}
		} else if !active && neighbors == 3 {
			result[pos] = struct{}{}
		}
	}
	return result
}

func iterate4d(state map[coord4d]struct{}) map[coord4d]struct{} {
	result := make(map[coord4d]struct{})
	open := make(map[coord4d]struct{})
	closed := make(map[coord4d]struct{})
	for k, _ := range state {
		open[k] = struct{}{}
	}
	for len(open) > 0 {
		var pos coord4d
		for k := range open {
			pos = k
			break
		}
		delete(open, pos)
		_, ok := closed[pos]
		if ok {
			continue
		}
		closed[pos] = struct{}{}
		_, active := state[pos]
		neighbors := 0
		forAllNeighbors4d(pos, func(c coord4d) {
			if active {
				_, ok = closed[c]
				if !ok {
					open[c] = struct{}{}
				}
			}
			_, ok = state[c]
			if ok {
				neighbors++
			}
		})
		if active && (neighbors == 2 || neighbors == 3) {
			result[pos] = struct{}{}
		} else if !active && neighbors == 3 {
			result[pos] = struct{}{}
		}
	}
	return result
}

func main() {
	initialState := make(map[coord3d]struct{})
	y := 0
	err := utils.OpenAndReadAll("input17.txt", func(s string) error {
		for x := 0; x < len(s); x++ {
			if s[x] == '#' {
				initialState[coord3d{x, y, 0}] = struct{}{}
			}
		}
		y += 1
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part One\n")
	newState3d := initialState
	for i := 0; i < 6; i++ {
		newState3d = iterate3d(newState3d)
	}
	fmt.Printf("Count: %d\n", len(newState3d))
	fmt.Printf("Part Two\n")
	newState4d := make(map[coord4d]struct{})
	for pos := range initialState {
		newState4d[coord4d{pos.x, pos.y, pos.z, 0}] = struct{}{}
	}
	for i := 0; i < 6; i++ {
		newState4d = iterate4d(newState4d)
	}
	fmt.Printf("Count: %d\n", len(newState4d))
}

