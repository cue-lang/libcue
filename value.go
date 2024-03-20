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

	"cuelang.org/go/cue"
)

//export cue_compile_string_raw
func cue_compile_string_raw(ctx C.cue_ctx, str *C.char, opts *C.struct_cue_bopt, len C.size_t, v *C.cue_value) C.cue_error {
	bopts := buildOptions(opts, len)
	val := cueContext(ctx).CompileString(C.GoString(str), bopts...)

	if err := val.Err(); err != nil {
		return cueErrorHandle(err)
	}
	*v = cueValueHandle(val)
	return 0
}

//export cue_compile_bytes_raw
func cue_compile_bytes_raw(ctx C.cue_ctx, buf unsafe.Pointer, bufLen C.size_t, opts *C.struct_cue_bopt, optLen C.size_t, v *C.cue_value) C.cue_error {
	bopts := buildOptions(opts, optLen)
	val := cueContext(ctx).CompileBytes(C.GoBytes(buf, C.int(bufLen)), bopts...)

	if err := val.Err(); err != nil {
		return cueErrorHandle(err)
	}
	*v = cueValueHandle(val)
	return 0
}

//export cue_top
func cue_top(ctx C.cue_ctx) C.cue_value {
	val := cueContext(ctx).CompileString("_")
	return cueValueHandle(val)
}

//export cue_bottom
func cue_bottom(ctx C.cue_ctx) C.cue_value {
	val := cueContext(ctx).CompileString("_|_")
	return cueValueHandle(val)
}

//export cue_unify
func cue_unify(x C.cue_value, y C.cue_value) C.cue_value {
	u := cueValue(x).Unify(cueValue(y))
	return cueValueHandle(u)
}

//export cue_instance_of_raw
func cue_instance_of_raw(v C.cue_value, typ C.cue_value, opts *C.struct_cue_eopt, len C.size_t) C.cue_error {
	err := cueValue(typ).Subsume(cueValue(v), options(opts, len)...)
	if err != nil {
		return cueErrorHandle(err)
	}
	return 0
}

//export cue_lookup_string
func cue_lookup_string(v C.cue_value, str *C.char, res *C.cue_value) C.cue_error {
	path := cue.ParsePath(C.GoString(str))
	target := cueValue(v).LookupPath(path)
	if err := target.Err(); err != nil {
		return cueErrorHandle(err)
	}
	*res = cueValueHandle(target)
	return 0
}

//export cue_from_int64
func cue_from_int64(ctx C.cue_ctx, v C.int64_t) C.cue_value {
	val := cueContext(ctx).Encode(int64(v))
	return cueValueHandle(val)
}

//export cue_from_uint64
func cue_from_uint64(ctx C.cue_ctx, v C.uint64_t) C.cue_value {
	val := cueContext(ctx).Encode(uint64(v))
	return cueValueHandle(val)
}

//export cue_from_bool
func cue_from_bool(ctx C.cue_ctx, v C.bool) C.cue_value {
	val := cueContext(ctx).Encode(bool(v))
	return cueValueHandle(val)
}

//export cue_from_double
func cue_from_double(ctx C.cue_ctx, v C.double) C.cue_value {
	val := cueContext(ctx).Encode(float64(v))
	return cueValueHandle(val)
}

//export cue_from_string
func cue_from_string(ctx C.cue_ctx, str *C.char) C.cue_value {
	val := cueContext(ctx).Encode(C.GoString(str))
	return cueValueHandle(val)
}

//export cue_from_bytes
func cue_from_bytes(ctx C.cue_ctx, buf unsafe.Pointer, len C.size_t) C.cue_value {
	val := cueContext(ctx).Encode(C.GoBytes(buf, C.int(len)))
	return cueValueHandle(val)
}

//export cue_dec_int64
func cue_dec_int64(v C.cue_value, res *C.int64_t) C.cue_error {
	n, err := cueValue(v).Int64()
	if err != nil {
		return cueErrorHandle(err)
	}
	*res = C.int64_t(n)
	return 0
}

//export cue_dec_uint64
func cue_dec_uint64(v C.cue_value, res *C.uint64_t) C.cue_error {
	n, err := cueValue(v).Uint64()
	if err != nil {
		return cueErrorHandle(err)
	}
	*res = C.uint64_t(n)
	return 0
}

//export cue_dec_bool
func cue_dec_bool(v C.cue_value, res *C.bool) C.cue_error {
	b, err := cueValue(v).Bool()
	if err != nil {
		return cueErrorHandle(err)
	}
	*res = C.bool(b)
	return 0
}

//export cue_dec_double
func cue_dec_double(v C.cue_value, res *C.double) C.cue_error {
	x, err := cueValue(v).Float64()
	if err != nil {
		return cueErrorHandle(err)
	}
	*res = C.double(x)
	return 0
}

//export cue_dec_string
func cue_dec_string(v C.cue_value, res **C.char) C.cue_error {
	s, err := cueValue(v).String()
	if err != nil {
		return cueErrorHandle(err)
	}
	*res = C.CString(s)
	return 0
}

//export cue_dec_bytes
func cue_dec_bytes(v C.cue_value, res *unsafe.Pointer, size *C.size_t) C.cue_error {
	b, err := cueValue(v).Bytes()
	if err != nil {
		return cueErrorHandle(err)
	}
	*res = C.CBytes(b)
	*size = C.size_t(len(b))
	return 0
}

//export cue_dec_json
func cue_dec_json(v C.cue_value, res *unsafe.Pointer, size *C.size_t) C.cue_error {
	b, err := cueValue(v).MarshalJSON()
	if err != nil {
		return cueErrorHandle(err)
	}
	*res = C.CBytes(b)
	*size = C.size_t(len(b))
	return 0
}

//export cue_validate_raw
func cue_validate_raw(v C.cue_value, opts *C.struct_cue_eopt, len C.size_t) C.cue_error {
	err := cueValue(v).Validate(options(opts, len)...)
	if err != nil {
		return cueErrorHandle(err)
	}
	return 0
}

//export cue_default
func cue_default(v C.cue_value, res *C.bool) C.cue_value {
	def, ok := cueValue(v).Default()
	if res == nil {
		return cueValueHandle(def)
	}
	*res = C.bool(ok)
	return cueValueHandle(def)
}

//export cue_is_equal
func cue_is_equal(x C.cue_value, y C.cue_value) C.bool {
	return C.bool(cueValue(x).Equals(cueValue(y)))
}

//export cue_concrete_kind
func cue_concrete_kind(v C.cue_value) C.cue_kind {
	switch cueValue(v).Kind() {
	case cue.BottomKind:
		return C.CUE_KIND_BOTTOM
	case cue.NullKind:
		return C.CUE_KIND_NULL
	case cue.BoolKind:
		return C.CUE_KIND_BOOL
	case cue.IntKind:
		return C.CUE_KIND_INT
	case cue.FloatKind:
		return C.CUE_KIND_FLOAT
	case cue.StringKind:
		return C.CUE_KIND_STRING
	case cue.BytesKind:
		return C.CUE_KIND_BYTES
	case cue.StructKind:
		return C.CUE_KIND_STRUCT
	case cue.ListKind:
		return C.CUE_KIND_LIST
	case cue.NumberKind:
		return C.CUE_KIND_NUMBER
	case cue.TopKind:
		return C.CUE_KIND_TOP
	}
	panic("unreachable")
}

//export cue_incomplete_kind
func cue_incomplete_kind(v C.cue_value) C.cue_kind {
	switch cueValue(v).IncompleteKind() {
	case cue.BottomKind:
		return C.CUE_KIND_BOTTOM
	case cue.NullKind:
		return C.CUE_KIND_NULL
	case cue.BoolKind:
		return C.CUE_KIND_BOOL
	case cue.IntKind:
		return C.CUE_KIND_INT
	case cue.FloatKind:
		return C.CUE_KIND_FLOAT
	case cue.StringKind:
		return C.CUE_KIND_STRING
	case cue.BytesKind:
		return C.CUE_KIND_BYTES
	case cue.StructKind:
		return C.CUE_KIND_STRUCT
	case cue.ListKind:
		return C.CUE_KIND_LIST
	case cue.NumberKind:
		return C.CUE_KIND_NUMBER
	case cue.TopKind:
		return C.CUE_KIND_TOP
	}
	panic("unreachable")
}

//export cue_value_error
func cue_value_error(v C.cue_value) C.cue_error {
	if err := cueValue(v).Err(); err != nil {
		return cueErrorHandle(err)
	}
	return 0
}
