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

#include <stdlib.h>
#include "cue.h"
#include "option.h"
#include "_cgo_export.h"

size_t
cue_cgo_handle_slice_len(uintptr_t *p) {
	size_t len = 0;

	while(*p != 0) {
		p++;
		len++;
	}
	return len;
}

void
cue_free_all(uintptr_t *p) {
	cue_free_all_inner_raw(p, cue_cgo_handle_slice_len(p));
	free(p);
}