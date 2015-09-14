# CMake build script
# gazlowe - My algorithm tests

cmake_minimum_required(VERSION 2.8)

project(gazlowe)

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
    unit_test_framework system thread date_time chrono)

include_directories("${PROJECT_SOURCE_DIR}"
    "${Boost_INCLUDE_DIRS}")

set(gazlowe_boost_dbg_libs
    "${Boost_SYSTEM_LIBRARY_DEBUG}"
    "${Boost_DATE_TIME_LIBRARY_DEBUG}"
    "${Boost_CHRONO_LIBRARY_DEBUG}"
    "${Boost_THREAD_LIBRARY_DEBUG}"
    "${Boost_UNIT_TEST_FRAMEWORK_LIBRARY_DEBUG}")

set(gazlowe_boost_opt_libs 
    "${Boost_SYSTEM_LIBRARY_RELEASE}"
    "${Boost_DATE_TIME_LIBRARY_RELEASE}"
    "${Boost_CHRONO_LIBRARY_RELEASE}"
    "${Boost_THREAD_LIBRARY_RELEASE}"
    "${Boost_UNIT_TEST_FRAMEWORK_LIBRARY_RELEASE}")

add_executable(gazlowe
    "${PROJECT_SOURCE_DIR}/src/main.cxx"
    "${PROJECT_SOURCE_DIR}/src/g_misc.cxx"
    "${PROJECT_SOURCE_DIR}/src/g_misc.hxx"
    "${PROJECT_SOURCE_DIR}/src/g_sorting_qsort.cxx"
    "${PROJECT_SOURCE_DIR}/src/g_leetcode_single_number.cxx"
    "${PROJECT_SOURCE_DIR}/src/g_leetcode_max_depth_btree.cxx"
    "${PROJECT_SOURCE_DIR}/src/g_leetcode_invert_btree.cxx"
    "${PROJECT_SOURCE_DIR}/src/g_leetcode_del_node_in_a_linked_list.cxx"
    "${PROJECT_SOURCE_DIR}/src/g_leetcode_add_digits.cxx"
    "${PROJECT_SOURCE_DIR}/src/g_leetcode_same_tree.cxx"
    "${PROJECT_SOURCE_DIR}/src/g_leetcode_linked_list_cycle.cxx"
    "${PROJECT_SOURCE_DIR}/src/g_leetcode_majority_element.cxx"
    "${PROJECT_SOURCE_DIR}/src/g_leetcode_queue_using_stacks.cxx"
    "${PROJECT_SOURCE_DIR}/src/g_leetcode_stack_using_queues.cxx"
    "${PROJECT_SOURCE_DIR}/src/g_leetcode_contains_duplicate.cxx")

target_link_libraries(gazlowe
    debug "${gazlowe_boost_dbg_libs}"
    optimized "${gazlowe_boost_opt_libs}")

