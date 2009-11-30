/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

go-galib initializers
*/


package main

type GAInitializer interface {
	// Initializes popsize length []GAGenome from i
	InitPop(i GAGenome, popsize int) ([]GAGenome);
	// String name of initializers
	String() string;
}

type GAShuffleInitializer struct{
}

func (i *GAShuffleInitializer) InitPop(first GAGenome, popsize int) (pop []GAGenome){
	pop = make([]GAGenome, popsize);
	for x := 0; x < popsize; x ++ {
		pop[x] = first.Copy();
		pop[x].Shuffle();
	}
	return pop;
}

func (i *GAShuffleInitializer) String() string {
	return "ShuffleInitializer";
}
