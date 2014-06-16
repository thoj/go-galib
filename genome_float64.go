/*
Copyright 2010 Thomas Jager <mail@jager.no> All rights reserved.

Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Floating point genome. For solving functions for example.
*/

package ga

import (
	"fmt"
	"math/rand"
)

type GAFloatGenome struct {
	Gene     []float64
	score    float64
	Max      float64
	Min      float64
	hasscore bool
	sfunc    func(ga *GAFloatGenome) float64
}

func NewFloatGenome(i []float64, sfunc func(ga *GAFloatGenome) float64, max float64, min float64) *GAFloatGenome {
	g := new(GAFloatGenome)
	g.Gene = i
	g.sfunc = sfunc
	g.Max = max
	g.Min = min
	return g
}

// Partially mapped crossover.
func (a *GAFloatGenome) Crossover(bi GAGenome, p1, p2 int) (GAGenome, GAGenome) {
	ca := a.Copy().(*GAFloatGenome)
	b := bi.(*GAFloatGenome)
	cb := b.Copy().(*GAFloatGenome)
	copy(ca.Gene[p1:p2+1], b.Gene[p1:p2+1])
	copy(cb.Gene[p1:p2+1], a.Gene[p1:p2+1])
	ca.Reset()
	cb.Reset()
	return ca, cb
}

func (a *GAFloatGenome) Splice(bi GAGenome, from, to, length int) {
	b := bi.(*GAFloatGenome)
	copy(a.Gene[to:length+to], b.Gene[from:length+from])
	a.Reset()
}

func (g *GAFloatGenome) Valid() bool {
	//TODO: Make this
	return true
}

func (g *GAFloatGenome) Switch(x, y int) {
	g.Gene[x], g.Gene[y] = g.Gene[y], g.Gene[x]
	g.Reset()
}

func (g *GAFloatGenome) Randomize() {
	l := len(g.Gene)
	for i := 0; i < l; i++ {
		g.Gene[i] = rand.Float64()*g.Max + g.Min
	}
	g.Reset()
}

func (g *GAFloatGenome) Copy() GAGenome {
	n := new(GAFloatGenome)
	n.Gene = make([]float64, len(g.Gene))
	copy(n.Gene, g.Gene)
	n.sfunc = g.sfunc
	n.score = g.score
	n.Max = g.Max
	n.Min = g.Min
	n.hasscore = g.hasscore
	return n
}

func (g *GAFloatGenome) Len() int { return len(g.Gene) }

func (g *GAFloatGenome) Score() float64 {
	if !g.hasscore {
		g.score = g.sfunc(g)
		g.hasscore = true
	}
	return g.score
}

func (g *GAFloatGenome) Reset() { g.hasscore = false }

func (g *GAFloatGenome) String() string { return fmt.Sprintf("%v", g.Gene) }
