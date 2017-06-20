@echo off


setlocal

set _emacs_bin=emacs

%_emacs_bin% --batch --script %*
endlocal

