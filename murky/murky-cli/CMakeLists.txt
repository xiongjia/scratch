# murky cli
file(GLOB murky_cli_src *.cpp)

add_executable(murky-cli ${murky_cli_src})

target_link_libraries(murky-cli
  debug "${ZLIB_STATIC_DBG_LIBRARIES}"
  optimized "${ZLIB_STATIC_REL_LIBRARIES}")

add_dependencies(murky-cli zlib)
include_directories("${PROJECT_SOURCE_DIR}/murky-cli"
  "${ZLIB_INCLUDE_DIRS}")

