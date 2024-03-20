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
#include "cue.h"
#include "build.h"
*/
import "C"
import (
	"unsafe"

	"cuelang.org/go/cue"
)

func buildOptions(opts *C.struct_cue_bopt, len C.size_t) []cue.BuildOption {
	var bopts []cue.BuildOption
	for _, opt := range unsafe.Slice(opts, len) {
		bopts = append(bopts, buildOption(opt))
	}
	return bopts
}

func buildOption(opt C.struct_cue_bopt) cue.BuildOption {
	switch opt.tag {
	case C.CUE_BUILD_FILENAME:
		return cue.Filename(C.GoString(opt.str))
	case C.CUE_BUILD_IMPORT_PATH:
		return cue.ImportPath(C.GoString(opt.str))
	case C.CUE_BUILD_INFER_BUILTINS:
		return cue.InferBuiltins(bool(opt.b))
	case C.CUE_BUILD_SCOPE:
		return cue.Scope(cueValue(opt.value))
	}
	panic("unreachable")
}
