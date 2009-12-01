/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

go-galib beeders
*/


package main

import (
//	"fmt";
//	"math";
	"rand";
)

type GABreeder interface {
	// Breeds two parent GAGenomes and returns two children
	Breed(a, b GAGenome) (ca, cb GAGenome);
	// String name of breeder
	String() string;
}

type GA2PointBreeder struct {
}

func (breeder *GA2PointBreeder) Breed(a, b GAGenome) (ca, cb GAGenome) {
	if a.Len() != b.Len() {
		panic("Length mismatch in pmx");
	}
	p1 := rand.Intn(a.Len());
	p2 := rand.Intn(b.Len());
	if p1 > p2 {
		p1, p2 = p2, p1;
	}
	ca, cb = a.Crossover(b, p1, p2);
//	fmt.Printf("%d >> %d, A = %s, B = %s, C1 = %s, C2 = %s\n", p1, p2, a, b, ca, cb);
	return;
}

func (b *GA2PointBreeder) String() string {
	return "GA2PointBreeder";
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

