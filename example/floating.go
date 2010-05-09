/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

subset sum solver
*/
package main

import (
	"math"
	"fmt"
	"rand"
	"time"
	"../_obj/ga"
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
	var sum float64
	for i := 1; i < len(g.Gene); i++ {
		sum += 100.0*math.Pow(math.Pow(g.Gene[i]-g.Gene[i-1], 2), 2) + math.Pow(1-g.Gene[i-1], 2)
	}
	return sum
}


func main() {
	rand.Seed(time.Nanoseconds())
	m := ga.NewGAGaussianMutator(0.4, 0)
	b := new(ga.GA2PointBreeder)
	s := ga.NewGATournamentSelector(0.2, 5)

	i := new(ga.GARandomInitializer)
	gao := ga.NewGA(i, s, m, b)
	gao.PMutate = 0.5 //Chance of mutation
	gao.PBreed = 0.2  //Chance of breeding

	fmt.Printf("%s\n", gao)
	genome := ga.NewFloatGenome(make([]float64, 20), rosenbrock, 1, -1)
	gao.Init(100, genome) //Total population
	for {
		gao.Optimize(100) // Run genetic algorithm for 20 generations.
		best := gao.Best().(*ga.GAFloatGenome)
		fmt.Printf("%s = %f\n", best, best.Score())
	}
	fmt.Printf("Calls to score = %d\n", scores)
}
