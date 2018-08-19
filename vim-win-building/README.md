# Build Vim in Windows

## Dependencies
- Microsoft Visual Studio ( e.g. VS 2017 Community )
- Python v2.7
- Windows SDK v7.1

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

