/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

go-galib genome
*/

package main

import (
	"rand";
)

type GAGenome interface {
	//Shuffle genomes araound
	Shuffle();
	//Copy a genome;
	Copy() GAGenome;
}

//Ordered list genome for problems where the order of Genes matter
type GAOrderedIntGenome struct {
	gene []int;
}

func NewGAOrderedIntGenome(i []int) *GAOrderedIntGenome {
	g := new(GAOrderedIntGenome);
	g.gene = i;
	return g;
}
func (g *GAOrderedIntGenome) Shuffle() {
	l := len(g.gene);
	for i := 0; i < l; i++ {
		x := rand.Intn(l);
		y := rand.Intn(l);
		g.gene[x], g.gene[y] = g.gene[y], g.gene[x];
	}
}

func (g *GAOrderedIntGenome) Copy() *GAOrderedIntGenome {
	n := new(GAOrderedIntGenome);
        n.gene = g.gene;
	return n;
}
