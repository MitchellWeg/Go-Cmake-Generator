cmake_minimum_required(VERSION 3.2)

set(CMAKE_MODULE_PATH ${CMAKE_CURRENT_LIST_DIR}/)

project(hello_world Go)

set(SOURCE_FILES
    src/main.go)

add_executable(main ${SOURCE_FILES})
