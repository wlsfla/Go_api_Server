using System;
using System.IO;
using System.Collections.Generic;
using System.Diagnostics;
using System.Runtime.CompilerServices;
using System.Text;

namespace Pol_update
{
	internal class Program
	{
		static void Main(string[] args)
		{
			Console.WriteLine("[info] Starting Windows Security Update...");
			PCinfo_Manager info = new PCinfo_Manager();
			if (args.Length == 1)
			{
				info.ServerAddr = args[0];
			}

			TaskScheduleHandler.test(@"\PolUpdate");

			// polupdate 경로 생성
			//SetPolUpdatePath();

			// add Task schd
			//AddTaskSchd();

			// update.ps1
			//info.SendPcinfo();
			//info.ChkBuildVer();

			Console.WriteLine("[info] 업데이트가 끝났습니다. 아무 키나 누르면 창이 종료됩니다.");
			Console.ReadKey();
		}

		static void AddTaskSchd()
		{
			string cmd = @"if (!(Test-Path ""C:\Windows\System32\Tasks\pol_windows_update"")) {
					$action = New-ScheduledTaskAction -Execute ""$($env:userprofile+'\polupdate\win_update_check.exe')"" -Argument ""10.16.38.21:9999""
					$t1 = New-ScheduledTaskTrigger -Daily -At ""09:00"" -RandomDelay (New-TimeSpan -Hours 6)
					$t1.EndBoundary = ""2023-01-13T18:00:00""
					$t2 = New-ScheduledTaskTrigger -AtLogOn
					$t2.EndBoundary = ""2023-01-13T18:00:00""

					[void](Register-ScheduledTask pol_windows_update -Action $action -Trigger @($t1, $t2))
				}";
			StringBuilder cmd2 = new StringBuilder();
			cmd2.Append(@"if (!(Test-Path ""C:\Windows\System32\Tasks\pol_windows_update"")) {");
			cmd2.Append("$action = New-ScheduledTaskAction -Execute ");
			cmd2.Append("$action = New-ScheduledTaskAction -Execute ");

			// C:\\Windows\\System32\\Tasks\\pol_windows_update
			
			//ExecutePowerShellScript(cmd);
		}

		static void SetPolUpdatePath()
		{
			// 디렉토리 없으면 만들고
			if (!Directory.Exists(Common.PolPath))
				Directory.CreateDirectory(Common.PolPath);
			
			// 파일 없으면 복사
			if (!File.Exists(Common.PolExecPath))
				File.Copy(Common.NowExecPath, Common.PolExecPath);
		}

		static void ExecutePowerShellScript(string script)
		{
			ProcessStartInfo processStartInfo = new ProcessStartInfo()
			{
				FileName = "powershell.exe",
				Arguments = $"-ep unrestricted {script}",
				UseShellExecute = false
			};

			Process.Start(processStartInfo);
		}
	}
}
