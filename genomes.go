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
	"container/vector";
)

type GAGenome interface {
	//Shuffle genens
	Shuffle();
	//Copy a genome;
	Copy() GAGenome;
	//Calculate score
	Score() int;
	//Reset cached score
	Reset();
	//Crossover for this genome
	Crossover(bi GAGenome, p1, p2 int) (ca GAGenome, cb GAGenome);
	//Switch for the genome
	Switch(x, y int);

	String() string;
	Len() int;
}

type GAGenomes []GAGenome

func (g GAGenomes) Len() int		{ return len(g) }
func (g GAGenomes) Less(i, j int) bool	{ return g[i].Score() < g[j].Score() }
func (g GAGenomes) Swap(i, j int)	{ g[i], g[j] = g[j], g[i] }

func AppendGenomes(slice, data GAGenomes) GAGenomes {
	l := len(slice);
	if l+len(data) > cap(slice) {
		newSlice := make(GAGenomes, (l+len(data))*2);
		for i, c := range slice {
			newSlice[i] = c
		}
		slice = newSlice;
	}
	slice = slice[0 : l+len(data)];
	for i, c := range data {
		slice[l+i] = c
	}
	return slice;
}


//Ordered list genome for problems where the order of Genes matter
type GAOrderedIntGenome struct {
	gene		[]int;
	score		int;
	hasscore	bool;
	sfunc		func(ga GAOrderedIntGenome) int;
}

func NewOrderedIntGenome(i []int, sfunc func(ga GAOrderedIntGenome) int) *GAOrderedIntGenome {
	g := new(GAOrderedIntGenome);
	g.gene = i;
	g.sfunc = sfunc;
	return g;
}
//Helper for Partially mapped crossover
func (a GAOrderedIntGenome) pmxmap(v, p1, p2 int) (int, bool) {
	for i, c := range a.gene {
		if c == v && ( i < p1 || i > p2) { 
			return i, true
		}
	}
	return 0, false;
}

// Partially mapped crossover.
func (a GAOrderedIntGenome) Crossover(bi GAGenome, p1, p2 int) (GAGenome, GAGenome) {
	ca := a.Copy().(*GAOrderedIntGenome);
	b := bi.(*GAOrderedIntGenome);
	cb := b.Copy().(*GAOrderedIntGenome);
	copy(ca.gene[p1:p2+1], b.gene[p1:p2+1]);
	copy(cb.gene[p1:p2+1], a.gene[p1:p2+1]);
	//Proto child needs fixing
	amap := new(vector.IntVector);
	bmap := new(vector.IntVector);
	for i := p1; i <= p2; i++ {
		ma, found := ca.pmxmap(ca.gene[i], p1, p2);
		if found {
			amap.Push(ma);
			if bmap.Len() > 0 {
				i1 := amap.Pop();
				i2 := bmap.Pop();
				ca.gene[i1], cb.gene[i2] = cb.gene[i2], ca.gene[i1];
			}
		}
		mb, found := cb.pmxmap(cb.gene[i], p1, p2);
		if found {
			bmap.Push(mb);
			if amap.Len() > 0 {
				i1 := amap.Pop();
				i2 := bmap.Pop();
				ca.gene[i1], cb.gene[i2] = cb.gene[i2], ca.gene[i1];
			}
		}
	}
	return ca, cb;
}
/*
func (g GAOrderedIntGenome) Valid() bool {
	t := g.Copy().(*GAOrderedIntGenome);
	sort.SortInts(t.gene);
	last := -9;
	for _, c := range t.gene {
		if last > -1 && c == last {
			fmt.Printf("%d - %d", c, last);
			return false;
		}
		last = c;
	}
	return true;
}
*/

func (g GAOrderedIntGenome) Switch(x, y int) {
	g.gene[x], g.gene[y] = g.gene[y], g.gene[x];
}

func (g GAOrderedIntGenome) Shuffle() {
	l := len(g.gene);
	for i := 0; i < l; i++ {
		x := rand.Intn(l);
		y := rand.Intn(l);
		g.gene[x], g.gene[y] = g.gene[y], g.gene[x];
	}
}

func (g GAOrderedIntGenome) Copy() GAGenome {
	n := new(GAOrderedIntGenome);
	n.gene = make([]int, len(g.gene));
	for i, c := range g.gene {
		n.gene[i] = c
	}
	n.sfunc = g.sfunc;
	return n;
}

func (g GAOrderedIntGenome) Len() int	{ return len(g.gene) }

func (g GAOrderedIntGenome) Score() int {
	if !g.hasscore {
		g.score = g.sfunc(g);
		g.hasscore = true;
	}
	return int(g.score);
}

func (g GAOrderedIntGenome) Reset()	{ g.hasscore = false }


func (g GAOrderedIntGenome) String() string	{ return fmt.Sprintf("%v", g.gene) }
