// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build cgo

package cgobench

/*
static void empty() {
}
*/
import "C"

func Empty() {
	C.empty()
}
