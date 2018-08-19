@echo off

:: Dependencies: MS SDK + VS 2017 + Python v2.7
set WIN_SDK=D:\app\Microsoft SDKs\Windows\v7.1\Include
set VS_2017="C:\Program Files (x86)\Microsoft Visual Studio\2017"
set VS_VARS=%VS_2017%\Community\VC\Auxiliary\Build\vcvarsall.bat
set NMAKE=%VS_2017%\Community\VC\Tools\MSVC\14.14.26428\bin\Hostx64\x64\nmake
set PY2_HOME=D:\app\py\python27
set VIM_DEST=.\dest

:: build settings
set SDK_INCLUDE_DIR=%WIN_SDK%
set CPU=AMD64
set TOOLCHAIN=x86_amd64
set FEATURES=BIG
set GUI=yes
set NETBEANS=no
set MBYTE=yes
set DYNAMIC_PYTHON=yes
set PYTHON=%PY2_HOME%
set PYTHON_VER=27

echo %VS_VARS%

call %VS_VARS% %TOOLCHAIN%

