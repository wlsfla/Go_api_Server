
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

write-output $hostname
write-output (GetWindowsVersion -Registry_Name $DisplayVersion)


# Get
# Invoke-WebRequest -Method GET -Uri http://localhost:8888/api/Send?param1=PARAM1&param2=PARAM2

# POST
# $text="param1=PARAM1&param2=PARAM2"
# $postParam=[System.Text.Encoding]::UTF8.GetBytes($text)
# Invoke-WebRequest -Method POST -Body $postParam -ContentType "application/x-www-form-urlencoded; charset=utf-8" -Uri http://localhost:8888/api/Send