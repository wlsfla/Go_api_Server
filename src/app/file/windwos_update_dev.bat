@echo off
echo "Windows Update is starting"
PowerShell -NoProfile -ExecutionPolicy Bypass -Command "Invoke-WebRequest -Method GET -Uri http://localhost:7979/update/ps -OutFile('windwos_update_dev.ps1')"
PowerShell -NoProfile -ExecutionPolicy Bypass -Command ".\windwos_update_dev.ps1"
pause
