/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Shifts the whole genome random length to the right.
*/

package ga

import (
	"math/rand"
)

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
