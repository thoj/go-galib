/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

go-galib initializers
*/


package main

type GAInitializer interface {
	// Initializes popsize length []GAGenome
	Init(popsize int) ([]GAGenome);
	// String name of initializers
	String() string;
}
