@echo off
for /f "tokens=*" %%f in (us.txt) do (scan.exe %%f 22 15000)
pause
