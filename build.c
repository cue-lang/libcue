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

size_t
cue_bopt_len(cue_bopt *opts) {
	size_t len = 0;

	while(opts->tag != CUE_BUILD_NONE) {
		opts++;
		len++;
	}
	return len;
}

cue_bopt
cue_filename(char *s) {
	cue_bopt bo = (cue_bopt) {
		.tag = CUE_BUILD_FILENAME,
		.str = s,
	};

	return bo;
}

cue_bopt
cue_import_path(char *s) {
	cue_bopt bo = (cue_bopt) {
		.tag = CUE_BUILD_IMPORT_PATH,
		.str = s,
	};

	return bo;
}

cue_bopt
cue_infer_builtins(bool b) {
	cue_bopt bo = (cue_bopt) {
		.tag = CUE_BUILD_INFER_BUILTINS,
		.b = b,
	};

	return bo;
}

cue_bopt
cue_scope(cue_value v) {
	cue_bopt bo = (cue_bopt) {
		.tag = CUE_BUILD_SCOPE,
		.value = v,
	};

	return bo;
}
