#                                               -*- Autoconf -*-
# Process this file with autoconf to produce a configure script.

AC_PREREQ([2.71])
AC_INIT([ashen],[1.0],[lexiongjia@gmail.com])

AC_CANONICAL_HOST
AC_CONFIG_HEADERS([src/ash_config.h])

AH_TOP([
#ifndef _ASH_CONFIGURE_H_
#define _ASH_CONFIGURE_H_ 1
])

AH_BOTTOM([
#endif /* !defined(_ASH_CONFIGURE_H_) */
])
AM_INIT_AUTOMAKE()

# Checks for programs.
AC_PROG_CC
CFLAGS+=" -std=c11 "

# Checks for libraries.

# Checks for header files.

# Checks for typedefs, structures, and compiler characteristics.

# Checks for library functions.

AC_CONFIG_FILES(Makefile src/Makefile)
AC_OUTPUT
