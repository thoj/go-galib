package main

import (
	"fmt";
	"rand";
	"time";
	"ga";
)

// Boring fitness/score function.
func score(g ga.GAOrderedIntGenome) int {
	var total int;
	for i, c := range g.Gene {
		total += c ^ i
	}
	return total;
}

func main() {
	rand.Seed(time.Nanoseconds());
	m := new(ga.GASwitchMutator);
	b := new(ga.GA2PointBreeder);

	s := new(ga.GATournamentSelector);
	s.Contestants = 4;	//Contestants in Tournament
	s.PElite = 0.3;		//Chance of best contestant winning, chance of next is PElite^2 and so on.

	i := new(ga.GARandomInitializer);
	gao := ga.NewGA(i, s, m, b);
	gao.PMutate = 0.1;	//Chance of mutation
	gao.PBreed = 0.2;	//Chance of breeding

	fmt.Printf("%s\n", gao);
	genome := ga.NewOrderedIntGenome([]int{10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, score);
	gao.Init(200, genome);	//Total population

	gao.Optimize(20);	// Run genetic algorithm for 20 generations.
	gao.PrintTop(10);
}
