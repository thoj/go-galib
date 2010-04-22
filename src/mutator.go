/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

go-galib mutators
*/

package ga

type GAMutator interface {
	// Runs mutate operation on a GAGenome
	Mutate(a GAGenome) GAGenome
	// String name of mutator
	String() string
}


//Do nothing mutator
type GANoopMutator struct{}

func (m GANoopMutator) Mutate(a GAGenome) GAGenome {
	n := a.Copy()
	return n
}
func (m GANoopMutator) String() string { return "GANoopMutator" }
