/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

subset sum solver
*/
package main

import (
	"fmt"
	"github.com/thoj/go-galib"
	"math"
	"math/rand"
	"time"
)

var scores int

func ackley(g *ga.GAFloatGenome) float64 {
	scores++
	var sum1 float64 = 0.0
	for _, c := range g.Gene {
		sum1 += float64(c * c)
	}
	t1 := math.Exp(-0.2 * (math.Sqrt((1.0 / 5.0) * sum1)))
	sum1 = 0.0
	for _, c := range g.Gene {
		sum1 += math.Cos(float64(2.0 * math.Pi * c))
	}
	t2 := math.Exp((1.0 / 5.0) * sum1)
	return (20 + math.Exp(1) - 20*t1 - t2)
}

func rosenbrock(g *ga.GAFloatGenome) float64 {
	scores++
	var sum float64
	for i := 1; i < len(g.Gene); i++ {
		sum += 100.0*math.Pow(math.Pow(g.Gene[i]-g.Gene[i-1], 2), 2) + math.Pow(1-g.Gene[i-1], 2)
	}
	return sum
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	param := ga.GAParameter{
		Initializer: new(ga.GARandomInitializer),
		Selector:    ga.NewGATournamentSelector(0.2, 5),
		Breeder:     new(ga.GA2PointBreeder),
		Mutator:     ga.NewGAGaussianMutator(0.4, 0),
		PMutate:     0.5,
		PBreed:      0.2}

	// Second parameter is the number of Optimize Processes.
	gao := ga.NewGAParallel(param, 2)

	genome := ga.NewFloatGenome(make([]float64, 20), rosenbrock, 1, -1)

	gao.Init(1000, genome) //Total population

	gao.OptimizeUntil(func(best ga.GAGenome) bool {
		return best.Score() < 1e-3
	})

	best := gao.Best().(*ga.GAFloatGenome)
	fmt.Printf("%s = %f\n", best, best.Score())
	fmt.Printf("Calls to score = %d\n", scores)
}
