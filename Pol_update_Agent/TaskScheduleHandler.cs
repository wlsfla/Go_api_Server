using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Microsoft.Win32.TaskScheduler;

namespace Pol_update
{
	internal static class TaskScheduleHandler
	{
        public static void test(string taskPath, string execPath, string args)
		{
			using (TaskService taskService = new TaskService())
			{
				Microsoft.Win32.TaskScheduler.Task task = taskService.GetTask(taskPath); // 현재 task name 객체 검색

				if (task == null)
				{
					TaskDefinition taskDefinition = taskService.NewTask();

					taskDefinition.Principal.DisplayName = "메모장 실행";
					taskDefinition.Principal.LogonType = TaskLogonType.InteractiveToken;
					taskDefinition.Principal.UserId = $"{Environment.UserDomainName}\\{Environment.UserName}";
					taskDefinition.Principal.RunLevel = TaskRunLevel.Highest;

					taskDefinition.Settings.MultipleInstances = TaskInstancesPolicy.IgnoreNew;
					taskDefinition.Settings.DisallowStartIfOnBatteries = false;
					taskDefinition.Settings.StopIfGoingOnBatteries = false;
					taskDefinition.Settings.AllowHardTerminate = false;
					taskDefinition.Settings.StartWhenAvailable = false;
					taskDefinition.Settings.RunOnlyIfNetworkAvailable = false;
					taskDefinition.Settings.IdleSettings.StopOnIdleEnd = false;
					taskDefinition.Settings.IdleSettings.RestartOnIdle = false;
					taskDefinition.Settings.AllowDemandStart = false;
					taskDefinition.Settings.RunOnlyIfIdle = false;
					taskDefinition.Settings.ExecutionTimeLimit = TimeSpan.Zero;
					taskDefinition.Settings.Priority = ProcessPriorityClass.High;
					taskDefinition.Settings.Hidden = false;
					taskDefinition.Settings.Enabled = true;

					DailyTrigger trigger = new DailyTrigger();
					trigger.RandomDelay = TimeSpan.FromHours(6); // c# TimeSpan. 6 Hours.
					trigger.DaysInterval = 1;
					trigger.StartBoundary = DateTime.Now;
					trigger.EndBoundary = new DateTime(2023, 1, 6);

					taskDefinition.Triggers.Add(trigger);

					taskDefinition.Actions.Add(new ExecAction("notepad.exe", "test")); // file name.

					taskService.RootFolder.RegisterTaskDefinition(taskPath, taskDefinition);
				}
			}
		}

	}
}
