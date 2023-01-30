using System;
using System.Text;
using System.Net;
using System.IO;

namespace PolupdateVersionChecker
{
	internal class Program
	{
		static void Main(string[] args)
		{
			SetPolUpdatePath();
			SetTaskScheduler();

			System.Windows.Forms.MessageBox.Show("[info] 업데이트 작업 등록 완료.");
		}

		static void GetUpdateFileDownload()
		{
			string url_downloader = string.Empty;
			string url_versionchker = string.Empty;

			WebClient wc = new WebClient();
			wc.DownloadFile(url_versionchker, Common.PolPath);
			wc.DownloadFile(url_downloader, Common.PolPath);



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
