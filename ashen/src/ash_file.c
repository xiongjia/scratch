/**
 *
 */

#include <stdio.h>
#include <stdlib.h>

#include "ash_file.h"
#include "ash_string.h"

struct _ash_file_t {
  ash_pool_t *pool;
  const char *filename;
  FILE *file;
};

ASH_DECLARE(ash_file_t *) ash_file_open(ash_pool_t *pool,
    const char *filename, const char *mode) {
  FILE *file;
  ash_file_t *fd;

  if (ASH_STR_IS_NULL_OR_EMPTY(filename) || ASH_STR_IS_NULL_OR_EMPTY(mode)
      || pool == NULL) {
    return NULL;
  }

  fd = ash_pool_alloc(pool, sizeof(ash_file_t));
  if (fd == NULL) {
    return NULL;
  }

  file = fopen(filename, mode);
  if (file == NULL) {
    return NULL;
  }

  fd->pool = pool;
  fd->filename = ash_str_duplicate(pool, filename);
  fd->file = file;
  return fd;
}

ASH_DECLARE(void) ash_file_close(ash_file_t *fd) {
  if (fd == NULL) {
    return;
  }

  if (fd->file != NULL) {
    fclose(fd->file);
    fd->file = NULL;
  }
  ash_pool_free(fd->pool, fd);
}

