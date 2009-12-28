/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

go-galib mutators
*/

package ga

import (
	"rand"
	"container/vector"
	"fmt"
)

type GAMutator interface {
	// Runs mutate operation on a GAGenome
	Mutate(a GAGenome) GAGenome
	// String name of mutator
	String() string
}

//This mutator copies the genome and switches two genes in
//the copy and returns the new mutated copy.
type GASwitchMutator struct{}

func (m GASwitchMutator) Mutate(a GAGenome) GAGenome {
	n := a.Copy()
	p1 := rand.Intn(a.Len())
	p2 := rand.Intn(a.Len())
	if p1 > p2 {
		p1, p2 = p2, p1
	}
	n.Switch(p1, p2)
	return n
}
func (m GASwitchMutator) String() string { return "GASwitchMutator" }

//Shifts the whole genome random length to the right.
type GAShiftMutator struct{}

func (m GAShiftMutator) Mutate(a GAGenome) GAGenome {
	n := a.Copy()
	l := a.Len()
	s := rand.Intn(l / 2)
	n.Splice(a, l-s, 0, s)
	n.Splice(a, 0, l-s, s)
	return n
}
func (m GAShiftMutator) String() string { return "GAShiftMutator" }

//Combines several mutators into one, each mutation has equal chance of occuring.
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
