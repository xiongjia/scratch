/*
 */

#ifndef _ASH_GENERAL_H_
#define _ASH_GENERAL_H_ 1

#include <stddef.h>
#include <string.h>

#ifdef __cplusplus
#  define ASH_DECLARE(_type) extern "C" _type
#  define ASH_DECLARE_CDECL(_type) extern "C" _type __cdecl
#else 
#  define ASH_DECLARE(_type) _type
#  define ASH_DECLARE_CDECL(_type) _type __cdecl
#endif /* defined(__cplusplus) */

#define ASH_MAX(_x, _y) ((_x) > (_y) ? _x : _y)
#define ASH_MIN(_x, _y) ((_x) < (_y) ? _x : _y)

#define ASH_OK        0
#define ASH_ERROR    -1

#endif /* !defined(_ASH_GENERAL_H_) */

