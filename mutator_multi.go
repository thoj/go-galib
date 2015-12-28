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
	"strings"
)

type GAMultiMutator struct {
	v     []GAMutator
	stats []int
}

// NewMultiMutator returns a new, empty, multi mutator.
func NewMultiMutator() *GAMultiMutator {
	return &GAMultiMutator{
		v:     make([]GAMutator, 0),
		stats: make([]int, 0),
	}
}

// Mutate mutates the genome using one of the mutators added using Add(). Each
// mutator has equal chance of being chosen.
func (m *GAMultiMutator) Mutate(a GAGenome) GAGenome {
	if len(m.v) == 0 {
		// No mutators, so nothing to do.
		return a.Copy()
	}
	r := rand.Intn(len(m.v))
	m.stats[r]++
	return m.v[r].Mutate(a)
}

// Add adds a mutator to the MultiMutator.
func (m *GAMultiMutator) Add(a GAMutator) {
	m.v = append(m.v, a)
	m.stats = append(m.stats, 0)
}

// String returns the name of the mutator.
func (m GAMultiMutator) String() string { return "GAMultiMutator" }

// Stats() returns a strings with usage details of the individual mutators.
func (m *GAMultiMutator) Stats() string {
	var o []string
	for i, sm := range m.v {
		o = append(o, fmt.Sprintf("%s %d times", sm, m.stats[i]))
	}
	return "Used " + strings.Join(o, ", ")
}
