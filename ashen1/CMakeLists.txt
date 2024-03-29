cmake_minimum_required(VERSION 3.11)

project(ashen1)

if(MSVC)
  add_compile_options(
    $<$<CONFIG:>:/MT>
    $<$<CONFIG:Debug>:/MTd>
    $<$<CONFIG:Release>:/MT>
  )
  set(CMAKE_C_STANDARD_LIBRARIES "ws2_32.lib")
  set(CMAKE_CXX_STANDARD_LIBRARIES "ws2_32.lib")
endif()

include_directories("${PROJECT_SOURCE_DIR}/src")

file(GLOB SRC_LIB
  "${PROJECT_SOURCE_DIR}/src/ash_types.h"
  "${PROJECT_SOURCE_DIR}/src/ash_cfg.*"
  "${PROJECT_SOURCE_DIR}/src/ash_hash.*"
  "${PROJECT_SOURCE_DIR}/src/ash_vfmt.*"
  "${PROJECT_SOURCE_DIR}/src/ash_list.*"
  "${PROJECT_SOURCE_DIR}/src/ash_vbuf.*"
  "${PROJECT_SOURCE_DIR}/src/ash_time.*"
  "${PROJECT_SOURCE_DIR}/src/ash_errno.*"
  "${PROJECT_SOURCE_DIR}/src/ash_file.*"
  "${PROJECT_SOURCE_DIR}/src/ash_log.*"
  "${PROJECT_SOURCE_DIR}/src/ash_mm.*"
  "${PROJECT_SOURCE_DIR}/src/ash_str.*"
  "${PROJECT_SOURCE_DIR}/src/ash_misc.*")

file(GLOB SRC_LIB_UNIX
  "${PROJECT_SOURCE_DIR}/src/ash_file_unix.*")

file(GLOB SRC_LIB_WIN
  "${PROJECT_SOURCE_DIR}/src/ash_file_win.*")

source_group("lib" FILES ${SRC_LIB} ${SRC_LIB_UNIX})

add_library(libashen1 STATIC
  "${PROJECT_SOURCE_DIR}/src/ash_types.h"
  "${PROJECT_SOURCE_DIR}/src/ash_cfg.c"
  "${PROJECT_SOURCE_DIR}/src/ash_cfg.h"
  "${PROJECT_SOURCE_DIR}/src/ash_hash.c"
  "${PROJECT_SOURCE_DIR}/src/ash_hash.h"
  "${PROJECT_SOURCE_DIR}/src/ash_vfmt.c"
  "${PROJECT_SOURCE_DIR}/src/ash_vfmt.h"
  "${PROJECT_SOURCE_DIR}/src/ash_list.c"
  "${PROJECT_SOURCE_DIR}/src/ash_list.h"
  "${PROJECT_SOURCE_DIR}/src/ash_vbuf.c"
  "${PROJECT_SOURCE_DIR}/src/ash_vbuf.h"
  "${PROJECT_SOURCE_DIR}/src/ash_time.c"
  "${PROJECT_SOURCE_DIR}/src/ash_time.h"
  "${PROJECT_SOURCE_DIR}/src/ash_errno.c"
  "${PROJECT_SOURCE_DIR}/src/ash_errno.h"
  "${PROJECT_SOURCE_DIR}/src/ash_file_win.c"
  "${PROJECT_SOURCE_DIR}/src/ash_file.c"
  "${PROJECT_SOURCE_DIR}/src/ash_file.h"
  "${PROJECT_SOURCE_DIR}/src/ash_log.c"
  "${PROJECT_SOURCE_DIR}/src/ash_log.h"
  "${PROJECT_SOURCE_DIR}/src/ash_mm.c"
  "${PROJECT_SOURCE_DIR}/src/ash_mm.h"
  "${PROJECT_SOURCE_DIR}/src/ash_str.c"
  "${PROJECT_SOURCE_DIR}/src/ash_str.h"
  "${PROJECT_SOURCE_DIR}/src/ash_misc.c"
  "${PROJECT_SOURCE_DIR}/src/ash_misc.h")

add_executable(${PROJECT_NAME}
  "${PROJECT_SOURCE_DIR}/src/main.c")

add_executable(ashen1-tests
  "${PROJECT_SOURCE_DIR}/src/ash_test_str.c"
  "${PROJECT_SOURCE_DIR}/src/ash_test_hash.c"
  "${PROJECT_SOURCE_DIR}/src/ash_test_time.c"
  "${PROJECT_SOURCE_DIR}/src/ash_test_file.c"
  "${PROJECT_SOURCE_DIR}/src/ash_tests_util.h"
  "${PROJECT_SOURCE_DIR}/src/ash_tests_util.c"
  "${PROJECT_SOURCE_DIR}/src/ash_tests.h"
  "${PROJECT_SOURCE_DIR}/src/ash_tests.c")

target_link_libraries(${PROJECT_NAME} libashen1)
target_link_libraries(ashen1-tests libashen1)
