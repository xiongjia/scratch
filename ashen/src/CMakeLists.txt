
set(main_src ash_main.c)
set(libash_src 
  ash_alloc.c ash_alloc.h ash_palloc.c ash_palloc.h
  ash_list.c ash_list.h ash_string.c ash_string.h
  ash_hash.c ash_hash.h ash_file.c ash_file.h ash_log.c ash_log.h)


include_directories(${CMAKE_BINARY_DIR}/src)
set(EXAMPLE "This is an example")
configure_file(ash_config.h.in ash_config.h)

add_executable(ash ${main_src})
add_library(libash STATIC ${libash_src})
target_link_libraries(ash libash)

install(TARGETS ash DESTINATION bin)

