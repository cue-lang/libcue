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

#ifndef cue_h
#define cue_h

#include <stddef.h>
#include <stdint.h>
#include <stdbool.h>

typedef uintptr_t cue_ctx;
typedef uintptr_t cue_value;
typedef uintptr_t cue_error;
typedef uintptr_t cue_attr;
typedef struct cue_bopt cue_bopt;
typedef struct cue_eopt cue_eopt;
typedef struct cue_attr_arg cue_attr_arg;
typedef enum cue_eopt_tag cue_eopt_tag;
typedef enum cue_bopt_tag cue_bopt_tag;
typedef enum cue_attr_kind cue_attr_kind;
typedef enum cue_kind cue_kind;

enum cue_attr_kind {
	CUE_ATTR_FIELD = 1<<0,
	CUE_ATTR_DECL  = 1<<1,
	CUE_ATTR_VALUE = CUE_ATTR_FIELD|CUE_ATTR_DECL,
};

struct cue_attr_arg {
	char *key, *val;
};

enum cue_kind {
	CUE_KIND_BOTTOM,
	CUE_KIND_NULL,
	CUE_KIND_BOOL,
	CUE_KIND_INT,
	CUE_KIND_FLOAT,
	CUE_KIND_STRING,
	CUE_KIND_BYTES,
	CUE_KIND_STRUCT,
	CUE_KIND_LIST,
	CUE_KIND_NUMBER,
	CUE_KIND_TOP,
};

enum cue_eopt_tag {
	CUE_OPT_NONE,
	CUE_OPT_ALL,
	CUE_OPT_ATTR,
	CUE_OPT_CONCRETE,
	CUE_OPT_DEFS,
	CUE_OPT_DISALLOW_CYCLES,
	CUE_OPT_DOCS,
	CUE_OPT_ERRORS_AS_VALUES,
	CUE_OPT_FINAL,
	CUE_OPT_HIDDEN,
	CUE_OPT_INLINE_IMPORTS,
	CUE_OPT_OPTIONALS,
	CUE_OPT_RAW,
	CUE_OPT_SCHEMA,
};

struct cue_eopt {
	cue_eopt_tag tag;
	bool value;
};

enum cue_bopt_tag {
	CUE_BUILD_NONE,
	CUE_BUILD_FILENAME,
	CUE_BUILD_IMPORT_PATH,
	CUE_BUILD_INFER_BUILTINS,
	CUE_BUILD_SCOPE,
};

struct cue_bopt {
	cue_bopt_tag tag;
	cue_value value;
	char *str;
	bool b;
};

#ifdef __cplusplus
extern "C" {
#endif

cue_ctx	cue_newctx(void);
char*	cue_error_string(cue_error);

cue_error	cue_compile_string(cue_ctx, char*, cue_bopt*, cue_value*);
cue_error	cue_compile_bytes(cue_ctx, void*, size_t, cue_bopt*, cue_value*);
cue_value	cue_top(cue_ctx);
cue_value	cue_bottom(cue_ctx);
cue_value	cue_unify(cue_value, cue_value);
cue_error	cue_instance_of(cue_value, cue_value, cue_eopt*);
cue_error	cue_lookup_string(cue_value, char*, cue_value*);
cue_value	cue_from_int64(cue_ctx, int64_t);
cue_value	cue_from_uint64(cue_ctx, uint64_t);
cue_value	cue_from_bool(cue_ctx, bool);
cue_value	cue_from_double(cue_ctx, double);
cue_value	cue_from_string(cue_ctx, char*);
cue_value	cue_from_bytes(cue_ctx, void*, size_t);
cue_error	cue_dec_int64(cue_value, int64_t*);
cue_error	cue_dec_uint64(cue_value, uint64_t*);
cue_error	cue_dec_bool(cue_value, bool*);
cue_error	cue_dec_double(cue_value, double*);
cue_error	cue_dec_string(cue_value, char**);
cue_error	cue_dec_bytes(cue_value, void**, size_t*);
cue_error	cue_dec_json(cue_value, void**, size_t*);
cue_error	cue_validate(cue_value, cue_eopt*);
cue_value	cue_default(cue_value, bool*);
cue_kind	cue_concrete_kind(cue_value);
cue_kind	cue_incomplete_kind(cue_value);
cue_error	cue_value_error(cue_value);
bool	cue_is_equal(cue_value, cue_value);

cue_bopt	cue_filename(char*);
cue_bopt	cue_import_path(char*);
cue_bopt	cue_infer_builtins(bool);
cue_bopt	cue_scope(cue_value);

cue_eopt	cue_all(void);
cue_eopt	cue_attributes(bool);
cue_eopt	cue_concrete(bool);
cue_eopt	cue_definitions(bool);
cue_eopt	cue_disallow_cycles(bool);
cue_eopt	cue_docs(bool);
cue_eopt	cue_errors_as_values(bool);
cue_eopt	cue_final(void);
cue_eopt	cue_hidden(bool);
cue_eopt	cue_inline_imports(bool);
cue_eopt	cue_optionals(bool);
cue_eopt	cue_raw(void);
cue_eopt	cue_schema(void);

cue_attr*	cue_attrs(cue_value, cue_attr_kind, size_t*);
char*	cue_attr_name(cue_attr);
char*	cue_attr_value(cue_attr);
size_t	cue_attr_numargs(cue_attr);
void	cue_attr_getarg(cue_attr, size_t, cue_attr_arg*);

void	cue_free(uintptr_t);
void	cue_free_all(uintptr_t*);
void	libc_free(void*);

#ifdef __cplusplus
}
#endif

#endif // cue_h
