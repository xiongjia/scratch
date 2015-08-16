# CMake build script
# Zeratul - A simple socks proxy. (Boost ASIO)
cmake_minimum_required(VERSION 2.8)

project(zeratul)

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
 
    set(CMAKE_C_STANDARD_LIBRARIES "")
    set(CMAKE_CXX_STANDARD_LIBRARIES "")
    add_definitions(-D_WIN32_WINNT=0x0501)
endif()

if (APPLE)
    # Enable C++ 11 on OS X
    add_definitions(-std=c++11 -stdlib=libc++)
endif()

# enable boost static flag
unset(Boost_LIBRARIES)
set(Boost_USE_STATIC ON)
set(Boost_USE_STATIC_LIBS ON)
set(Boost_USE_MULTITHREADED ON)
set(Boost_USE_STATIC_RUNTIME ON)
find_package(Boost REQUIRED COMPONENTS
    thread date_time system chrono log log_setup filesystem regex)

include_directories("${Boost_INCLUDE_DIRS}"
    "${PROJECT_SOURCE_DIR}/src")

set(zeratul_dbg_libs
    "${Boost_REGEX_LIBRARY_DEBUG}"
    "${Boost_FILESYSTEM_LIBRARY_DEBUG}"
    "${Boost_LOG_LIBRARY_DEBUG}"
    "${Boost_LOG_SETUP_LIBRARY_DEBUG}"
    "${Boost_DATE_TIME_LIBRARY_DEBUG}"
    "${Boost_SYSTEM_LIBRARY_DEBUG}"
    "${Boost_CHRONO_LIBRARY_DEBUG}"
    "${Boost_THREAD_LIBRARY_DEBUG}")

set(zeratul_opt_libs
    "${Boost_REGEX_LIBRARY_RELEASE}"
    "${Boost_FILESYSTEM_LIBRARY_RELEASE}"
    "${Boost_LOG_LIBRARY_RELEASE}"
    "${Boost_LOG_SETUP_LIBRARY_RELEASE}"
    "${Boost_DATE_TIME_LIBRARY_RELEASE}"
    "${Boost_SYSTEM_LIBRARY_RELEASE}"
    "${Boost_CHRONO_LIBRARY_RELEASE}"
    "${Boost_THREAD_LIBRARY_RELEASE}") 

add_library(libzeratul STATIC
    "${PROJECT_SOURCE_DIR}/src/z_types.hxx"
    "${PROJECT_SOURCE_DIR}/src/z_misc.hxx"
    "${PROJECT_SOURCE_DIR}/src/z_misc.cxx"
    "${PROJECT_SOURCE_DIR}/src/zeratul.hxx"
    "${PROJECT_SOURCE_DIR}/src/zeratul.cxx")

add_executable(zeratul
    "${PROJECT_SOURCE_DIR}/src/main.cxx")

target_link_libraries(zeratul
    debug "${zeratul_dbg_libs}"
    optimized "${zeratul_opt_libs}"
    libzeratul)

