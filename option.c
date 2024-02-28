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
#include "option.h"

cue_eopt
cue_all(void) {
	cue_eopt o = (cue_eopt) {
		.tag = OPT_ALL,
	};

	return o;
}

cue_eopt
cue_attributes(bool b) {
	cue_eopt o = (cue_eopt) {
		.tag = OPT_ATTR,
		.value = b,
	};

	return o;
}

cue_eopt
cue_concrete(bool b) {
	cue_eopt o = (cue_eopt) {
		.tag = OPT_CONCRETE,
		.value = b,
	};

	return o;
}

cue_eopt
cue_definitions(bool b) {
	cue_eopt o = (cue_eopt) {
		.tag = OPT_DEFS,
		.value = b,
	};

	return o;
}

cue_eopt
cue_disallow_cycles(bool b) {
	cue_eopt o = (cue_eopt) {
		.tag = OPT_DISALLOW_CYCLES,
		.value = b,
	};

	return o;
}

cue_eopt
cue_docs(bool b) {
	cue_eopt o = (cue_eopt) {
		.tag = OPT_DOCS,
		.value = b,
	};

	return o;
}

cue_eopt
cue_errors_as_values(bool b) {
	cue_eopt o = (cue_eopt) {
		.tag = OPT_ERRORS_AS_VALUES,
		.value = b,
	};

	return o;
}

cue_eopt
cue_final(void) {
	cue_eopt o = (cue_eopt) {
		.tag = OPT_FINAL,
	};

	return o;
}

cue_eopt
cue_hidden(bool b) {
	cue_eopt o = (cue_eopt) {
		.tag = OPT_HIDDEN,
		.value = b,
	};

	return o;
}

cue_eopt
cue_inline_imports(bool b) {
	cue_eopt o = (cue_eopt) {
		.tag = OPT_INLINE_IMPORTS,
		.value = b,
	};

	return o;
}

cue_eopt
cue_optionals(bool b) {
	cue_eopt o = (cue_eopt) {
		.tag = OPT_OPTIONALS,
		.value = b,
	};

	return o;
}

cue_eopt
cue_raw(void) {
	cue_eopt o = (cue_eopt) {
		.tag = OPT_RAW,
	};

	return o;
}

cue_eopt
cue_schema(void) {
	cue_eopt o = (cue_eopt) {
		.tag = OPT_SCHEMA,
	};

	return o;
}