/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Combines several mutators into one, each mutation has equal chance of occuring.
*/

package ga

import (
	"rand"
	"container/vector"
	"fmt"
)

type GAMultiMutator struct {
	v     *vector.Vector
	stats []int
}

func NewMultiMutator() *GAMultiMutator {
	m := new(GAMultiMutator)
	m.v = new(vector.Vector)
	m.stats = make([]int, 100)
	return m
}

func (m GAMultiMutator) Mutate(a GAGenome) GAGenome {
	if m.v.Len() == 0 {
		panic("No mutators added!")
	}
	r := float64(1.0 / float64(m.v.Len()))
	for i := 0; i < m.v.Len()-1; i++ {
		if rand.Float64() < r {
			sm := m.v.At(i).(GAMutator)
			m.stats[i]++
			return sm.Mutate(a)
		}
	}
	sm := m.v.At(m.v.Len() - 1).(GAMutator)
	m.stats[m.v.Len()-1]++
	return sm.Mutate(a)
}

//Add mutator
func (m *GAMultiMutator) Add(a GAMutator) { m.v.Push(a) }
func (m GAMultiMutator) String() string   { return "GAMultiMutator" }
func (m *GAMultiMutator) Stats() string {
	o := "Used "
	for i := 0; i < m.v.Len(); i++ {
		sm := m.v.At(i).(GAMutator)
		o = fmt.Sprintf("%s%s %d times, ", o, sm, m.stats[i])
	}
	return o
}
