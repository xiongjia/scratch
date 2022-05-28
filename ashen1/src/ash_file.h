/*
 */

#ifndef _ASH_FILE_H_
#define _ASH_FILE_H_ 1

#include "ash_types.h"

const char *ash_file_get_filename(const char *full_filename);

void ash_file_write_stdout(const char *fmt, ...);

void ash_file_vwrite_stdout(const char *fmt, va_list ap);

void ash_file_write_stderr(const char *fmt, ...);

void ash_file_vwrite_stderr(const char *fmt, va_list ap);

#endif /* !defined(_ASH_FILE_H_) */
