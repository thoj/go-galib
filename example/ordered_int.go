/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Example of uing the ordered int genome and mutators
*/
package main

import (
	"fmt"
	"rand"
	"time"
	"../_obj/ga"
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
	rand.Seed(time.Nanoseconds())
	//m := new(ga.GASwitchMutator);
	m := ga.NewMultiMutator()
	msh := new(ga.GAShiftMutator)
	msw := new(ga.GASwitchMutator)
	m.Add(msh)
	m.Add(msw)
	b := new(ga.GA2PointBreeder)

	s := new(ga.GATournamentSelector)
	s.Contestants = 5 //Contestants in Tournament
	s.PElite = 0.7    //Chance of best contestant winning, chance of next is PElite^2 and so on.

	i := new(ga.GARandomInitializer)
	gao := ga.NewGA(i, s, m, b)
	gao.PMutate = 0.2 //Chance of mutation
	gao.PBreed = 0.2  //Chance of breeding

	fmt.Printf("%s\n", gao)
	genome := ga.NewOrderedIntGenome([]int{10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, score)
	gao.Init(200, genome) //Total population

	gao.Optimize(20) // Run genetic algorithm for 20 generations.
	gao.PrintTop(10)
	fmt.Printf("Calls to score = %d\n", scores);
	fmt.Printf("%s\n", m.Stats())
}
