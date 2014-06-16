/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Fixed length Bitstring genome for problems like subset sum
*/

package ga

import (
	"fmt"
	"math/rand"
)

type GAFixedBitstringGenome struct {
	Gene     []bool
	score    float64
	hasscore bool
	sfunc    func(ga *GAFixedBitstringGenome) float64
}

func NewFixedBitstringGenome(i []bool, sfunc func(ga *GAFixedBitstringGenome) float64) *GAFixedBitstringGenome {
	g := new(GAFixedBitstringGenome)
	g.Gene = i
	g.sfunc = sfunc
	g.Reset()
	return g
}

//Simple 2 point crossover
func (a *GAFixedBitstringGenome) Crossover(bi GAGenome, p1, p2 int) (GAGenome, GAGenome) {
	ca := a.Copy().(*GAFixedBitstringGenome)
	b := bi.(*GAFixedBitstringGenome)
	cb := b.Copy().(*GAFixedBitstringGenome)
	copy(ca.Gene[p1:p2+1], b.Gene[p1:p2+1])
	copy(cb.Gene[p1:p2+1], a.Gene[p1:p2+1])
	ca.Reset()
	cb.Reset()
	return ca, cb
}

func (a *GAFixedBitstringGenome) Splice(bi GAGenome, from, to, length int) {
	b := bi.(*GAFixedBitstringGenome)
	copy(a.Gene[to:length+to], b.Gene[from:length+from])
	a.Reset()
}

func (g *GAFixedBitstringGenome) Valid() bool { return true }

func (g *GAFixedBitstringGenome) Switch(x, y int) {
	g.Gene[x], g.Gene[y] = g.Gene[y], g.Gene[x]
	g.Reset()
}

func (g *GAFixedBitstringGenome) Randomize() {
	l := len(g.Gene)
	for i := 0; i < l; i++ {
		x := rand.Intn(2)
		if x == 1 {
			g.Gene[i] = true
		} else {
			g.Gene[i] = false
		}
	}
	g.Reset()
}

func (g *GAFixedBitstringGenome) Copy() GAGenome {
	n := new(GAFixedBitstringGenome)
	n.Gene = make([]bool, len(g.Gene))
	copy(n.Gene, g.Gene)
	n.sfunc = g.sfunc
	n.score = g.score
	n.hasscore = g.hasscore
	return n
}

func (g *GAFixedBitstringGenome) Len() int { return len(g.Gene) }

func (g *GAFixedBitstringGenome) Score() float64 {
	if !g.hasscore {
		g.score = g.sfunc(g)
		g.hasscore = true
	}
	return g.score
}

func (g *GAFixedBitstringGenome) Reset() { g.hasscore = false }

func (g *GAFixedBitstringGenome) String() string {
	return fmt.Sprintf("%v", g.Gene)
}
