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
	"unsafe"
)

//export cue_compile_string
func cue_compile_string(ctx C.cue_ctx, str *C.char, opts *C.struct_cue_bopt, len C.size_t) C.cue_value {
	bopts := buildOptions(opts, len)
	val := cueContext(ctx).CompileString(C.GoString(str), bopts...)
	return cueValueHandle(val)
}

//export cue_compile_bytes
func cue_compile_bytes(ctx C.cue_ctx, buf unsafe.Pointer, bufLen C.size_t, opts *C.struct_cue_bopt, optLen C.size_t) C.cue_value {
	bopts := buildOptions(opts, optLen)
	val := cueContext(ctx).CompileBytes(C.GoBytes(buf, C.int(bufLen)), bopts...)
	return cueValueHandle(val)
}

// export cue_top
func cue_top(ctx C.cue_ctx) C.cue_value {
	val := cueContext(ctx).CompileString("_")
	return cueValueHandle(val)
}

// export cue_bottom
func cue_bottom(ctx C.cue_ctx) C.cue_value {
	val := cueContext(ctx).CompileString("_|_")
	return cueValueHandle(val)
}
