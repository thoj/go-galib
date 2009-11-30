/*
Copyright 2009 Thomas Jager <mail@jager.no> All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

go-galib selectors
*/


package main

type GASelector interface {
	// Select n from pop
	Select(pop []GAGenome, n int) ([]GAGenome);
	// String name of selector
	String() string;
}
