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

#ifndef build_h
#define build_h

#include <stdbool.h>
#include "cue.h"

typedef enum build_option build_option;
typedef struct cue_bopt cue_bopt;

enum build_option {
	BUILD_FILENAME,
	BUILD_IMPORT_PATH,
	BUILD_INFER_BUILTINS,
	BUILD_SCOPE,
};

struct cue_bopt {
	build_option tag;
	cue_value value;
	char *str;
	bool b;
};

#endif // build_h
