package ga

import (
	"math"
	"reflect"
	"testing"
)

func TestRandomMutatorReplaces(t *testing.T) {
	// A float genome with min/max set to 10. After one mutation, exactly one gene must have
	// been replaced with 10.
	g := NewFloatGenome([]float64{0, 0}, nil, 10, 10)
	m := &GAMutatorRandom{}
	gn := m.Mutate(g).(*GAFloatGenome)

	if !reflect.DeepEqual(gn.Gene, []float64{0, 10}) &&
		!reflect.DeepEqual(gn.Gene, []float64{10, 0}) {
		t.Errorf("GAMutatorRandom.Mutate(%v) = %v; want {0, 10} or {10, 0}", g.Gene, gn.Gene)
	}
}

func TestRandomMutatorUniform(t *testing.T) {
	// A float genome with min/max set to 1. We keep mutating it and checking which gene has
	// been mutated, to keep track of the distribution of replacements.
	g := NewFloatGenome([]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil, 1, 1)
	m := &GAMutatorRandom{}

	count := make([]int, len(g.Gene))
	for i := 0; i < 100000; i++ {
		gn := m.Mutate(g).(*GAFloatGenome)
		for i, v := range gn.Gene {
			if v > 0 {
				count[i]++
			}
		}
	}

	min, max := math.MaxInt64, 0
	for _, c := range count {
		if c < min {
			min = c
		}
		if c > max {
			max = c
		}
	}

	if s := float64(min+1) / float64(max+1); s < 0.90 {
		t.Errorf("GARandomMutator replacement counters similarity = %v; want >= 0.9", s)
		t.Errorf("Gene replacement counters: %v", count)
	}
}
