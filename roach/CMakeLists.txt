# CMake build script
cmake_minimum_required(VERSION 2.8)

project(roach)

if (MSVC)
    # Enable the static libraries on Windows
    foreach (flag_var 
        CMAKE_CXX_FLAGS CMAKE_CXX_FLAGS_DEBUG CMAKE_CXX_FLAGS_RELEASE
        CMAKE_CXX_FLAGS_MINSIZEREL CMAKE_CXX_FLAGS_RELWITHDEBINFO
        CMAKE_C_FLAGS CMAKE_C_FLAGS_DEBUG CMAKE_C_FLAGS_RELEASE
        CMAKE_C_FLAGS_MINSIZEREL CMAKE_C_FLAGS_RELWITHDEBINFO)
        # update to the static version of the run time library
        string(REPLACE "/MD" "/MT" ${flag_var} "${${flag_var}}")
    endforeach()
 
    set(CMAKE_C_STANDARD_LIBRARIES "ws2_32.lib advapi32.lib psapi.lib iphlpapi.lib")
    set(CMAKE_CXX_STANDARD_LIBRARIES "ws2_32.lib advapi32.lib psapi.lib iphlpapi.lib")
endif()

# enable boost static flag
unset(Boost_LIBRARIES)
set(Boost_USE_STATIC ON)
set(Boost_USE_STATIC_LIBS ON)
set(Boost_USE_MULTITHREADED ON)
set(Boost_USE_STATIC_RUNTIME ON)

find_package(Boost REQUIRED COMPONENTS thread date_time system chrono)

# finding libuv ( https://github.com/libuv/libuv )
find_path(UV_INCLUDE_DIR uv.h)
find_library(UV_LIBRARY_DEBUG NAMES libuv)
find_library(UV_LIBRARY_RELEASE NAMES libuv)

# finding http-parser ( https://github.com/joyent/http-parser/ )
find_path(HttpParser_INCLUDE_DIR http_parser.h)
find_library(HttpParser_LIBRARY_DEBUG NAMES http_parser)
find_library(HttpParser_LIBRARY_RELEASE NAMES http_parser)

include_directories("${Boost_INCLUDE_DIRS}"
    "${UV_INCLUDE_DIR}"
    "${HttpParser_INCLUDE_DIR}"
    "${PROJECT_SOURCE_DIR}/src"
    "${PROJECT_SOURCE_DIR}/include")

set(roach_dbg_libs
    "${UV_LIBRARY_DEBUG}"
    "${HttpParser_LIBRARY_DEBUG}"
    "${Boost_DATE_TIME_LIBRARY_DEBUG}"
    "${Boost_SYSTEM_LIBRARY_DEBUG}"
    "${Boost_CHRONO_LIBRARY_DEBUG}"
    "${Boost_THREAD_LIBRARY_DEBUG}")

set(roach_opt_libs
    "${UV_LIBRARY_RELEASE}"
    "${HttpParser_LIBRARY_RELEASE}"
    "${Boost_DATE_TIME_LIBRARY_RELEASE}"
    "${Boost_SYSTEM_LIBRARY_RELEASE}"
    "${Boost_CHRONO_LIBRARY_RELEASE}"
    "${Boost_THREAD_LIBRARY_RELEASE}") 

add_library(roach STATIC
    "${PROJECT_SOURCE_DIR}/src/roach.cxx"
    "${PROJECT_SOURCE_DIR}/include/roach.hxx"
    "${PROJECT_SOURCE_DIR}/src/roach_context.cxx"
    "${PROJECT_SOURCE_DIR}/include/roach_context.hxx"
    "${PROJECT_SOURCE_DIR}/src/roach_log.cxx"
    "${PROJECT_SOURCE_DIR}/include/roach_log.hxx"
    "${PROJECT_SOURCE_DIR}/src/roach_serv.cxx"
    "${PROJECT_SOURCE_DIR}/include/roach_serv.hxx"
    "${PROJECT_SOURCE_DIR}/src/roach_misc.cxx"
    "${PROJECT_SOURCE_DIR}/src/roach_misc.hxx")

add_executable(roach-example
    "${PROJECT_SOURCE_DIR}/example/example.cxx")

target_link_libraries(roach-example
    debug "${roach_dbg_libs}"
    optimized "${roach_opt_libs}"
    roach) 

