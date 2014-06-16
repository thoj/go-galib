/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Example of uing the ordered int genome and mutators
*/
package main

import (
	"fmt"
	"github.com/thoj/go-galib"
	"math/rand"
	"time"
)

var scores int

// Boring fitness/score function.
func score(g *ga.GAOrderedIntGenome) float64 {
	var total int
	for i, c := range g.Gene {
		total += c ^ i
	}
	scores++
	return float64(total)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	m := ga.NewMultiMutator()
	msh := new(ga.GAShiftMutator)
	msw := new(ga.GASwitchMutator)
	m.Add(msh)
	m.Add(msw)

	param := ga.GAParameter{
		Initializer: new(ga.GARandomInitializer),
		Selector:    ga.NewGATournamentSelector(0.7, 5),
		Breeder:     new(ga.GA2PointBreeder),
		Mutator:     m,
		PMutate:     0.1,
		PBreed:      0.7}

	gao := ga.NewGA(param)

	genome := ga.NewOrderedIntGenome([]int{10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, score)

	gao.Init(5, genome) //Total population

	gao.Optimize(10) // Run genetic algorithm for 20 generations.
	gao.PrintTop(10)

	fmt.Printf("Calls to score = %d\n", scores)
	fmt.Printf("%s\n", m.Stats())
}
