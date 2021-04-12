cmake_minimum_required(VERSION 3.11)

project(tachi)
find_package(Boost REQUIRED)

include_directories(
  "${Boost_INCLUDE_DIR}"
  "${PROJECT_SOURCE_DIR}/src")

add_executable(
  ${PROJECT_NAME}
  "${PROJECT_SOURCE_DIR}/src/main.cpp")

target_link_libraries(${PROJECT_NAME})

