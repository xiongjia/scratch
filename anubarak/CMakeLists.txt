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
    program_options thread system date_time chrono filesystem program_options)

# finding SDL2
find_path(SDL2_INCLUDE_DIR SDL2/sdl.h)
find_library(SDL2_LIBRARY_DEBUG NAMES SDL2)
find_library(SDL2_LIBRARY_RELEASE NAMES SDL2)
find_library(SDL2_MAIN_LIBRARY_DEBUG NAMES SDL2main)
find_library(SDL2_MAIN_LIBRARY_RELEASE NAMES SDL2main)

# include directories
include_directories("${LUA_INCLUDE_DIR}"
    "${SDL2_INCLUDE_DIR}"
    "${Boost_INCLUDE_DIRS}"
    "${PROJECT_SOURCE_DIR}/src/main")

set(dbg_libs
    "${Boost_PROGRAM_OPTIONS_LIBRARY_DEBUG}"
    "${Boost_THREAD_LIBRARY_DEBUG}"
    "${Boost_SYSTEM_LIBRARY_DEBUG}"
    "${Boost_DATE_TIME_LIBRARY_DEBUG}"
    "${Boost_CHRONO_LIBRARY_DEBUG}"
    "${Boost_FILESYSTEM_LIBRARY_DEBUG}"
    "${Boost_PROGRAM_OPTIONS_LIBRARY_DEBUG}"
    "${SDL2_LIBRARY_DEBUG}"
    "${SDL2_MAIN_LIBRARY_DEBUG}"
    "${LUA_LIBRARY_DEBUG}")

set(opt_libs
    "${Boost_PROGRAM_OPTIONS_LIBRARY_RELEASE}"
    "${Boost_THREAD_LIBRARY_RELEASE}"
    "${Boost_SYSTEM_LIBRARY_RELEASE}"
    "${Boost_DATE_TIME_LIBRARY_RELEASE}"
    "${Boost_CHRONO_LIBRARY_RELEASE}"
    "${Boost_FILESYSTEM_LIBRARY_RELEASE}"
    "${Boost_PROGRAM_OPTIONS_LIBRARY_RELEASE}"
    "${SDL2_LIBRARY_RELEASE}"
    "${SDL2_MAIN_LIBRARY_RELEASE}"
    "${LUA_LIBRARY_RELEASE}")

add_library(libanubarak STATIC
    "${PROJECT_SOURCE_DIR}/src/main/ab_types.h"
    "${PROJECT_SOURCE_DIR}/src/main/ab_logger.h"
    "${PROJECT_SOURCE_DIR}/src/main/ab_logger.cxx"
    "${PROJECT_SOURCE_DIR}/src/main/ab_mod_logger.h"
    "${PROJECT_SOURCE_DIR}/src/main/ab_mod_logger.cxx"
    "${PROJECT_SOURCE_DIR}/src/main/ab_module.h"
    "${PROJECT_SOURCE_DIR}/src/main/ab_module.cxx"
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

