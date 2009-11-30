/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

go-galib selectors
*/


package main

type GASelector interface {
	// Select n from pop
	Select(pop []GAGenome, n int) ([]GAGenome);
	
	// String name of selector
	String() string;
}

type GARoulettSelector struct {
	scorefunc func(GAGenome);
}

func NewRouletteSelector(f func(GAGenome)) (s *GARoulettSelector){
	s = new(GARoulettSelector);
	s.scorefunc = f;
	return s;
}
func (s *GARoulettSelector) Select(pop []GAGenome, n int) (g []GAGenome) {
	//TODO: do something useful here
	g = make([]GAGenome, 10);
	return;
}
func (s *GARoulettSelector) String() string {
	return "GARoulettSelector";
}
