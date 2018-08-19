@set SRC=vim-src
@set DEST=%VIM_DEST%

xcopy %SRC%\runtime %DEST% /D /S /H /I /Y
xcopy %SRC%\src\xxd\xxd.exe %DEST% /D /Y
xcopy %SRC%\src\GvimExt\gvimext.dll %DEST% /D /Y
xcopy %SRC%\src\*.exe %DEST% /D /Y
