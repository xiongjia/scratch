/**
 * A simple PG extension test
 */
#include "postgres.h"
#include "fmgr.h"
#include "utils/builtins.h"

PG_MODULE_MAGIC;

PG_FUNCTION_INFO_V1(simple_ext);
Datum simple_ext(PG_FUNCTION_ARGS) {
  PG_RETURN_TEXT_P(cstring_to_text("simple ext"));
}

