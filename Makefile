include $(GOROOT)/src/Make.inc

TARG=ga

GOFILES=ga.go \
	ga_parallel.go \
	selector.go \
	mutator.go \
	mutator_shift.go \
	mutator_switch.go \
	mutator_gaussian.go \
	mutator_multi.go \
	initializer.go \
	breeder.go \
	genome.go \
	genome_ordered_int.go \
	genome_fixed_bitstring.go \
	genome_float64.go \

include $(GOROOT)/src/Make.pkg
