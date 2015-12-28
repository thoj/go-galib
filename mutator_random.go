/*
Copyright 2015 Ludovico Gardenghi <lu@dovi.co>.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

This mutator replaces one gene with a random one. It works by creating a new random genome and
splicing one gene at a random position on the given one, so it is not very efficient (in many cases
it would suffice to generate one single gene rather than a full genome).
*/

package ga

import (
	"math/rand"
)

type GAMutatorRandom struct{}

// Mutate returns a genome which is identical to the given one except for one
// gene, which is replaced with a random one.
func (m GAMutatorRandom) Mutate(a GAGenome) GAGenome {
	r := a.Copy()
	r.Randomize()
	p := rand.Intn(a.Len())

	ac := a.Copy()
	ac.Splice(r, p, p, 1)
	return ac
}
func (m GAMutatorRandom) String() string { return "GAMutatorRandom" }
