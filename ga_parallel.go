/*
Copyright 2010 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Crude Parallel Genetic Algorithm
*/

package ga

import (
	"fmt"
)

type GAParallel struct {
	ga        []*GA
	Parameter GAParameter
	numproc   int
}

func NewGAParallel(parameter GAParameter, numproc int) *GAParallel {
	gap := new(GAParallel)
	gap.Parameter = parameter
	gap.ga = make([]*GA, numproc)
	gap.numproc = numproc
	for i := 0; i < numproc; i++ {
		gap.ga[i] = NewGA(parameter)
	}
	return gap
}

func (ga *GAParallel) String() string {
	return fmt.Sprintf("Initializer = %s, Selector = %s, Mutator = %s Breeder = %s",
		ga.Parameter.Initializer,
		ga.Parameter.Selector,
		ga.Parameter.Mutator,
		ga.Parameter.Breeder)
}

func (ga *GAParallel) Init(popsize int, init GAGenome) {
	for i := 0; i < ga.numproc; i++ {
		ga.ga[i].Init(popsize, init)
	}
}

func optimize_worker(ga *GA, gen int, c chan int) {
	ga.Optimize(gen)
	c <- 1
}

func (ga *GAParallel) Optimize(gen int) {
	c := make(chan int, ga.numproc)
	for i := 0; i < ga.numproc; i++ {
		go optimize_worker(ga.ga[i], gen, c)
	}
	for i := 0; i < ga.numproc; i++ {
		<-c
	}
	nselect := gen * 2
	children := make([]GAGenomes, ga.numproc)
	for i := 0; i < ga.numproc; i++ {
		children[i] = make(GAGenomes, nselect)
		for j := 0; j < nselect; j++ {
			children[i][j] = ga.ga[i].Parameter.Selector.SelectOne(ga.ga[i].pop)
		}
	}
	j := ga.numproc - 1
	for i := 0; i < ga.numproc; i++ {
		ga.ga[i].pop = AppendGenomes(ga.ga[i].pop, children[j])
		j--
	}
}


func (ga *GAParallel) OptimizeUntil(stop func(best GAGenome) bool) {
	for !stop(ga.Best()) {
		ga.Optimize(1)
	}
}


func (ga *GAParallel) Best() GAGenome {
	best := ga.ga[0].Best()
	for i := 1; i < ga.numproc; i++ {
		nbest := ga.ga[i].Best()
		if nbest.Score() < best.Score() {
			best = nbest
		}
	}
	return best
}
