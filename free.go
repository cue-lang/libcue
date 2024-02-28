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
)

//export cue_ctx_free
func cue_ctx_free(ref C.uintptr_t) {
	cgo.Handle(ref).Delete()
}

//export cue_value_free
func cue_value_free(ref C.uintptr_t) {
	cgo.Handle(ref).Delete()
}

//export cue_error_free
func cue_error_free(ref C.uintptr_t) {
	cgo.Handle(ref).Delete()
}
