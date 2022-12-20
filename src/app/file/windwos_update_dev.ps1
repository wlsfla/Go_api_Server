$hostname = hostname
$winver = (Get-Item "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion") | select @{l="ReleaseId";e={$_.Getvalue("ReleaseId")}},@{l="DisplayVersion";e={$_.GetValue("DisplayVersion")}}
if ([string]::IsNullOrEmpty($winver.DisplayVersion)) { $winver = $winver.ReleaseId } else { $winver = $winver.DisplayVersion}

$url = "http://localhost:7979/api/info_reg/$hostname/$winver"
Invoke-WebRequest -Method GET -Uri $url

