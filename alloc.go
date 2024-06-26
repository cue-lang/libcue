// Copyright 2024 The CUE Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

/*
#include <stdint.h>
#include <stdbool.h>
#include "cue.h"
*/
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

//export cue_free
func cue_free(ref C.uintptr_t) {
	cgo.Handle(ref).Delete()
}

//export cue_free_all_inner_raw
func cue_free_all_inner_raw(p *C.uintptr_t, n C.size_t) {
	for _, ref := range unsafe.Slice(p, n) {
		cue_free(ref)
	}
}

func calloc[T any](n int, size int) ([]T, *T) {
	ptr := (*T)(C.malloc(C.size_t(n * size)))
	s := unsafe.Slice(ptr, n)

	for i := 0; i < n; i++ {
		var zero T
		s[i] = zero
	}
	return s, ptr
}
