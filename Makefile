include $(GOROOT)/src/Make.$(GOARCH)

TARG=ga
GOFILES=ga.go selectors.go mutators.go initializers.go breeders.go genomes.go

include $(GOROOT)/src/Make.pkg 
