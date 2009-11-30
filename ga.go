/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

go-galib gene
*/


package main

import (
	"fmt";
)

type GA struct {
	pop []GAGenome;
	
	initializer GAInitializer;
	selector GASelector;
	mutator GAMutator;
	breeder GABreeder;
}

func NewGA(i GAInitializer, s GASelector, m GAMutator, b GABreeder) *GA {
	ga := new(GA);
	ga.initializer = i;
	ga.selector = s;
	ga.mutator = m;
	ga.breeder = b;
	return ga;
}

func (ga *GA) String() string {
	return fmt.Sprintf("Initializer = %s, Selector = %s, Mutator = %s Breeder = %s", ga.initializer, ga.selector, ga.mutator, ga.breeder);
}
