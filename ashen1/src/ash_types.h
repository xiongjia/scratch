/**
 */

#ifndef _ASH_TYPES_H_
#define _ASH_TYPES_H_ 1

typedef unsigned __int64 size_t;

typedef unsigned char uchar_t;
typedef char char_t;

typedef unsigned __int64 uint64_t;
typedef unsigned long uint32_t;
typedef long int32_t;

#define boolean_t int
#define ASH_FALSE (0)
#define ASH_TRUE  (1)

#ifndef NULL
# define NULL   ((void *)0)
#endif

#define ASH_ARRAY_COUNT(_src) (sizeof(_src) / sizeof(_src[0]))

#endif /* !defined(_ASH_TYPES_H_) */

