# Build Vim in Windows

## Dependencies
- Microsoft Visual Studio ( e.g. VS 2017 Community )
- Python v2.7
- Windows SDK v7.1

## Settings
All the settings is saved in `configure.cmd`.   
```
set WIN_SDK=D:\app\Microsoft SDKs\Windows\v7.1\Include
set VS_2017="C:\Program Files (x86)\Microsoft Visual Studio\2017"
set VS_VARS=%VS_2017%\Community\VC\Auxiliary\Build\vcvarsall.bat
set NMAKE=%VS_2017%\Community\VC\Tools\MSVC\14.14.26428\bin\Hostx64\x64\nmake
set PY2_HOME=D:\app\py\python27
set VIM_DEST=.\dest

set FEATURES=BIG
```

## Folders
```
~ vim-win-building/
  + build/
  + dest/
  + vim-src/
```
- `build` - the building script filesa
- `vim-src` - vim source ( https://github.com/vim/vim )
- `dest` - the folder of vim output files

## Usage:
- Enable building settings: `.\build\configure.cmd`
- Building vim via MS nmake: `.\build\build.cmd`
- Clean building files: `.\build\build.cmd clean`
- Copy Vim files to dest folder: `.\build\copy-vim.cmd`

## Links
- http://blog.mgiuffrida.com/2015/06/27/building-vim-on-windows.html

