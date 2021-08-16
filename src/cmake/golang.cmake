set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${CMAKE_CURRENT_BINARY_DIR}/lib)
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${CMAKE_CURRENT_BINARY_DIR}/lib)
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_CURRENT_BINARY_DIR}/bin)

# go get function
function(ExternalGoProject_Add TARG)
  add_custom_target(${TARG} env GOPATH=${GOPATH} ${CMAKE_Go_COMPILER} get ${ARGN})
endfunction(ExternalGoProject_Add)

function(add_go_executable NAME)
  file(GLOB GO_SOURCE RELATIVE "${CMAKE_CURRENT_SOURCE_DIR}" "*.go")

  add_custom_command(OUTPUT ${CMAKE_RUNTIME_OUTPUT_DIRECTORY}/.timestamp
    COMMAND ${CMAKE_Go_COMPILER} build -o "${CMAKE_RUNTIME_OUTPUT_DIRECTORY}/${NAME}" ${CMAKE_GO_FLAGS} ${GO_SOURCE}
    COMMENT "${CMAKE_Go_COMPILER} build -o ${CMAKE_RUNTIME_OUTPUT_DIRECTORY}/${NAME} ${CMAKE_GO_FLAGS} ${GO_SOURCE}"
    # DEPENDS ${GO_SOURCE}
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR})

  add_custom_target(${NAME} ALL DEPENDS ${CMAKE_RUNTIME_OUTPUT_DIRECTORY}/.timestamp ${ARGN})

endfunction(add_go_executable)

function(ADD_GO_LIBRARY NAME BUILD_TYPE)
  if(BUILD_TYPE STREQUAL "STATIC")
    set(BUILD_MODE -buildmode=c-archive)
    set(LIB_NAME "lib${NAME}.a")
  else()
    set(BUILD_MODE -buildmode=c-shared)
    if(APPLE)
      set(LIB_NAME "lib${NAME}.dylib")
    else()
      set(LIB_NAME "lib${NAME}.so")
    endif()
  endif()

  file(GLOB GO_SOURCE RELATIVE ${CMAKE_CURRENT_SOURCE_DIR} "*.go")
  add_custom_command(OUTPUT ${OUTPUT_DIR}/.timestamp
    COMMAND env GOPATH=${GOPATH} ${CMAKE_Go_COMPILER} build ${BUILD_MODE} -o "${CMAKE_LIBRARY_OUTPUT_DIRECTORY}/${LIB_NAME}" ${CMAKE_GO_FLAGS} ${GO_SOURCE}
    COMMENT "env GOPATH=${GOPATH} ${CMAKE_Go_COMPILER} build ${BUILD_MODE} -o ${CMAKE_LIBRARY_OUTPUT_DIRECTORY}/${LIB_NAME} ${CMAKE_GO_FLAGS} ${GO_SOURCE}"
    WORKING_DIRECTORY .)

  add_custom_target(${NAME} ALL DEPENDS ${OUTPUT_DIR}/.timestamp ${ARGN})

endfunction(ADD_GO_LIBRARY)
