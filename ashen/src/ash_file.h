/**
 */

#ifndef _ASH_FILE_H_
#define _ASH_FILE_H_ 1

#include "ash_general.h"
#include "ash_palloc.h"

typedef struct _ash_file_t ash_file_t;

ASH_DECLARE(ash_file_t *) ash_file_open(ash_pool_t *pool,
    const char *filename, const char *mode);

ASH_DECLARE(void) ash_file_close(ash_file_t *fd);

#endif /* !defined(_ASH_FILE_H_) */

