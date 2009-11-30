/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

go-galib mutators
*/

package main

type GAMutator interface {
	// Runs mutate operation on a GAGenome
	Mutate(a *GAGenome) *GAGenome;
	// String name of mutator
	String() string;
}

//This mutator switchs copies the genome and switches two genes in
//the copy and returns the new mutated copy.
type GASwitchMutator struct {
}

func (m GASwitchMutator) Mutate(a *GAGenome) *GAGenome {
	return new(GAGenome);
}
func (m GASwitchMutator) String() string {
	return "GASwitchMutator";
}
