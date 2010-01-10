all:
	$(MAKE) -C ordered_int/
	mv ordered_int/ordered_int ordered_int.bin
	
	$(MAKE) -C subset_sum/
	mv subset_sum/subset_sum subset_sum.bin

clean:
	$(MAKE) -C ordered_int/ clean
	$(MAKE) -C subset_sum/ clean
	rm *.bin
