/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

go-galib selectors
*/


package main

import (
	"math";
	"rand";
	"sort";
)

type GASelector interface {
	// Select n from pop
	SelectOne(pop GAGenomes) GAGenome;
	
	// String name of selector
	String() string;
}


type GATournamentSelector struct {
	PElite float64;
	Contestants int;
}

func (s *GATournamentSelector) SelectOne(pop GAGenomes) GAGenome {
	if s.Contestants < 2 {
		panic("Set selector.Contestants > 1");
	}
	if s.PElite == 0{
		panic("Set selector.PElite to float64 (0.5 is a good choice for most problems)");
	}
	g := make(GAGenomes, s.Contestants);
	l := len(pop);
	//fmt.Printf("Length = %d, Contestants = %d\n", l, len(g));
	for i := 0; i < s.Contestants; i++ {
		g[i] = pop[rand.Intn(l)];
	}
	sort.Sort(g);
	//fmt.Printf("%+v\n", g);
	r := rand.Float64();
	for i := 0; i < s.Contestants - 1; i++ {
		if s.PElite * math.Pow((float64(1) - s.PElite), float64(i+1)) < r {
			return g[i];
		}
	}
	return g[s.Contestants - 1];
}
func (s *GATournamentSelector) String() string {
	return "GARoulettSelector";
}
