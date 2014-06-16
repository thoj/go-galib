/*
Copyright 2010 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Genetic Algorithm
*/

package ga

import (
	"fmt"
	"math/rand"
	"sort"
)

type GAParameter struct {
	//Chance of breeding
	PBreed float64
	//Chance of mutation
	PMutate float64

	// Initializer, Selector, Mutator, Breeder Objects this GA will use
	Initializer GAInitializer
	Selector    GASelector
	Mutator     GAMutator
	Breeder     GABreeder
}

type GA struct {
	pop     GAGenomes
	popsize int

	Parameter GAParameter
}

func NewGA(parameter GAParameter) *GA {
	ga := new(GA)
	ga.Parameter = parameter
	return ga
}

func (ga *GA) String() string {
	return fmt.Sprintf("Initializer = %s, Selector = %s, Mutator = %s Breeder = %s",
		ga.Parameter.Initializer,
		ga.Parameter.Selector,
		ga.Parameter.Mutator,
		ga.Parameter.Breeder)
}

func (ga *GA) Init(popsize int, i GAGenome) {
	ga.pop = ga.Parameter.Initializer.InitPop(i, popsize)
	ga.popsize = popsize
}

func (ga *GA) Optimize(gen int) {
	for i := 0; i < gen; i++ {
		l := len(ga.pop) // Do not try to breed/mutate new in this gen
		for p := 0; p < l; p++ {
			//Breed two inviduals selected with selector.
			if ga.Parameter.PBreed > rand.Float64() {
				children := make(GAGenomes, 2)
				children[0], children[1] = ga.Parameter.Breeder.Breed(
					ga.Parameter.Selector.SelectOne(ga.pop),
					ga.Parameter.Selector.SelectOne(ga.pop))

				ga.pop = AppendGenomes(ga.pop, children)
			}
			//Mutate
			if ga.Parameter.PMutate > rand.Float64() {
				children := make(GAGenomes, 1)
				children[0] = ga.Parameter.Mutator.Mutate(ga.pop[p])
				ga.pop = AppendGenomes(ga.pop, children)
			}
		}
		//cleanup remove some from pop
		// this should probably use a type of selector
		sort.Sort(ga.pop)
		ga.pop = ga.pop[0:ga.popsize]
	}
}

func (ga *GA) OptimizeUntil(stop func(best GAGenome) bool) {
	for !stop(ga.Best()) {
		ga.Optimize(1)
	}
}

func (ga *GA) Best() GAGenome {
	sort.Sort(ga.pop)
	return ga.pop[0]
}

func (ga *GA) PrintTop(n int) {
	sort.Sort(ga.pop)
	if len(ga.pop) < n {
		for i := 0; i < len(ga.pop); i++ {
			fmt.Printf("%2d: %s Score = %f\n", i, ga.pop[i], ga.pop[i].Score())
		}
		return
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%2d: %s Score = %f\n", i, ga.pop[i], ga.pop[i].Score())
	}
}

func (ga *GA) PrintPop() {
	fmt.Printf("Current Population:\n")
	for i := 0; i < len(ga.pop); i++ {
		fmt.Printf("%2d: %s Score = %f\n", i, ga.pop[i], ga.pop[i].Score())
	}
}
