/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Ordered list genome for problems where the order of Genes matter, tSP for example.
*/

package ga

import (
	//"container/vector"
	//"container/list"
	"fmt"
	"math/rand"
	"sort"
)

type GAOrderedIntGenome struct {
	Gene     []int
	score    float64
	hasscore bool
	sfunc    func(ga *GAOrderedIntGenome) float64
}

func NewOrderedIntGenome(i []int, sfunc func(ga *GAOrderedIntGenome) float64) *GAOrderedIntGenome {
	g := new(GAOrderedIntGenome)
	g.Gene = i
	g.sfunc = sfunc
	return g
}

//Helper for Partially mapped crossover
func (a *GAOrderedIntGenome) pmxmap(v, p1, p2 int) (int, bool) {
	for i, c := range a.Gene {
		if c == v && (i < p1 || i > p2) {
			return i, true
		}
	}
	return 0, false
}

// Partially mapped crossover.
func (a *GAOrderedIntGenome) Crossover(bi GAGenome, p1, p2 int) (GAGenome, GAGenome) {
	ca := a.Copy().(*GAOrderedIntGenome)
	b := bi.(*GAOrderedIntGenome)
	cb := b.Copy().(*GAOrderedIntGenome)
	copy(ca.Gene[p1:p2+1], b.Gene[p1:p2+1])
	copy(cb.Gene[p1:p2+1], a.Gene[p1:p2+1])
	//Proto child needs fixing
	//amap := new(vector.IntVector)
	//bmap := new(vector.IntVector)
	amap := make([]int, 0)
	bmap := make([]int, 0)
	for i := p1; i <= p2; i++ {
		ma, found := ca.pmxmap(ca.Gene[i], p1, p2)
		if found {
			//amap.Push(ma)
			amap = append(amap, ma)
			//if bmap.Len() > 0 {
			if len(bmap) > 0 {
				//i1 := amap.Pop()
				//i2 := bmap.Pop()
				var i1, i2 int
				i1, amap = amap[len(amap)-1], amap[:len(amap)-1]
				i2, bmap = bmap[len(bmap)-1], bmap[:len(bmap)-1]
				ca.Gene[i1], cb.Gene[i2] = cb.Gene[i2], ca.Gene[i1]
			}
		}
		mb, found := cb.pmxmap(cb.Gene[i], p1, p2)
		if found {
			//bmap.Push(mb)
			bmap = append(bmap, mb)
			//if amap.Len() > 0 {
			if len(amap) > 0 {
				//i1 := amap.Pop()
				//i2 := bmap.Pop()
				var i1, i2 int
				i1, amap = amap[len(amap)-1], amap[:len(amap)-1]
				i2, bmap = bmap[len(bmap)-1], bmap[:len(bmap)-1]
				ca.Gene[i1], cb.Gene[i2] = cb.Gene[i2], ca.Gene[i1]
			}
		}
	}
	ca.Reset()
	cb.Reset()
	return ca, cb
}

func (a *GAOrderedIntGenome) Splice(bi GAGenome, from, to, length int) {
	b := bi.(*GAOrderedIntGenome)
	copy(a.Gene[to:length+to], b.Gene[from:length+from])
	a.Reset()
}

func (g *GAOrderedIntGenome) Valid() bool {
	t := g.Copy().(*GAOrderedIntGenome)
	sort.Ints(t.Gene)
	last := -9
	for _, c := range t.Gene {
		if last > -1 && c == last {
			fmt.Printf("%d - %d", c, last)
			return false
		}
		last = c
	}
	return true
}

func (g *GAOrderedIntGenome) Switch(x, y int) {
	g.Gene[x], g.Gene[y] = g.Gene[y], g.Gene[x]
	g.Reset()
}

func (g *GAOrderedIntGenome) Randomize() {
	l := len(g.Gene)
	for i := 0; i < l; i++ {
		x := rand.Intn(l)
		y := rand.Intn(l)
		g.Gene[x], g.Gene[y] = g.Gene[y], g.Gene[x]
	}
	g.Reset()
}

func (g *GAOrderedIntGenome) Copy() GAGenome {
	n := new(GAOrderedIntGenome)
	n.Gene = make([]int, len(g.Gene))
	copy(n.Gene, g.Gene)
	n.sfunc = g.sfunc
	n.score = g.score
	n.hasscore = g.hasscore
	return n
}

func (g *GAOrderedIntGenome) Len() int { return len(g.Gene) }

func (g *GAOrderedIntGenome) Score() float64 {
	if !g.hasscore {
		g.score = g.sfunc(g)
		g.hasscore = true
	}
	return g.score
}

func (g *GAOrderedIntGenome) Reset() { g.hasscore = false }

func (g *GAOrderedIntGenome) String() string { return fmt.Sprintf("%v", g.Gene) }
