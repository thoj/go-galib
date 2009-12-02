/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

go-galib genome
*/

package ga

import (
	"rand";
	"fmt";
	"container/vector";
)

type GAGenome interface {
	//Randomize.Genens
	Randomize();
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
	//Splice two genomes;
	Splice(bi GAGenome, from, to, length int);

	String() string;
	Len() int;
}

type GAGenomes []GAGenome

func (g GAGenomes) Len() int		{ return len(g) }
func (g GAGenomes) Less(i, j int) bool	{ return g[i].Score() < g[j].Score() }
func (g GAGenomes) Swap(i, j int)	{ g[i], g[j] = g[j], g[i] }


//Helper
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
	Gene		[]int;
	score		int;
	hasscore	bool;
	sfunc		func(ga GAOrderedIntGenome) int;
}

func NewOrderedIntGenome(i []int, sfunc func(ga GAOrderedIntGenome) int) *GAOrderedIntGenome {
	g := new(GAOrderedIntGenome);
	g.Gene = i;
	g.sfunc = sfunc;
	return g;
}
//Helper for Partially mapped crossover
func (a GAOrderedIntGenome) pmxmap(v, p1, p2 int) (int, bool) {
	for i, c := range a.Gene {
		if c == v && (i < p1 || i > p2) {
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
	copy(ca.Gene[p1:p2+1], b.Gene[p1:p2+1]);
	copy(cb.Gene[p1:p2+1], a.Gene[p1:p2+1]);
	//Proto child needs fixing
	amap := new(vector.IntVector);
	bmap := new(vector.IntVector);
	for i := p1; i <= p2; i++ {
		ma, found := ca.pmxmap(ca.Gene[i], p1, p2);
		if found {
			amap.Push(ma);
			if bmap.Len() > 0 {
				i1 := amap.Pop();
				i2 := bmap.Pop();
				ca.Gene[i1], cb.Gene[i2] = cb.Gene[i2], ca.Gene[i1];
			}
		}
		mb, found := cb.pmxmap(cb.Gene[i], p1, p2);
		if found {
			bmap.Push(mb);
			if amap.Len() > 0 {
				i1 := amap.Pop();
				i2 := bmap.Pop();
				ca.Gene[i1], cb.Gene[i2] = cb.Gene[i2], ca.Gene[i1];
			}
		}
	}
	ca.Reset();
	cb.Reset();
	return ca, cb;
}

func (a GAOrderedIntGenome) Splice(bi GAGenome, from, to, length int ) {
	b := bi.(*GAOrderedIntGenome);
	//fmt.Printf("Splice: copy(a.Gene[%d:%d], b.Gene[%d:%d])\n", to, length, from, length);
	copy(a.Gene[to:length+to], b.Gene[from:length+from]);
}
/*
func (g GAOrderedIntGenome) Valid() bool {
	t := g.Copy().(*GAOrderedIntGenome);
	sort.SortInts(t.Gene);
	last := -9;
	for _, c := range t.Gene {
		if last > -1 && c == last {
			fmt.Printf("%d - %d", c, last);
			return false;
		}
		last = c;
	}
	return true;
}
*/

func (g GAOrderedIntGenome) Switch(x, y int)	{ g.Gene[x], g.Gene[y] = g.Gene[y], g.Gene[x]; g.hasscore = false; }

func (g GAOrderedIntGenome) Randomize() {
	l := len(g.Gene);
	for i := 0; i < l; i++ {
		x := rand.Intn(l);
		y := rand.Intn(l);
		g.Gene[x], g.Gene[y] = g.Gene[y], g.Gene[x];
	}
	g.Reset();
}

func (g GAOrderedIntGenome) Copy() GAGenome {
	n := new(GAOrderedIntGenome);
	n.Gene = make([]int, len(g.Gene));
	for i, c := range g.Gene {
		n.Gene[i] = c
	}
	n.sfunc = g.sfunc;
	n.score = g.score;
	g.hasscore = true;
	return n;
}

func (g GAOrderedIntGenome) Len() int	{ return len(g.Gene) }

func (g GAOrderedIntGenome) Score() int {
	if !g.hasscore {
		g.score = g.sfunc(g);
		g.hasscore = true;
	}
	return int(g.score);
}

func (g GAOrderedIntGenome) Reset()	{ g.hasscore = false }


func (g GAOrderedIntGenome) String() string	{ return fmt.Sprintf("%v", g.Gene) }
