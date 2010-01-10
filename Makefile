include $(GOROOT)/src/Make.$(GOARCH)

TARG=ga
GOFILES=ga.go \
	selectors.go \
	mutators/mutators.go \
	mutators/shift.go \
	mutators/switch.go \
	mutators/multi.go \
	initializers.go \
	breeders.go \
	genomes/genomes.go \
	genomes/ordered_int.go \
	genomes/fixed_bitstring.go \

include $(GOROOT)/src/Make.pkg 
