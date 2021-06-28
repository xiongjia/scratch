#                                               -*- Autoconf -*-
# Process this file with autoconf to produce a configure script.

AC_PREREQ([2.71])
AC_INIT(tautotools, 1.0, lexiongjia@gmail.com)
AC_CONFIG_SRCDIR([src/main.cpp])
AC_CONFIG_HEADERS([src/tautotools_config.h])

AH_TOP([
#ifndef _TAUTOTOOLS_H_
#define _TAUTOTOOLS_H_ 1
])

AH_BOTTOM([
#endif
])

AM_INIT_AUTOMAKE()


AC_ARG_ENABLE(examples, 
              AS_HELP_STRING([--disable-examples],[disable building of examples]),
              [],[enable_examples=yes])

AM_CONDITIONAL(BUILD_EXAMPLES,[test $enable_examples != no])

# Checks for programs.
AC_PROG_CXX

# Checks for libraries.

# Checks for header files.
AC_CHECK_HEADERS([stdio.h stdlib.h])

# Checks for typedefs, structures, and compiler characteristics.

# Checks for library functions.


AC_CONFIG_FILES(Makefile src/Makefile)

AC_OUTPUT


echo 
echo "----------------------------------------------------------------"
echo " Config is DONE: "
echo
echo "        Build examples: $enable_examples"
echo "----------------------------------------------------------------"
echo
