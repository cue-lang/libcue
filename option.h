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

#ifndef option_h
#define option_h

#include <stdbool.h>

typedef enum option option;
typedef struct cue_eopt cue_eopt;

enum option {
	OPT_ALL,
	OPT_ATTR,
	OPT_CONCRETE,
	OPT_DEFS,
	OPT_DISALLOW_CYCLES,
	OPT_DOCS,
	OPT_ERRORS_AS_VALUES,
	OPT_FINAL,
	OPT_HIDDEN,
	OPT_INLINE_IMPORTS,
	OPT_OPTIONALS,
	OPT_RAW,
	OPT_SCHEMA,
};

struct cue_eopt {
	option tag;
	bool value;
};

size_t	cue_eopt_len(cue_eopt*);

#endif // option_h
