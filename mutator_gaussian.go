/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Shifts the whole genome random length to the right.
*/

package ga

import (
	"math/rand"
)

type GAGaussianMutator struct {
	StdDev float64
	Mean   float64
}

func NewGAGaussianMutator(stddev float64, mean float64) *GAGaussianMutator {
	if stddev == 0 {
		return nil
	}
	return &GAGaussianMutator{StdDev: stddev, Mean: mean}
}

func (m GAGaussianMutator) Mutate(a GAGenome) GAGenome {
	n := a.Copy().(*GAFloatGenome)
	l := a.Len()
	s := rand.Intn(l)
	n.Gene[s] += rand.NormFloat64()*m.StdDev + m.Mean
	return n
}
func (m GAGaussianMutator) String() string { return "GAGaussianMutator" }
