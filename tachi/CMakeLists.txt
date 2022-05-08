cmake_minimum_required(VERSION 3.11)

# Project settings
project(tachi)
if(MSVC)
  add_compile_options(
    $<$<CONFIG:>:/MT>
    $<$<CONFIG:Debug>:/MTd>
    $<$<CONFIG:Release>:/MT>
  )
endif()

file(GLOB SRC_GROUP_MAIN 
  ".editorconfig"
  "src/main.*")

source_group("main" FILES ${SRC_GROUP_MAIN})

include_directories("${PROJECT_SOURCE_DIR}/src")

add_executable(${PROJECT_NAME}
  "${PROJECT_SOURCE_DIR}/src/main.c")

target_link_libraries(${PROJECT_NAME})
