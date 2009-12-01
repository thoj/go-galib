/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

go-galib gene
*/


package main

import (
	"fmt";
	"sort";
)

type GA struct {
	pop GAGenomes;
	
	initializer GAInitializer;
	selector GASelector;
	mutator GAMutator;
	breeder GABreeder;
}

func NewGA(i GAInitializer, s GASelector, m GAMutator, b GABreeder) *GA {
	ga := new(GA);
	ga.initializer = i;
	ga.selector = s;
	ga.mutator = m;
	ga.breeder = b;
	return ga;
}

func (ga *GA) String() string {
	return fmt.Sprintf("Initializer = %s, Selector = %s, Mutator = %s Breeder = %s", ga.initializer, ga.selector, ga.mutator, ga.breeder);
}

func (ga *GA) Init(popsize int, i GAGenome) {
	ga.pop = ga.initializer.InitPop(i, popsize);
}

func (ga *GA) Optimize(gen int) {
	for i := 0; i < gen; i++ {
		//Breed two inviduals selected with selector.
		children := make(GAGenomes, 2);
		children[0], children[1] = ga.breeder.Breed(ga.selector.SelectOne(ga.pop), ga.selector.SelectOne(ga.pop));
		ga.pop = AppendGenomes(ga.pop, children);
	}
}
func (ga *GA) PrintTop(n int) {
	fmt.Printf("Top %d Induviduals\n", n);
	sort.Sort(ga.pop);
	if len(ga.pop) < n {
	for i := 0; i < len(ga.pop); i++ {
		fmt.Printf("%2d: %s Score = %d\n", i, ga.pop[i], ga.pop[i].Score());
	}
		return;
 	}
	for i := 0; i < n; i++ {
		fmt.Printf("%2d: %s Score = %d\n", i, ga.pop[i], ga.pop[i].Score());
	}
}

func (ga *GA) PrintPop() {
	fmt.Printf("Current Population:\n");
	for i := 0; i < len(ga.pop); i++ {
		fmt.Printf("%2d: %s Score = %d\n", i, ga.pop[i], ga.pop[i].Score());
	}
}
