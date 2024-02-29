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

func attrKind(kind C.cue_attr_kind) cue.AttrKind {
	switch kind {
	case C.CUE_ATTR_FIELD:
		return cue.FieldAttr
	case C.CUE_ATTR_DECL:
		return cue.DeclAttr
	case C.CUE_ATTR_VALUE:
		return cue.ValueAttr
	}
	panic("unreachable")
}

// export cue_attrs
func cue_attrs(v C.cue_value, kind C.cue_attr_kind) *C.cue_attr {
	attrs := cueValue(v).Attributes(attrKind(kind))
	if len(attrs) == 0 {
		return nil
	}

	s, ptr := calloc[C.cue_attr](len(attrs)+1, C.sizeof_cue_attr)
	for i, a := range attrs {
		s[i] = cueAttrHandle(a)
	}
	return ptr
}

//export cue_attr_name
func cue_attr_name(a C.cue_attr) *C.char {
	attr := cueAttr(a)
	return C.CString(attr.Name())
}

//export cue_attr_value
func cue_attr_value(a C.cue_attr) *C.char {
	attr := cueAttr(a)
	return C.CString(attr.Contents())
}
