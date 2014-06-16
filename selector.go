/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

go-galib selectors
*/

package ga

import (
	"math"
	"math/rand"
	"sort"
)

type GASelector interface {
	// Select one from pop
	SelectOne(pop GAGenomes) GAGenome

	// String name of selector
	String() string
}

//This selector first selects selector.Contestants random GAGenomes
//from the population then selects one based on PElite chance.
//The best contestant has PElite chance of getting selected.
//The next best contestant has PElite^2 chance of getting selected and so on
type GATournamentSelector struct {
	PElite      float64
	Contestants int
}

func NewGATournamentSelector(pelite float64, contestants int) *GATournamentSelector {
	if pelite == 0 {
		return nil
	}
	return &GATournamentSelector{pelite, contestants}
}

func (s *GATournamentSelector) SelectOne(pop GAGenomes) GAGenome {
	if s.Contestants < 2 || s.PElite == 0 {
		panic("Contestants and PElite are not set")
	}
	g := make(GAGenomes, s.Contestants)
	l := len(pop)
	//fmt.Printf("Length = %d, Contestants = %d\n", l, len(g));
	for i := 0; i < s.Contestants; i++ {
		g[i] = pop[rand.Intn(l)]
	}
	sort.Sort(g)
	//fmt.Printf("%+v\n", g);
	r := rand.Float64()
	for i := 0; i < s.Contestants-1; i++ {
		if s.PElite*math.Pow((float64(1)-s.PElite), float64(i+1)) < r {
			return g[i]
		}
	}
	return g[s.Contestants-1]
}
func (s *GATournamentSelector) String() string {
	return "GATournamentSelector"
}
