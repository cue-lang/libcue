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
#include "option.h"
*/
import "C"

import (
	"unsafe"

	"cuelang.org/go/cue"
)

func options(opts *C.struct_cue_eopt, len C.size_t) []cue.Option {
	var eopts []cue.Option
	for _, opt := range unsafe.Slice(opts, len) {
		eopts = append(eopts, option(opt))
	}
	return eopts
}

func option(opt C.struct_cue_eopt) cue.Option {
	switch opt.tag {
	case C.OPT_ALL:
		return cue.All()
	case C.OPT_ATTR:
		return cue.Attributes(bool(opt.value))
	case C.OPT_CONCRETE:
		return cue.Concrete(bool(opt.value))
	case C.OPT_DEFS:
		return cue.Definitions(bool(opt.value))
	case C.OPT_DISALLOW_CYCLES:
		return cue.DisallowCycles(bool(opt.value))
	case C.OPT_DOCS:
		return cue.Docs(bool(opt.value))
	case C.OPT_ERRORS_AS_VALUES:
		return cue.ErrorsAsValues(bool(opt.value))
	case C.OPT_FINAL:
		return cue.Final()
	case C.OPT_HIDDEN:
		return cue.Hidden(bool(opt.value))
	case C.OPT_INLINE_IMPORTS:
		return cue.InlineImports(bool(opt.value))
	case C.OPT_OPTIONALS:
		return cue.Optional(bool(opt.value))
	case C.OPT_RAW:
		return cue.Raw()
	case C.OPT_SCHEMA:
		return cue.Schema()
	}
	panic("unreachable")
}