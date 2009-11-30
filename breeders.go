/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

go-galib beeders
*/


package main

type GABreeder interface {
	// Breeds two parent GAGenomes and returns two children
	Breed(a, b GAGenome) (GAGenome, GAGenome);
	// String name of breeder
	String() string;
}
