

'
path: "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion"

1) [name: DisplayVersion] 이 있다면 해당 값 출력
2) 없으면 [name: ReleaseId] 출력


'

[string]$Registry_path = "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion"
[string]$Registry_Name = "ReleaseId"
Get-ItemProperty -Path $Registry_path -Name $Registry_Name|Format-List -Property $Registry_Name
