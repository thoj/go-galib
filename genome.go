/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

go-galib genome
*/

package ga

// Genome interface, Not final.
type GAGenome interface {
	//Randomize.Genens
	Randomize()
	//Copy a genome;
	Copy() GAGenome
	//Calculate score
	Score() float64
	//Reset cached score
	Reset()
	//Crossover for this genome
	Crossover(bi GAGenome, p1, p2 int) (ca GAGenome, cb GAGenome)
	//Switch for the genome
	Switch(x, y int)
	//Splice two genomes;
	Splice(bi GAGenome, from, to, length int)
	//Check if genome is valid
	Valid() bool

	String() string
	Len() int
}

type GAGenomes []GAGenome

func (g GAGenomes) Len() int           { return len(g) }
func (g GAGenomes) Less(i, j int) bool { return g[i].Score() < g[j].Score() }
func (g GAGenomes) Swap(i, j int)      { g[i], g[j] = g[j], g[i] }


func AppendGenomes(slice, data GAGenomes) GAGenomes {
	l := len(slice)
	if l+len(data) > cap(slice) {
		newSlice := make(GAGenomes, (l+len(data))*2)
		for i, c := range slice {
			newSlice[i] = c
		}
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	for i, c := range data {
		slice[l+i] = c
	}
	return slice
}
