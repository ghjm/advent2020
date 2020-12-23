package main

import (
	"fmt"
	"github.com/ghjm/advent2020/pkg/utils"
	"math"
	"strconv"
	"strings"
)

type orientedTile = struct {
	name        string
	edges       map[string]int64
	constraints map[string]struct{}
	transforms  []string
}

type orientReq = struct {
	cur, wanted string
}

func getEdgeID(edge []bool) int64 {
	var binStr string
	for _, b := range edge {
		if b {
			binStr += "1"
		} else {
			binStr += "0"
		}
	}
	i1, err := strconv.ParseInt(binStr, 2, 64)
	if err != nil {
		panic(err)
	}
	binStr = ""
	for i := len(edge)-1; i >= 0; i-- {
		b := edge[i]
		if b {
			binStr += "1"
		} else {
			binStr += "0"
		}
	}
	i2, err := strconv.ParseInt(binStr, 2, 64)
	if err != nil {
		panic(err)
	}
	if i1 < i2 {
		return i1
	} else {
		return i2
	}
}

func nullOp(ot *orientedTile, check bool) bool {
	return true
}

func flipVerticalOp(ot *orientedTile, check bool) bool {
	_, ok1 := ot.constraints["top"]
	_, ok2 := ot.constraints["bottom"]
	if ok1 || ok2 {
		return false
	}
	if !check {
		ot.transforms = append(ot.transforms, "flipvert")
		edge := ot.edges["top"]
		ot.edges["top"] = ot.edges["bottom"]
		ot.edges["bottom"] = edge
	}
	return true
}

func flipHorizontalOp(ot *orientedTile, check bool) bool {
	_, ok1 := ot.constraints["left"]
	_, ok2 := ot.constraints["right"]
	if ok1 || ok2 {
		return false
	}
	if !check {
		ot.transforms = append(ot.transforms, "fliphoriz")
		edge := ot.edges["left"]
		ot.edges["left"] = ot.edges["right"]
		ot.edges["right"] = edge
	}
	return true
}

func rotateOp(steps int) func(ot *orientedTile, check bool) bool {
	return func(ot *orientedTile, check bool) bool {
		if len(ot.constraints) > 0 {
			return false
		}
		if !check {
			for i := 0; i < steps; i++ {
				ot.transforms = append(ot.transforms, "rotate")
				topEdge := ot.edges["top"]
				ot.edges["top"] = ot.edges["left"]
				ot.edges["left"] = ot.edges["bottom"]
				ot.edges["bottom"] = ot.edges["right"]
				ot.edges["right"] = topEdge
			}
		}
		return true
	}
}

var orientOps = map[orientReq][]func(ot *orientedTile, check bool) bool{
	{"top", "top"}: {nullOp},
	{"top", "left"}: {rotateOp(3)},
	{"top", "bottom"}: {rotateOp(2), flipVerticalOp},
	{"top", "right"}: {rotateOp(1)},
	{"left", "top"}: {rotateOp(1)},
	{"left", "left"}: {nullOp},
	{"left", "bottom"}: {rotateOp(3)},
	{"left", "right"}: {rotateOp(2), flipHorizontalOp},
	{"bottom", "top"}: {rotateOp(2), flipVerticalOp},
	{"bottom", "left"}: {rotateOp(1)},
	{"bottom", "bottom"}: {nullOp},
	{"bottom", "right"}: {rotateOp(3)},
	{"right", "top"}: {rotateOp(3)},
	{"right", "left"}: {rotateOp(2), flipHorizontalOp},
	{"right", "bottom"}: {rotateOp(1)},
	{"right", "right"}: {nullOp},
}

func orient(ot *orientedTile, reqs map[string]int64, check bool) bool {
	for reqSide, reqEdge := range reqs {
		if ot.edges[reqSide] != reqEdge {
			var curSide string
			for otSide, otEdge := range ot.edges {
				if otEdge == reqEdge {
					curSide = otSide
					break
				}
			}
			if curSide == "" {
				return false
			}
			possibleOps := orientOps[orientReq{curSide, reqSide}]
			success := false
			for i := range possibleOps {
				if possibleOps[i](ot, check) {
					success = true
					break
				}
			}
			if !success {
				return false
			}
		}
	}
	return true
}

var neighbors = []struct{
	dx int
	dy int
	mySide string
	facingSide string}{
	{-1, 0, "left", "right"},
	{1, 0, "right", "left"},
	{0, -1, "top", "bottom"},
	{0, 1, "bottom", "top"},
}

func checkNeighbors(fullMap [][]*orientedTile, x int, y int, ot *orientedTile) (bool, map[*orientedTile]map[string]int64) {
	orientations := make(map[*orientedTile]map[string]int64)
	workable := true
	for _, n := range neighbors {
		nx := x + n.dx
		ny := y + n.dy
		if ny < 0 || ny >= len(fullMap) || nx < 0 || nx >= len(fullMap[0]) {
			continue
		}
		neigh := fullMap[ny][nx]
		if neigh == nil {
			continue
		}
		neighborOrient := map[string]int64{n.facingSide: ot.edges[n.mySide]}
		if !orient(neigh, neighborOrient, true) {
			workable = false
			break
		}
	}
	return workable, orientations
}

func transformTile(origTile [][]bool, transform string) [][]bool {
	size := len(origTile)
	if size != len(origTile[0]) {
		panic("tile not square")
	}
	newTile := make([][]bool, size)
	for i := 0; i < size; i++ {
		newTile[i] = make([]bool, size)
	}
	transfunc := map[string]func(origX, origY int) (int, int){
		"flipvert": func(origX, origY int) (int, int) {
			return size - origX - 1, origY
		},
		"fliphoriz": func(origX, origY int) (int, int) {
			return origX, size - origY - 1
		},
		"rotate": func(origX, origY int) (int, int) {
			return size - origY - 1, origX
		},
	}[transform]
	for origY := range origTile {
		for origX := range origTile[0] {
			newX, newY := transfunc(origX, origY)
			newTile[newY][newX] = origTile[origY][origX]
		}
	}
	return newTile
}

func main() {
	tiles := make(map[string][][]bool)
	var curTileName string
	var curTile [][]bool
	err := utils.OpenAndReadAll("input20_test.txt", func(s string) error {
		if s == "" {
			return nil
		} else if strings.HasPrefix(s, "Tile") {
			if curTile != nil {
				tiles[curTileName] = curTile
			}
			curTileName = s[5:9]
			curTile = make([][]bool, 0)
			return nil
		} else {
			curRow := make([]bool, 0)
			for i := range s {
				curRow = append(curRow, s[i] == '#')
			}
			curTile = append(curTile, curRow)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	if curTile != nil {
		tiles[curTileName] = curTile
	}
	fmt.Printf("Part One\n")
	edges := make(map[int64][]string)
	originalTiles := make(map[string]orientedTile)
	for tile, tileMap := range tiles {
		topEdge := getEdgeID(tileMap[0])
		bottomEdge := getEdgeID(tileMap[len(tileMap)-1])
		leftSide := make([]bool, 0)
		rightSide := make([]bool, 0)
		for _, row := range tileMap {
			leftSide = append(leftSide, row[0])
			rightSide = append(rightSide, row[len(row)-1])
		}
		leftEdge := getEdgeID(leftSide)
		rightEdge := getEdgeID(rightSide)
		for _, edge := range []int64{topEdge, leftEdge, bottomEdge, rightEdge} {
			_, ok := edges[edge]
			if !ok {
				edges[edge] = make([]string, 0)
			}
			edges[edge] = append(edges[edge], tile)
		}
		originalTiles[tile] = orientedTile{
			name:        tile,
			edges:       map[string]int64{"top": topEdge, "left": leftEdge, "bottom": bottomEdge, "right": rightEdge},
			constraints: make(map[string]struct{}),
			transforms:  make([]string, 0),
		}
	}
	outerTiles := make(map[string]struct{})
	cornerTiles := make(map[string]struct{})
	part2Unsolvable := false
	for _, tileList := range edges {
		if len(tileList) == 1 {
			_, ok := outerTiles[tileList[0]]
			if ok {
				cornerTiles[tileList[0]] = struct{}{}
				delete(outerTiles, tileList[0])
			} else {
				outerTiles[tileList[0]] = struct{}{}
			}
		} else if len(tileList) > 2 {
			part2Unsolvable = true
		}
	}
	prod := int64(1)
	for ct := range cornerTiles {
		tileInt := utils.AtoiOrPanic64(ct)
		prod *= tileInt
	}
	fmt.Printf("Product of corners: %d\n", prod)

	fmt.Printf("Part Two\n")
	if part2Unsolvable {
		panic("part 2 unsolvable")
	}
	sizeF := math.Sqrt(float64(len(tiles)))
	if sizeF != float64(int(sizeF)) {
		panic("size is not a square")
	}
	size := int(sizeF)
	var fullMap [][]*orientedTile

	complete := false
	CORNERS:
	for firstTileName := range cornerTiles {
		myTiles := make(map[string]orientedTile)
		for k, v := range originalTiles {
			newTile := orientedTile{
				name:        v.name,
				edges:       v.edges,
				constraints: make(map[string]struct{}),
				transforms:  make([]string, 0),
			}
			myTiles[k] = newTile
		}
		fullMap = make([][]*orientedTile, size)
		for i := 0; i < size; i++ {
			fullMap[i] = make([]*orientedTile, size)
		}
		firstTile := myTiles[firstTileName]
		outerSides := make([]int64, 0)
		for _, edge := range firstTile.edges {
			if len(edges[edge]) == 1 {
				outerSides = append(outerSides, edge)
			}
		}
		if len(outerSides) != 2 {
			panic("corner not a corner")
		}
		if !orient(&firstTile, map[string]int64{"top": outerSides[0], "left": outerSides[1]}, false) {
			if !orient(&firstTile, map[string]int64{"top": outerSides[1], "left": outerSides[0]}, false) {
				continue CORNERS
			}
		}
		delete(myTiles, firstTileName)
		fullMap[0][0] = &firstTile
		for y := range fullMap {
			for x := range fullMap[0] {
				if y == 0 && x == 0 {
					continue
				}
				var anchor *orientedTile
				var anchorSide string
				var anchorAltSide string
				var mySide string
				if x == 0 {
					anchor = fullMap[y-1][x]
					anchorSide = "bottom"
					anchorAltSide = "top"
					mySide = "top"
				} else {
					anchor = fullMap[y][x-1]
					anchorSide = "right"
					anchorAltSide = "left"
					mySide = "left"
				}
				found := false
				for tile, ot := range myTiles {
					testOt := orientedTile{
						name:        ot.name,
						edges:       ot.edges,
						constraints: make(map[string]struct{}),
						transforms:  make([]string, len(ot.transforms)),
					}
					for k := range ot.constraints {
						testOt.constraints[k] = struct{}{}
					}
					for i := range ot.transforms {
						testOt.transforms[i] = ot.transforms[i]
					}
					if anchor.name == "3079" && tile == "2311" {
						fmt.Printf("foo\n")
					}
					edgeFound := false
					for _, edge := range testOt.edges {
						if edge == anchor.edges[anchorSide] {
							edgeFound = true
							break
						}
					}
					if !edgeFound {
						for _, edge := range testOt.edges {
							if edge == anchor.edges[anchorAltSide] {
								edgeFound = true
								break
							}
						}
						if edgeFound {
							if !orient(anchor, map[string]int64{anchorSide: anchor.edges[anchorAltSide]}, true) {
								continue
							}
						}
					}
					var workable bool
					var orientations map[*orientedTile]map[string]int64
					if !orient(&testOt, map[string]int64{mySide: anchor.edges[anchorSide]}, false) {
						continue
					}
					testOt.constraints[mySide] = struct{}{}
					workable, orientations = checkNeighbors(fullMap, x, y, &testOt)
					if !workable {
						if flipVerticalOp(&testOt, false) {
							workable, orientations = checkNeighbors(fullMap, x, y, &testOt)
						}
					}
					if !workable {
						if flipHorizontalOp(&ot, false) {
							workable, orientations = checkNeighbors(fullMap, x, y, &testOt)
						}
					}
					if !workable {
						continue
					}
					for neigh := range orientations {
						if !orient(neigh, orientations[neigh], false) {
							panic("orientation failed after succeeding")
						}
					}
					fullMap[y][x] = &testOt
					for _, n := range neighbors {
						nx := x + n.dx
						ny := y + n.dy
						if ny < 0 || ny >= len(fullMap) || nx < 0 || nx >= len(fullMap[0]) {
							continue
						}
						if fullMap[ny][nx] != nil {
							ot.constraints[n.mySide] = struct{}{}
							fullMap[ny][nx].constraints[n.facingSide] = struct{}{}
						}
					}
					delete(myTiles, tile)
					found = true
					break
				}
				if !found {
					continue CORNERS
				}
			}
		}
		complete = true
		break
	}
	if !complete {
		panic("Failed to solve map")
	}
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			for _, n := range neighbors {
				ny := y + n.dy
				nx := x + n.dx
				if ny < 0 || nx < 0 || ny >= size || nx >= size {
					continue
				}
				myEdge := fullMap[y][x].edges[n.mySide]
				facingEdge := fullMap[ny][nx].edges[n.facingSide]
				if myEdge != facingEdge {
					panic(fmt.Sprintf("Map solution is wrong: y=%d x=%d has %s=%d but y=%d x=%d has %s=%d",
						y, x, n.mySide, myEdge, ny, nx, n.facingSide, facingEdge))
				}
			}
		}
	}

	finalOTs := make(map[string]*orientedTile)
	for y := range fullMap {
		for x := range fullMap[0] {
			ot := fullMap[y][x]
			finalOTs[ot.name] = ot
			fmt.Printf("%s %s        ", ot.name, strings.Join(ot.transforms, ","))
		}
		fmt.Printf("\n")
	}
	transformedTiles := make(map[string][][]bool)
	for tileName := range tiles {
		ot := finalOTs[tileName]
		tile := tiles[tileName]
		for _, transform := range ot.transforms {
			tile = transformTile(tile, transform)
		}
		transformedTiles[tileName] = tile
	}

	var tileSize int
	for _, tile := range tiles {
		tileSize = len(tile)
		break
	}
	tileSize -= 2
	bigSize := size * tileSize
	bigMap := make([][]bool, bigSize)
	anyTile := false
	for outerY := 0; outerY < size; outerY++ {
		for innerY := 0; innerY < tileSize; innerY++ {
			bigY := outerY*tileSize + innerY
			bigMap[bigY] = make([]bool, bigSize)
			for outerX := 0; outerX < size; outerX++ {
				tileName := fullMap[outerY][outerX].name
				if !anyTile {
					anyTile = true
				}
				for innerX := 0; innerX < tileSize; innerX++ {
					bigX := outerX*tileSize + innerX
					bigMap[bigY][bigX] = transformedTiles[tileName][innerY+1][innerX+1]
				}
			}
		}
	}

	//for y := range bigMap {
	//	for x := range bigMap[0] {
	//		if bigMap[y][x] {
	//			fmt.Printf("#")
	//		} else {
	//			fmt.Printf(".")
	//		}
	//	}
	//	fmt.Printf("\n")
	//}
}

