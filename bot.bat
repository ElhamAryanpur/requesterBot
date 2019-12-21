@echo off
cls
color 0b

set /p url=Please Enter URL: 

:loop
compiled/bot_win64.exe %url% &
goto loop