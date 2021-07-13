
set(main_src ash_main.c)
set(libash_src ash_alloc.c ash_alloc.h)


include_directories(${CMAKE_BINARY_DIR}/src)
set(EXAMPLE "This is an example")
configure_file(ash_config.h.in ash_config.h)

add_executable(ash ${main_src})
add_library(libash STATIC ${libash_src})
target_link_libraries(ash libash)

install(TARGETS ash DESTINATION bin)

