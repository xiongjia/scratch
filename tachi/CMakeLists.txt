cmake_minimum_required(VERSION 3.11)

# Project settings
project(tachi)

# boost
unset(Boost_LIBRARIES)
set(Boost_USE_STATIC ON)
set(Boost_USE_STATIC_LIBS ON)
set(Boost_USE_MULTITHREADED ON)
set(Boost_USE_STATIC_RUNTIME ON)
find_package(Boost REQUIRED COMPONENTS
  log log_setup system filesystem date_time thread chrono)

include_directories(
  "${Boost_INCLUDE_DIR}"
  "${PROJECT_SOURCE_DIR}/src")

set(debug_libs
  "${Boost_LOG_LIBRARY_DEBUG}"
  "${Boost_LOG_SETUP_LIBRARY_DEBUG}"
  "${Boost_THREAD_LIBRARY_DEBUG}"
  "${Boost_DATE_TIME_LIBRARY_DEBUG}"
  "${Boost_CHRONO_LIBRARY_DEBUG}"
  "${Boost_FILESYSTEM_LIBRARY_DEBUG}"
  "${Boost_SYSTEM_LIBRARY_DEBUG}")

set(optimized_libs 
  "${Boost_LOG_LIBRARY_RELEASE}"
  "${Boost_LOG_SETUP_LIBRARY_RELEASE}"
  "${Boost_THREAD_LIBRARY_RELEASE}"
  "${Boost_DATE_TIME_LIBRARY_RELEASE}"
  "${Boost_CHRONO_LIBRARY_RELEASE}"
  "${Boost_FILESYSTEM_LIBRARY_RELEASE}"
  "${Boost_SYSTEM_LIBRARY_RELEASE}")

add_executable(${PROJECT_NAME}
  "${PROJECT_SOURCE_DIR}/src/main.cpp")

target_link_libraries(${PROJECT_NAME}
  debug "${debug_libs}" optimized "${optimized_libs}")
