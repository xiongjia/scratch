# CMake build script
cmake_minimum_required(VERSION 2.8)
project(anubarak)

if(MSVC)
    # Enable the static libraries on Windows
    foreach(flag_var
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

# finding lua
find_path(LUA_INCLUDE_DIR lua.h)
find_library(LUA_LIBRARY_DEBUG NAMES lua)
find_library(LUA_LIBRARY_RELEASE NAMES lua)

# finding boost
unset(Boost_LIBRARIES)
set(Boost_USE_STATIC ON)
set(Boost_USE_STATIC_LIBS ON)
set(Boost_USE_MULTITHREADED ON)
set(Boost_USE_STATIC_RUNTIME ON)
find_package(Boost REQUIRED COMPONENTS
    program_options)

# include directories
include_directories("${LUA_INCLUDE_DIR}"
    "${Boost_INCLUDE_DIRS}"
    "${PROJECT_SOURCE_DIR}/src/main")

set(dbg_libs
    "${Boost_PROGRAM_OPTIONS_LIBRARY_DEBUG}"
    "${LUA_LIBRARY_DEBUG}")

set(opt_libs
    "${Boost_PROGRAM_OPTIONS_LIBRARY_RELEASE}"
    "${LUA_LIBRARY_RELEASE}")

add_library(libanubarak STATIC
    "${PROJECT_SOURCE_DIR}/src/main/ab_types.h"
    "${PROJECT_SOURCE_DIR}/src/main/ab_misc.h"
    "${PROJECT_SOURCE_DIR}/src/main/ab_misc.cxx"
    "${PROJECT_SOURCE_DIR}/src/main/anubarak.cxx"
    "${PROJECT_SOURCE_DIR}/src/main/anubarak.h")

add_executable(anubarak
    "${PROJECT_SOURCE_DIR}/src/main/main.cxx")

target_link_libraries(anubarak
    debug "${dbg_libs}"
    optimized "${opt_libs}"
    libanubarak)

