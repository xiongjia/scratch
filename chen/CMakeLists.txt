# CMake build script
# Chen - My Network protocol tests
cmake_minimum_required(VERSION 3.11)

project(chen)
set_property(GLOBAL PROPERTY USE_FOLDERS ON)
set(CMAKE_EXPORT_COMPILE_COMMAND ON)

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
  add_definitions(-D_WIN32_WINNT=0x0501)
add_definitions(-DBOOST_CONFIG_SUPPRESS_OUTDATED_MESSAGE)
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
find_package(Boost 1.66.0 REQUIRED COMPONENTS
  unit_test_framework thread date_time system program_options
  chrono filesystem regex)

include_directories("${Boost_INCLUDE_DIRS}"
    "${PROJECT_SOURCE_DIR}/src")

set(chen_dbg_libs
  "${Boost_REGEX_LIBRARY_DEBUG}"
  "${Boost_FILESYSTEM_LIBRARY_DEBUG}"
  "${Boost_DATE_TIME_LIBRARY_DEBUG}"
  "${Boost_SYSTEM_LIBRARY_DEBUG}"
  "${Boost_CHRONO_LIBRARY_DEBUG}"
  "${Boost_THREAD_LIBRARY_DEBUG}"
  "${Boost_PROGRAM_OPTIONS_LIBRARY_DEBUG}"
  "${Boost_UNIT_TEST_FRAMEWORK_LIBRARY_DEBUG}")

set(chen_opt_libs
  "${Boost_REGEX_LIBRARY_RELEASE}"
  "${Boost_FILESYSTEM_LIBRARY_RELEASE}"
  "${Boost_DATE_TIME_LIBRARY_RELEASE}"
  "${Boost_SYSTEM_LIBRARY_RELEASE}"
  "${Boost_CHRONO_LIBRARY_RELEASE}"
  "${Boost_THREAD_LIBRARY_RELEASE}"
  "${Boost_PROGRAM_OPTIONS_LIBRARY_RELEASE}"
  "${Boost_UNIT_TEST_FRAMEWORK_LIBRARY_RELEASE}")

add_library(libchen STATIC
  "${PROJECT_SOURCE_DIR}/src/chen_types.hxx"
  "${PROJECT_SOURCE_DIR}/src/chen_misc.hxx"
  "${PROJECT_SOURCE_DIR}/src/chen_misc.cxx"
  "${PROJECT_SOURCE_DIR}/src/chen_log.hxx"
  "${PROJECT_SOURCE_DIR}/src/chen_log.cxx"
  "${PROJECT_SOURCE_DIR}/src/chen_prog_options.cxx"
  "${PROJECT_SOURCE_DIR}/src/chen_prog_options.hxx"
  "${PROJECT_SOURCE_DIR}/src/chen_serv.hxx"
  "${PROJECT_SOURCE_DIR}/src/chen_serv.cxx"
  "${PROJECT_SOURCE_DIR}/src/chen.hxx")

add_executable(chentest
  "${PROJECT_SOURCE_DIR}/test/chen_test_log.cxx"
  "${PROJECT_SOURCE_DIR}/test/chen_test.cxx")

file(GLOB SOURCES_GROUP_MISC
  "src/chen_log.*"
  "src/chen_prog_options.*"
  "src/chen_misc.*")

file(GLOB SOURCES_GROUP_CORE
  "src/chen.hxx"
  "src/chen_types.hxx")

file(GLOB SOURCES_GROUP_SERV
  "src/chen_serv.*")

file(GLOB SOURCES_GROUP_CFG
  ".editorconfig")

source_group("misc" FILES ${SOURCES_GROUP_MISC})
source_group("core" FILES ${SOURCES_GROUP_CORE})
source_group("serv" FILES ${SOURCES_GROUP_SERV})
source_group("conf" FILES ${SOURCES_GROUP_CONF})

add_executable(chen
  ".editorconfig"
  "${PROJECT_SOURCE_DIR}/src/main.cxx")

target_link_libraries(chen
  debug "${chen_dbg_libs}"
  optimized "${chen_opt_libs}"
  libchen)

target_link_libraries(chentest
  debug "${chen_dbg_libs}"
  optimized "${chen_opt_libs}"
  libchen)

