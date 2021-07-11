
set(main_src ash_main.c)

add_executable(ash ${main_src})

# target_link_libraries(myApp
#     libfoo
#     libbar
# )

install(TARGETS ash DESTINATION bin)

