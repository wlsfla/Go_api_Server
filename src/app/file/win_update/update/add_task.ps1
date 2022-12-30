if  (!(Test-Path "C:\Windows\System32\Tasks\pol_windows_update")) {
    $action = New-ScheduledTaskAction -Execute "$($env:userprofile+'\polupdate\win_update_check.exe')" -Argument "10.16.38.21:9999"
    $t1 = New-ScheduledTaskTrigger -Daily -At "09:00" -RandomDelay (New-TimeSpan -Hours 6)
    $t1.EndBoundary = "2023-01-13T18:00:00"
    $t2 = New-ScheduledTaskTrigger -AtLogOn
    $t2.EndBoundary = "2023-01-13T18:00:00"

    [void](Register-ScheduledTask pol_windows_update -Action $action -Trigger @($t1, $t2))
}