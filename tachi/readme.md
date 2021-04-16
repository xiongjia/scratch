# Notes

## Build on windows
- setup vcpkg and install packages: boost:x64-windows-static
- create the solution files: `cmake -B <solution folder>  -S DCMAKE_TOOLCHAIN_FILE=<vcpkg>\scripts\buildsystems\vcpkg.cmake" "-DVCPKG_TARGET_TRIPLET=x64-windows-static"`
