/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

go-galib gene
*/


package main

type GA struct {
	pop []GAGenome;
	
	initializer *GAInitializer;
	selector *GASelector;
	mutator *GAMutator;
	breeder *GABreeder;
}

func NewGA(i *GAInitializer, s *GASelector, m *GAMutator, b *GABreeder) {
	ga := new(GA);
	ga.initializer = i;
	ga.selector = s;
	ga.mutator = m;
	ga.breeder = b;
}
