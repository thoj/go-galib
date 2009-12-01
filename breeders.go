/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

go-galib beeders
*/


package main

type GABreeder interface {
	// Breeds two parent GAGenomes and returns two children
	Breed(a, b GAGenome) (ca, cb GAGenome);
	// String name of breeder
	String() string;
}

type GARandomBreeder struct {
}

func (breeder *GARandomBreeder) Breed(a, b GAGenome) (ca, cb GAGenome) {
	ca = a.Copy();
	ca.Shuffle();
	cb = b.Copy();
	cb.Shuffle();
	return;
}

func (b *GARandomBreeder) String() string {
	return "GARandomBreeder";
}

