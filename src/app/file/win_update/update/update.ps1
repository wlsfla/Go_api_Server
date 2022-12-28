$winver = (Get-Item "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion") | select @{l="ReleaseId";e={$_.Getvalue("ReleaseId")}},@{l="DisplayVersion";e={$_.GetValue("DisplayVersion")}},@{l="CurrentBuildNumber";e={$_.GetValue("CurrentBuildNumber")}},@{l="UBR";e={$_.GetValue("UBR")}},@{l="Hostname";e={hostname}}
if ([string]::IsNullOrEmpty($winver.DisplayVersion)) { $winver.DisplayVersion = $winver.ReleaseId }

# /api/info_reg/:hostname/:winver/:build
$ReleaseId = $winver.ReleaseId
$DisplayVersion = $winver.DisplayVersion
$CurrentBuildNumber = $winver.CurrentBuildNumber
$UBR = $winver.UBR
$Hostname = $winver.Hostname
$pcinfo = "$DisplayVersion.$CurrentBuildNumber.$UBR"
$currbuildver = "$CurrentBuildNumber.$UBR"

[void](Invoke-WebRequest -Method GET -Uri "http://10.16.38.21:9999/api/info_reg/$Hostname/$DisplayVersion/$currbuildver")


$dict = New-Object 'System.Collections.Generic.Dictionary[String, String]'
$dict.Add("1803", "17134.2208")
$dict.Add("1809", "17763.3653")
$dict.Add("1903", "18362.1256")
$dict.Add("1909", "18363.2274")
$dict.Add("20H2", "19042.2311")
$dict.Add("21H1", "19043.2311")
$dict.Add("21H2", "19044.2311")


if (($DisplayVersion -eq "1803") -or !($dict.ContainsKey($DisplayVersion))) {
    # winver 1803은 수동 업데이트
    [void](Invoke-WebRequest -Method GET -Uri "http://10.16.38.21:9999/api/result/2")
    write-host "[info] Cannot Update. ($pcinfo) `n"

    write-host "End Program. Press Enter Key."
    pause
    exit
}

if (($currbuildver -gt $dict[$winver.DisplayVersion]) -or ($currbuildver -eq $dict[$winver.DisplayVersion])) {
    # 현재 빌드버전이 높거나 같으면 pass
    [void](Invoke-WebRequest -Method GET -Uri "http://10.16.38.21:9999/api/result/1")
    write-host "[info] Already Update. ($pcinfo) `n"

    write-host "End Program. Press Enter Key."
    pause
    exit
}

# 현재 빌드버전이 낮으면 update url open
write-host "[info] Access This URL. And Download File In Internet Explorer `n `n >>> http://10.16.38.21:9999/file/$DisplayVersion `n "

pause