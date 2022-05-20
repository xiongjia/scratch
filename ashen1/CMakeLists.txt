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

add_library(libashen1 STATIC
  "${PROJECT_SOURCE_DIR}/src/ash_types.h"
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
  "${PROJECT_SOURCE_DIR}/src/ash_tests.h"
  "${PROJECT_SOURCE_DIR}/src/ash_tests.c")

target_link_libraries(${PROJECT_NAME} libashen1)
target_link_libraries(ashen1-tests libashen1)
