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
	"cuelang.org/go/cue"
)

//export cue_fields
func cue_fields(v C.cue_value, count *C.size_t) *C.cue_value {
	iter, err := cueValue(v).Fields()
	if err != nil {
		return nil
	}

	var fields []cue.Value
	for iter.Next() {
		fields = append(fields, iter.Value())
	}

	*count = C.size_t(len(fields))
	
	if len(fields) == 0 {
		return nil
	}

	s, ptr := calloc[C.cue_value](len(fields), C.sizeof_cue_value)
	for i, f := range fields {
		s[i] = cueValueHandle(f)
	}

	return ptr
}
