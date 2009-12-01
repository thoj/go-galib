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
	//Calculate score
	Score() int;
	//Reset cached score
	Reset();

	String() string;
}

type GAGenomes []GAGenome;

func (g GAGenomes) Len() int {
	return len(g);
}
func (g GAGenomes) Less(i, j int) bool {
	return g[i].Score() < g[j].Score();
}
func (g GAGenomes) Swap(i, j int) {
	g[i], g[j] = g[j], g[i];
}

//Ordered list genome for problems where the order of Genes matter
type GAOrderedIntGenome struct {
	gene	[]int;
	score	int;
	hasscore bool;
	sfunc	func(ga *GAOrderedIntGenome) int;
}

func NewOrderedIntGenome(i []int, sfunc func(ga *GAOrderedIntGenome) int) *GAOrderedIntGenome {
	g := new(GAOrderedIntGenome);
	g.gene = i;
	g.sfunc = sfunc;
	return g;
}
func (g *GAOrderedIntGenome) Shuffle() {
	l := len(g.gene)-1;
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
		n.gene[i] = c
	}
	n.sfunc = g.sfunc;
	return n;
}

func (g *GAOrderedIntGenome) Len() int {
	return len(g.gene);
}

func (g *GAOrderedIntGenome) Score() int {
	if ! g.hasscore {
		g.score = g.sfunc(g);
		g.hasscore = true;
	}
	return int(g.score);
}

func (g *GAOrderedIntGenome) Reset()  {
	g.hasscore = false;
}


func (g *GAOrderedIntGenome) String() string	{ return fmt.Sprintf("%v", g.gene) }
