package main

import (
	"fmt"
	"github.com/ghjm/advent2020/pkg/utils"
	"gonum.org/v1/gonum/mat"
	"regexp"
)

type command struct {
	instr string
	param int
}

type direction struct {
	dx int
	dy int
}

var cardinalDirections = map[string]direction{
	"N": {dx: 0, dy: -1},
	"S": {dx: 0, dy: 1},
	"W": {dx: -1, dy: 0},
	"E": {dx: 1, dy: 0},
}

var angleToDirection = map[int]string{
	0: "N",
	90: "E",
	180: "S",
	270: "W",
}

var rotate90 = map[int]*mat.Dense{
	90: mat.NewDense(2, 2, []float64{
		0, -1,
		1, 0,
	}),
	180: mat.NewDense(2, 2, []float64{
		-1, 0,
		0, -1,
	}),
	270: mat.NewDense(2, 2, []float64{
		0, 1,
		-1, 0,
	}),
}

func main() {
	commands := make([]command, 0)
	re := regexp.MustCompile(`^([NSEWLRF])(\d+)$`)
	err := utils.OpenAndReadAll("input12.txt", func(s string) error {
		m := re.FindStringSubmatch(s)
		cmd := command{
			instr: m[1],
			param: utils.AtoiOrPanic(m[2]),
		}
		commands = append(commands, cmd)
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part One\n")
	curAngle := 90
	curX := 0
	curY := 0
	for _, cmd := range commands {
		card, ok := cardinalDirections[cmd.instr]
		if ok {
			curX += card.dx * cmd.param
			curY += card.dy * cmd.param
		} else if cmd.instr == "R" || cmd.instr == "L" {
			sign := 0
			if cmd.instr == "R" {
				sign = 1
			} else {
				sign = -1
			}
			curAngle = ((curAngle + cmd.param * sign) + 360) % 360
			_, ok := angleToDirection[curAngle]
			if !ok {
				panic(fmt.Sprintf("Instruction %s%d resulted in non cardinal angle", cmd.instr, cmd.param))
			}
		} else if cmd.instr == "F" {
			dir := cardinalDirections[angleToDirection[curAngle]]
			curX += dir.dx * cmd.param
			curY += dir.dy * cmd.param
		}
	}
	fmt.Printf("Manhattan distance: %d\n", utils.Abs(curX) + utils.Abs(curY))

	fmt.Printf("Part Two\n")
	boatX := 0
	boatY := 0
	waypointX := 10
	waypointY := -1
	for _, cmd := range commands {
		card, ok := cardinalDirections[cmd.instr]
		if ok {
			waypointX += card.dx * cmd.param
			waypointY += card.dy * cmd.param
		} else if cmd.instr == "R" || cmd.instr == "L" {
			rot := cmd.param
			if cmd.instr == "L" {
				rot *= -1
			}
			rot = (rot + 360) % 360
			if rot != 0 {
				rotMat, ok := rotate90[rot]
				if !ok {
					panic(fmt.Sprintf("Instruction %s%d resulted in non cardinal angle", cmd.instr, cmd.param))
				}
				oldVec := mat.NewVecDense(2, []float64{float64(waypointX), float64(waypointY)})
				newVec := mat.NewVecDense(2, make([]float64, 2))
				newVec.MulVec(rotMat, oldVec)
				waypointX = int(newVec.AtVec(0))
				waypointY = int(newVec.AtVec(1))
			}
		} else if cmd.instr == "F" {
			boatX += waypointX * cmd.param
			boatY += waypointY * cmd.param
		}
	}
	fmt.Printf("Manhattan distance: %d\n", utils.Abs(boatX) + utils.Abs(boatY))
}

