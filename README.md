# libcue: use CUE from C and C-like languages

This is a work in progress, expect constant churn and breakage.

## Building

Building both as a static archive and a shared object is supported.

```
go build -o libcue.so -buildmode=c-shared
go build -o libcue.a -buildmode=c-archive
```

## Using

Include `cue.h` in your C code
and link with `libcue`
(either statically or dynamically).

## Issue tracking

Please raise all issues in
[the main CUE repository](https://github.com/cue-lang/cue/issues),
giving the title of the issue a `libcue: ` prefix.
