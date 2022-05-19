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

add_executable(${PROJECT_NAME}
  "${PROJECT_SOURCE_DIR}/src/main.c")

target_link_libraries(${PROJECT_NAME})

