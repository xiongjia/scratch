/*
 */

#ifndef _ASH_GENERAL_H_
#define _ASH_GENERAL_H_ 1

#include <stddef.h>
#include <string.h>

#ifdef __cplusplus
#  define ASH_DECLARE(_type)  extern "C" _type
#else 
#  define ASH_DECLARE(_type)  _type
#endif /* defined(__cplusplus) */

#define ASH_MAX(_x, _y) ((_x) > (_y) ? _x : _y)
#define ASH_MIN(_x, _y) ((_x) < (_y) ? _x : _y)

#define ASH_OK        0
#define ASH_ERROR    -1

#endif /* !defined(_ASH_GENERAL_H_) */

