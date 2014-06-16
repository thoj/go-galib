/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

This mutator copies the genome and switches two genes in
the copy and returns the new mutated copy.
*/

package ga

import (
	"math/rand"
)

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
