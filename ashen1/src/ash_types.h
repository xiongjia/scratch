/**
 */

#ifndef _ASH_TYPES_H_
#define _ASH_TYPES_H_ 1

typedef unsigned __int64 size_t;

typedef unsigned char uchar_t;
typedef char char_t;

typedef unsigned long uint32_t;
typedef long int32_t;

typedef char boolean_t;
#define  ASH_FALSE   ((boolean_t)0)
#define  ASH_TRUE    ((boolean_t)1)

#ifndef NULL
# define NULL   ((void *)0)
#endif

#endif /* !defined(_ASH_TYPES_H_) */

