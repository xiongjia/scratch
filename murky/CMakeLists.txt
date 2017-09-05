cmake_minimum_required(VERSION 3.8.0)

project(murky)

if (MSVC)
  set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} /W4")
  # Enable the static libraries on Windows (/MT)
  foreach (flag_var
    CMAKE_CXX_FLAGS CMAKE_CXX_FLAGS_DEBUG CMAKE_CXX_FLAGS_RELEASE
    CMAKE_CXX_FLAGS_MINSIZEREL CMAKE_CXX_FLAGS_RELWITHDEBINFO
    CMAKE_C_FLAGS CMAKE_C_FLAGS_DEBUG CMAKE_C_FLAGS_RELEASE
    CMAKE_C_FLAGS_MINSIZEREL CMAKE_C_FLAGS_RELWITHDEBINFO)
    # update to the static version of the run time library
    string(REPLACE "/MD" "/MT" ${flag_var} "${${flag_var}}")
  endforeach()
  set(CMAKE_C_STANDARD_LIBRARIES "")
  set(CMAKE_CXX_STANDARD_LIBRARIES "")
endif()

if (APPLE)
  # Enable C++ 11 on OS X
  add_definitions(-std=c++11 -stdlib=libc++)
endif()

# Dependencies
include(deps/CMakeLists.txt)

# Applications
add_subdirectory(murky-cli)

