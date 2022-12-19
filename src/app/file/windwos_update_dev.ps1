
[string]$hostname = hostname
[string]$DisplayVersion = 'DisplayVersion'
[string]$ReleaseId = "ReleaseId"
[string]$Product_Name = 'ProductName'

function GetWindowsVersion {
    param (
        [string]$Registry_path = "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion",
        [string]$Registry_Name
    )

    return (Get-ItemProperty -Path $Registry_path -Name $Registry_Name|Select-Object -Property $Registry_Name).DisplayVersion
}


$DisplayVersion = GetWindowsVersion -Registry_Name $DisplayVersion

$url = "http://localhost:7979/api/info_reg/$hostname/$DisplayVersion"
Invoke-WebRequest -Method GET -Uri $url -OutFile($)


# $winver = (Get-Item "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion") | select @{l="ReleaseId";e={$_.Getvalue("ReleaseId")}},@{l="DisplayVersion";e={$_.GetValue("DisplayVersion")}}