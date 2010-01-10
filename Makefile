include $(GOROOT)/src/Make.$(GOARCH)

TARG=ga
GOFILES=ga.go \
	selectors.go \
	mutators.go \
	initializers.go \
	breeders.go \
	genomes/genomes.go \
	genomes/ordered_int.go \
	genomes/fixed_bitstring.go \

include $(GOROOT)/src/Make.pkg 
