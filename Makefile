include $(GOROOT)/src/Make.$(GOARCH)

TARG=ga

GOFILES=ga.go \
	selector.go \
	mutator.go \
	mutator_shift.go \
	mutator_switch.go \
	mutator_multi.go \
	initializer.go \
	breeder.go \
	genome.go \
	genome_ordered_int.go \
	genome_fixed_bitstring.go \

include $(GOROOT)/src/Make.pkg
