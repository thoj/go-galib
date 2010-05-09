/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

go-galib gene
*/

package ga

import (
	"fmt"
	"sort"
	"rand"
)

type GA struct {
	pop GAGenomes

	initializer GAInitializer
	selector    GASelector
	mutator     GAMutator
	breeder     GABreeder

	PMutate float64
	PBreed  float64

	popsize int
}

func NewGA(i GAInitializer, s GASelector, m GAMutator, b GABreeder) *GA {
	ga := new(GA)
	ga.initializer = i
	ga.selector = s
	ga.mutator = m
	ga.breeder = b
	ga.PMutate = 0.05
	ga.PBreed = 0.1
	return ga
}

func (ga *GA) String() string {
	return fmt.Sprintf("Initializer = %s, Selector = %s, Mutator = %s Breeder = %s", ga.initializer, ga.selector, ga.mutator, ga.breeder)
}

func (ga *GA) Init(popsize int, i GAGenome) {
	ga.pop = ga.initializer.InitPop(i, popsize)
	ga.popsize = popsize
}

func (ga *GA) Optimize(gen int) {
	for i := 0; i < gen; i++ {
		l := len(ga.pop) // Do not try to breed/mutate new in this gen
		for p := 0; p < l; p++ {
			//Breed two inviduals selected with selector.
			if ga.PBreed > rand.Float64() {
				children := make(GAGenomes, 2)
				children[0], children[1] = ga.breeder.Breed(ga.selector.SelectOne(ga.pop), ga.selector.SelectOne(ga.pop))
				ga.pop = AppendGenomes(ga.pop, children)
			}
			//Mutate
			if ga.PMutate > rand.Float64() {
				children := make(GAGenomes, 1)
				children[0] = ga.mutator.Mutate(ga.pop[p])
				ga.pop = AppendGenomes(ga.pop, children)
			}
		}
		//cleanup remove some from pop
		// this should probably use a type of selector
		sort.Sort(ga.pop)
		ga.pop = ga.pop[0:ga.popsize]
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
