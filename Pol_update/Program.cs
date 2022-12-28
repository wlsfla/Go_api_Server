using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Pol_update
{
	internal class Program
	{
		static void Main(string[] args)
		{
			Console.WriteLine("[info] Starting Windows Security Update...");

			// polupdate 경로 생성
			//SetPolUpdatePath();

			// add Task schd
			//AddTaskSchd();

			// update.ps1
			Console.WriteLine(
				new PCinfo_Manager().ServerAddr = ""
			);


			//Console.WriteLine("End");
			//Console.ReadKey();
		}

		static void AddTaskSchd()
		{
			// add Task.ps1
			ExecutePowerShellScript(
				@"if  (!(Test-Path ""C:\Windows\System32\Tasks\pol_windows_update"")) {
					$action = New-ScheduledTaskAction -Execute ""$($env:userprofile+'\polupdate\win_update_check.exe')"" -Argument ""10.16.38.21:9999""
					$t1 = New-ScheduledTaskTrigger -Daily -At ""09:00"" -RandomDelay (New-TimeSpan -Hours 6)
					$t1.EndBoundary = ""2023-01-13T18:00:00""
					$t2 = New-ScheduledTaskTrigger -AtLogOn
					$t2.EndBoundary = ""2023-01-13T18:00:00""

					[void](Register-ScheduledTask pol_windows_update -Action $action -Trigger @($t1, $t2))
				}"
			);
		}

		static void SetPolUpdatePath()
		{
			ExecutePowerShellScript(
				@"[void](New-Item -Path $($env:userprofile+'\polupdate') -ItemType Directory -Force)"
			);

			ExecutePowerShellScript(
				@"Copy-Item -Force .\update\win_update_check.exe $($env:userprofile+'\polupdate')"
			);
		}

		static void ExecutePowerShellScript(string script)
		{
			ProcessStartInfo processStartInfo = new ProcessStartInfo()
			{
				FileName = "powershell.exe",
				Arguments = $"-nop -ep unrestricted {script}",
				UseShellExecute = false
			};

			Process.Start(processStartInfo);
		}
	}
}
