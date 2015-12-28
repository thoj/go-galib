/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Combines several mutators into one, each mutation has equal chance of occuring.
*/

package ga

import (
	"fmt"
	"math/rand"
)

type GAMultiMutator struct {
	v     []GAMutator
	stats []int
}

func NewMultiMutator() *GAMultiMutator {
	m := new(GAMultiMutator)
	m.v = make([]GAMutator, 0)
	m.stats = make([]int, 0)
	return m
}

func (m *GAMultiMutator) Mutate(a GAGenome) GAGenome {
	if len(m.v) == 0 {
		// No mutators, so nothing to do.
		return a.Copy()
	}
	r := rand.Intn(len(m.v))
	m.stats[r]++
	return m.v[r].Mutate(a)
}

//Add mutator
func (m *GAMultiMutator) Add(a GAMutator) {
	m.v = append(m.v, a)
	m.stats = append(m.stats, 0)
}
func (m GAMultiMutator) String() string { return "GAMultiMutator" }
func (m *GAMultiMutator) Stats() string {
	o := "Used "
	for i := 0; i < len(m.v); i++ {
		sm := m.v[i].(GAMutator)
		o = fmt.Sprintf("%s%s %d times, ", o, sm, m.stats[i])
	}
	return o
}
