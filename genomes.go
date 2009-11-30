/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

go-galib genome
*/

package main

import (
	"rand";
	"fmt";
)

type GAGenome interface {
	//Shuffle genomes araound
	Shuffle();
	//Copy a genome;
	Copy() GAGenome;

	String() string;
}

//Ordered list genome for problems where the order of Genes matter
type GAOrderedIntGenome struct {
	gene	[]int;
}

func NewOrderedIntGenome(i []int) *GAOrderedIntGenome {
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

func (g *GAOrderedIntGenome) Copy() GAGenome {
	n := new(GAOrderedIntGenome);
	n.gene = make([]int, len(g.gene));
	for i, c := range g.gene {
		n.gene[i] = c; 
	}
	return n;
}

func (g *GAOrderedIntGenome) String() string	{ return fmt.Sprintf("%v", g.gene) }
