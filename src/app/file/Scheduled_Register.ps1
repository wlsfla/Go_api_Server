$hostname = hostname
$winver = (Get-Item "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion") | select @{l="ReleaseId";e={$_.Getvalue("ReleaseId")}},@{l="DisplayVersion";e={$_.GetValue("DisplayVersion")}}
if ([string]::IsNullOrEmpty($winver.DisplayVersion)) { $winver = $winver.ReleaseId } else { $winver = $winver.DisplayVersion}

Invoke-WebRequest -Method GET -Uri "http://localhost:7979/api/info_reg/$hostname/$winver"

Clear-Host
if ($winver -eq "1803") {
    Write-Host "[error] Cannot Update"
    Invoke-WebRequest -Method GET -Uri "http://localhost:7979/api/result/2"
    
    Pause
    exit
}

$file = (join-Path -Path ([Environment]::GetEnvironmentVariable('userprofile')) -ChildPath "windwos_update_dev.ps1")
$action = New-ScheduledTaskAction -Execute "powershell.exe" -Argument "-w hidden -noni -nop -ep Bypass -Command ""$file"""
$trigger = New-ScheduledTaskTrigger -Daily -At "08:50" -RandomDelay (New-TimeSpan -Hours 8) # random time 08:50 to 16:50 
$trigger.EndBoundary = "2022-12-30T18:00:00" # expired date
$task = New-ScheduledTask -Action $action -Trigger $trigger
Register-ScheduledTask pol_windows_update -InputObject $task

Clear-Host