include $(GOROOT)/src/Make.$(GOARCH)

TARG=ga

GOFILES=src/ga.go \
	src/selector.go \
	src/mutator.go \
	src/mutator_shift.go \
	src/mutator_switch.go \
	src/mutator_multi.go \
	src/initializer.go \
	src/breeder.go \
	src/genome.go \
	src/genome_ordered_int.go \
	src/genome_fixed_bitstring.go \

include $(GOROOT)/src/Make.pkg
