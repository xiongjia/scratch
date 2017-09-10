# CMake build script
# gazlowe - My algorithm tests

cmake_minimum_required(VERSION 2.8)

project(ssl-scratch)

if (MSVC)
    # Enable the static libraries on Windows
    foreach (flag_var
        CMAKE_CXX_FLAGS CMAKE_CXX_FLAGS_DEBUG
        CMAKE_CXX_FLAGS_RELEASE CMAKE_CXX_FLAGS_MINSIZEREL
        CMAKE_CXX_FLAGS_RELWITHDEBINFO
        CMAKE_C_FLAGS CMAKE_C_FLAGS_DEBUG
        CMAKE_C_FLAGS_RELEASE CMAKE_C_FLAGS_MINSIZEREL
        CMAKE_C_FLAGS_RELWITHDEBINFO)
        # update to the static version of the run time library
        string(REPLACE "/MD" "/MT" ${flag_var} "${${flag_var}}")
    endforeach()

    set(win_sdk_libs
        "${CMAKE_STATIC_LIBRARY_PREFIX}ws2_32${CMAKE_STATIC_LIBRARY_SUFFIX}"
        "${CMAKE_STATIC_LIBRARY_PREFIX}advapi32${CMAKE_STATIC_LIBRARY_SUFFIX}")

    set(CMAKE_C_STANDARD_LIBRARIES  "")
    set(CMAKE_CXX_STANDARD_LIBRARIES "")
endif()

if (APPLE)
    # Enable C++ 11 on OS X
    add_definitions(-std=c++11 -stdlib=libc++)
endif()

find_package(OpenSSL)

include_directories("${PROJECT_SOURCE_DIR}"
    "${OPENSSL_INCLUDE_DIR}")

set(ssl-scratch_dbg_libs
    "${CMAKE_STATIC_LIBRARY_PREFIX}ws2_32${CMAKE_STATIC_LIBRARY_SUFFIX}"
    "${CMAKE_STATIC_LIBRARY_PREFIX}advapi32${CMAKE_STATIC_LIBRARY_SUFFIX}"
    "${CMAKE_STATIC_LIBRARY_PREFIX}user32${CMAKE_STATIC_LIBRARY_SUFFIX}"
    "${CMAKE_STATIC_LIBRARY_PREFIX}crypt32${CMAKE_STATIC_LIBRARY_SUFFIX}"
    "${OPENSSL_LIBRARIES}")

set(ssl-scratch_opt_libs
    "${CMAKE_STATIC_LIBRARY_PREFIX}ws2_32${CMAKE_STATIC_LIBRARY_SUFFIX}"
    "${CMAKE_STATIC_LIBRARY_PREFIX}advapi32${CMAKE_STATIC_LIBRARY_SUFFIX}"
    "${CMAKE_STATIC_LIBRARY_PREFIX}user32${CMAKE_STATIC_LIBRARY_SUFFIX}"
    "${CMAKE_STATIC_LIBRARY_PREFIX}crypt32${CMAKE_STATIC_LIBRARY_SUFFIX}"
    "${OPENSSL_LIBRARIES}")

add_executable(ssl-scratch
    "${PROJECT_SOURCE_DIR}/src/main.cxx")

target_link_libraries(ssl-scratch
    debug "${ssl-scratch_dbg_libs}"
    optimized "${ssl-scratch_dbg_libs}")

