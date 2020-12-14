package main

import (
	"fmt"
	"github.com/ghjm/advent2020/pkg/utils"
	"math"
	"math/big"
	"strings"
	"sync"
)

// https://rosettacode.org/wiki/Chinese_remainder_theorem#Go
var one = big.NewInt(1)
func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}

func main() {
	startTime := 0
	buses := make([]string, 0)
	strChan := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		startTime = utils.AtoiOrPanic(<-strChan)
		buses = strings.Split(<-strChan, ",")
		wg.Done()
	}()
	err := utils.OpenAndReadAll("input13.txt", func(s string) error {
		strChan <- s
		return nil
	})
	if err != nil {
		panic(err)
	}
	wg.Wait()
	close(strChan)

	fmt.Printf("Part One\n")
	bestTime := math.MaxInt32
	bestBus := 0
	for _, s := range buses {
		if s == "x" {
			continue
		}
		bus := utils.AtoiOrPanic(s)
		nextTime := startTime - (startTime % bus) + bus
		if nextTime < bestTime {
			bestTime = nextTime
			bestBus = bus
		}
	}
	waitTime := bestTime - startTime
	fmt.Printf("Wait times bus: %d\n", waitTime * bestBus)

	fmt.Printf("Part Two\n")
	coef := make([]*big.Int, 0)
	mod := make([]*big.Int, 0)
	for i, s := range buses {
		if s == "x" {
			continue
		}
		bus := utils.AtoiOrPanic(s)
		bigBus := big.NewInt(int64(bus))
		if !bigBus.ProbablyPrime(0) {
			panic(fmt.Sprintf("Bus %d is not prime\n", bus))
		}
		a := (bus - i) % bus
		if a < 0 {
			a += bus
		}
		bigCoef := big.NewInt(int64(a))
		coef = append(coef, bigCoef)
		mod = append(mod, bigBus)
	}
	result, err := crt(coef, mod)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Timestamp: %d\n", result)
}