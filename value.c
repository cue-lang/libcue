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

#include "cue.h"
#include "build.h"
#include "option.h"
#include "_cgo_export.h"

cue_error
cue_compile_string(cue_ctx ctx, char *s, cue_bopt *opts, cue_value *v) {
	return cue_compile_string_raw(ctx, s, opts, cue_bopt_len(opts), v);
}

cue_error
cue_compile_bytes(cue_ctx ctx, void *b, size_t len, cue_bopt *opts, cue_value *v) {
	return cue_compile_bytes_raw(ctx, b, len, opts, cue_bopt_len(opts), v);
}

cue_error
cue_instance_of(cue_value x, cue_value y, cue_eopt *opts) {
	return cue_instance_of_raw(x, y, opts, cue_eopt_len(opts));
}

cue_error
cue_validate(cue_value v, cue_eopt *opts) {
	return cue_validate_raw(v, opts, cue_eopt_len(opts));
}
