@echo off


setlocal

set _emacs_bin=emacs
set HOME=%_emacs_home%
%_emacs_bin% --batch --script %*
endlocal

