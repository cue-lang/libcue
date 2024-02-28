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

	"cuelang.org/go/cue"
)

func cueContext(ctx C.cue_ctx) *cue.Context {
	return cgo.Handle(ctx).Value().(*cue.Context)
}

func cueContextHandle(ctx *cue.Context) C.cue_ctx {
	return C.cue_ctx(cgo.NewHandle(ctx))
}

func cueValue(x C.cue_value) cue.Value {
	return cgo.Handle(x).Value().(cue.Value)
}

func cueValueHandle(v cue.Value) C.cue_value {
	return C.cue_value(cgo.NewHandle(v))
}

func cueError(x C.cue_error) error {
	return cgo.Handle(x).Value().(error)
}

func cueErrorHandle(err error) C.cue_error {
	return C.cue_value(cgo.NewHandle(err))
}