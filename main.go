package main

import (
	"fmt";
	"rand";
	"time";
)

func score (g GAGenome) {
}
func main() {
	rand.Seed(time.Nanoseconds());
	m := new(GASwitchMutator);
	b := new(GARandomBreeder);
	s := NewRouletteSelector(score);
	i := new(GAShuffleInitializer);
	ga := NewGA(i, s, m, b);
	fmt.Printf("%s\n", ga);
	genome := NewOrderedIntGenome([]int{1,2,3,4,5,6,7,8,9,0});
	ga.Init(10, genome);
	ga.PrintPop();
}
