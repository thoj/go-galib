package main

import (
)

func doMutate(a GAMutator, g *GAGenome) {
	a.Mutate(g);
}

func main() {
	m := new(GASwitchMutator);
	g := new(GAGenome);
	doMutate(m, g);
}
