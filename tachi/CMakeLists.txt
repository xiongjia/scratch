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

file(GLOB SRC_GROUP_MAIN 
  ".editorconfig"
  "src/main.*"
  "src/tachi-proc.*"
  "src/tachi-config.*")

source_group("main" FILES ${SRC_GROUP_MAIN})

include_directories("${PROJECT_SOURCE_DIR}/src")

add_executable(${PROJECT_NAME}
  "${PROJECT_SOURCE_DIR}/src/main.c"
  "${PROJECT_SOURCE_DIR}/src/tachi-proc.c" "${PROJECT_SOURCE_DIR}/src/tachi-proc.h"
  "${PROJECT_SOURCE_DIR}/src/tachi-config.c" "${PROJECT_SOURCE_DIR}/src/tachi-config.h")

target_link_libraries(${PROJECT_NAME})
