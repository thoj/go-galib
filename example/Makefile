all:
	$(MAKE) -C src/ordered_int/
	mv src/ordered_int/ordered_int ordered_int
 
	$(MAKE) -C src/subset_sum/
	mv src/subset_sum/subset_sum subset_sum
 
clean:
	$(MAKE) -C src/ordered_int/ clean
	$(MAKE) -C src/subset_sum/ clean
	rm ordered_int subset_sum
 
