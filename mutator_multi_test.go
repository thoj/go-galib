package ga

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

// A mutator that does nothing but keeps track of how many times Mutate() hsa been called.
type testMutator int64

func (m *testMutator) Mutate(_ GAGenome) GAGenome {
	*m++
	return nil
}

func (m *testMutator) String() string {
	return fmt.Sprintf("%v", *m)
}

type testMutators []*testMutator

func (tm testMutators) String() string {
	var out []string
	for _, m := range tm {
		out = append(out, m.String())
	}
	return strings.Join(out, ", ")
}

func buildMutatorsAndMutate(nmut int, iters int64) testMutators {
	mm := NewMultiMutator()
	var mutators testMutators
	for i := 0; i < nmut; i++ {
		tm := testMutator(0)
		mutators = append(mutators, &tm)
		mm.Add(&tm)
	}
	for i := int64(0); i < iters; i++ {
		mm.Mutate(nil)
	}
	return mutators
}

// Computes the min/max ratio for the given test mutators.
func similarity(t *testing.T, mutators testMutators) float64 {
	min, max := int64(math.MaxInt64), int64(0)
	for _, m := range mutators {
		if int64(*m) < min {
			min = int64(*m)
		}
		if int64(*m) > max {
			max = int64(*m)
		}
	}
	return float64(min+1) / float64(max+1)
}

// Tests that each mutator in a MultiMutator is called approximately the same
// number of times.
func TestMultiMutatorEqualProbability(t *testing.T) {
	tests := []struct {
		nmut      int     // Number of mutators.
		iters     int64   // Number of time Mutate() is called.
		threshold float64 // Minimum allowed similarity between mutator call count.
	}{
		{1, 100000, 1},
		{2, 100000, 0.95},
		{10, 100000, 0.90},
		{100, 1000000, 0.90},
	}
	for _, test := range tests {
		mutators := buildMutatorsAndMutate(test.nmut, test.iters)
		if got := similarity(t, mutators); got < test.threshold {
			t.Errorf("Similarity(%v, %v) = %v; want >= %v",
				test.nmut, test.iters, got, test.threshold)
			t.Errorf("Mutator counters: [%s]", mutators)
		}
	}
}
