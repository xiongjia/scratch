pushd vim-src\src
if /i [%1] == [clean] (
  call %NMAKE% -f Make_mvc.mak clean
) else (
  call %NMAKE% -f Make_mvc.mak
)
popd
