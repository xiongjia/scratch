cmake_minimum_required(VERSION 3.8.0)

project(murky)

# Dependencies
include(deps/CMakeLists.txt)

# Applications
add_subdirectory(murky-cli)

