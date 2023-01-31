using System;
using System.IO;

namespace PolupdateDownloader
{
	internal class Common
	{
		private static string polpath;

		/// <summary>
		/// Directory : %userprofile%\polupdate\
		/// </summary>
		public static string PolPath
		{
			get
			{
				// Set PolPath
				if (!Directory.Exists(polpath))
					Directory.CreateDirectory(polpath);

				return polpath;
			}
		}

		/// <summary>
		/// %userprofile%\polupdate\PolupdateDownloader_202212.exe
		/// </summary>
		public static string PolExecPath { get; }

		/// <summary>
		/// 현재 실행 파일 전체 경로
		/// </summary>
		public static string CurrExecPath { get; }

		/// <summary>
		/// 현재 실행 파일 폴더
		/// </summary>
		public static string CurrExecDir { get; }

		static Common()
		{
			polpath = Path.Combine(Environment.GetEnvironmentVariable("userprofile"), "polupdate");

			string versionDir = "202212"
			string filename = System.Windows.Forms.Application.ProductName + "_202212" + ".exe"; // PolupdateDownloader_202212.exe
			PolExecPath = Path.Combine(polpath, filename);
			CurrExecPath = System.Windows.Forms.Application.ExecutablePath;
			CurrExecDir = System.Windows.Forms.Application.StartupPath;
		}
	}
}
