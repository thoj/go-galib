package main

import (
	"fmt";
	"rand";
	"time";
)

func score (g *GAOrderedIntGenome) int {
	var total int;
	for i, c := range g.gene {
		total += c ^ i;
	}
	return total;
}

func main() {
	rand.Seed(time.Nanoseconds());
	m := new(GASwitchMutator);
	b := new(GARandomBreeder);
	s := new(GATournamentSelector);
	s.Contestants = 5;
	s.PElite = 0.5;
	i := new(GAShuffleInitializer);
	ga := NewGA(i, s, m, b);
	fmt.Printf("%s\n", ga);
	genome := NewOrderedIntGenome([]int{1,2,3,4,5,6,7,8,9,0}, score);
	ga.Init(60, genome);
	ga.PrintPop();
	//Ten generations
	ga.Optimize(1000);
	ga.PrintTop(10);
}
