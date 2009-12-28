/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

go-galib initializers
*/


package ga

type GAInitializer interface {
	// Initializes popsize length []GAGenome from i
	InitPop(i GAGenome, popsize int) []GAGenome
	// String name of initializers
	String() string
}

type GARandomInitializer struct{}

func (i *GARandomInitializer) InitPop(first GAGenome, popsize int) (pop []GAGenome) {
	pop = make([]GAGenome, popsize)
	for x := 0; x < popsize; x++ {
		pop[x] = first.Copy()
		pop[x].Randomize()
	}
	return pop
}

func (i *GARandomInitializer) String() string { return "RandomInitializer" }
