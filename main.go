package main

import (
	"fmt";
	"rand";
	"time";
)

// Boring fitness/score function.
func score (g GAOrderedIntGenome) int {
	var total int;
	for i, c := range g.gene {
		total += c ^ i;
	}
	return total;
}

func main() {
	rand.Seed(time.Nanoseconds());
	m := new(GASwitchMutator);
//	b := new(GARandomBreeder);
	b := new(GA2PointBreeder);
	s := new(GATournamentSelector);
	s.Contestants = 4; //Contestants in Tournament
	s.PElite = 0.3; //Chance of best contestant winning, chance of next is PElite^2 and so on.
	i := new(GAShuffleInitializer);
	ga := NewGA(i, s, m, b);
	ga.PMutate = 0.1; //Chance of mutation
	ga.PBreed = 0.2; //Chance of breeding
	fmt.Printf("%s\n", ga);
	genome := NewOrderedIntGenome([]int{10,11,12,13,14,15,16,1,2,3,4,5,6,7,8,9,0}, score);
	ga.Init(200, genome); //Total population
	
	ga.Optimize(20); // Run genetic algorithm for 20 generations.
	ga.PrintTop(10);
}
