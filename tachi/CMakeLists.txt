cmake_minimum_required(VERSION 3.11)

# Project settings
project(tachi)
if(MSVC)
  add_compile_options(
    $<$<CONFIG:>:/MT>
    $<$<CONFIG:Debug>:/MTd>
    $<$<CONFIG:Release>:/MT>
  )
  set(CMAKE_C_STANDARD_LIBRARIES "ws2_32.lib")
  set(CMAKE_CXX_STANDARD_LIBRARIES "ws2_32.lib")
endif()

file(GLOB SRC_GROUP_CORE_LIB
  "src/tachi-proc.*"
  "src/tachi-misc.*"
  "src/tachi-config.*")

file(GLOB SRC_GROUP_TEST
  "src/tachi-test.*")

file(GLOB SRC_GROUP_MAIN ".editorconfig" "mian.c")

source_group("core" FILES ${SRC_GROUP_CORE_LIB})
source_group("main" FILES ${SRC_GROUP_MAIN})
source_group("test" FILES ${SRC_GROUP_TEST})

add_library(libtachi STATIC
  "${PROJECT_SOURCE_DIR}/src/tachi-proc.c"
  "${PROJECT_SOURCE_DIR}/src/tachi-proc.h"
  "${PROJECT_SOURCE_DIR}/src/tachi-misc.c"
  "${PROJECT_SOURCE_DIR}/src/tachi-misc.h"
  "${PROJECT_SOURCE_DIR}/src/tachi-config.c"
  "${PROJECT_SOURCE_DIR}/src/tachi-config.h")

include_directories("${PROJECT_SOURCE_DIR}/src")

add_executable(${PROJECT_NAME}
  "${PROJECT_SOURCE_DIR}/src/main.c")

add_executable(tachitest
  "${PROJECT_SOURCE_DIR}/src/tachi-test.c")

target_link_libraries(${PROJECT_NAME} libtachi)


enable_testing()
add_test(NAME tachitest_test1 COMMAND tachitest all)
set_tests_properties(tachitest_test1 PROPERTIES PASS_REGULAR_EXPRESSION "PASS")
