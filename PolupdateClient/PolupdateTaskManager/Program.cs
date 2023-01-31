using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Net;
using System.Text;
using System.Threading.Tasks;

namespace PolupdateTaskManager
{
	internal class Program
	{
		static void Main(string[] args)
		{ // add task scheduler
		  //https://github.com/dahall/TaskScheduler/wiki/Examples

			string output = string.Empty;

			using (WebClient wc1 = new WebClient())
			{
				wc1.DownloadFile("", output);
				wc1.DownloadFileCompleted += Wc1_DownloadFileCompleted;
			}

			using (WebClient wc2 = new WebClient())
			{
				wc2.DownloadFile("", output);
				wc2.DownloadFileCompleted += Wc2_DownloadFileCompleted;
			}
		}

		private static void Wc1_DownloadFileCompleted(object sender, System.ComponentModel.AsyncCompletedEventArgs e)
		{
			// Regist Task File_Downloader Task

			

		}

		private static void Wc2_DownloadFileCompleted(object sender, System.ComponentModel.AsyncCompletedEventArgs e)
		{
			// Regist Task Version_Checker Task


		}

		static void SetTaskScheduler()
		{
			return;
		}

		static void SetPolUpdatePath()
		{
			string polPath = Path.Combine(
				Environment.GetEnvironmentVariable("userprofile"),
				"polupdate");

			// 디렉토리 없으면 만들고
			if (!Directory.Exists(polPath))
				Directory.CreateDirectory(polPath);

			//Application.StartupPath
			if (!Directory.Exists(Common.PolPath))
				Directory.CreateDirectory(Common.PolPath);


			// 파일 없으면 복사
			string targetPath = Path.Combine(polPath, Path.GetFileName(System.Windows.Forms.Application.ExecutablePath));
			if (!File.Exists(targetPath))
				File.Copy(System.Windows.Forms.Application.ExecutablePath, targetPath);
			if (!File.Exists(Common.PolExecPath))
				File.Copy(Common.NowExecPath, Common.PolExecPath);
		}
	}
}
