using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Pol_update
{
	internal static class Common
	{
		/// <summary>
		/// %userprofile%\polupdate 폴더 경로
		/// </summary>
		public static string PolPath { get; set; }

		/// <summary>
		/// 현재 실행 파일 전체 경로
		/// </summary>
		public static string PolExecPath { get; set; }

		/// <summary>
		/// 현재 실행 파일 전체 경로
		/// </summary>
		public static string NowExecPath
		{
			get { return System.Windows.Forms.Application.ExecutablePath; }
		}

		/// <summary>
		/// 현재 실행 파일 폴더
		/// </summary>
		public static string NowExecDir
		{
			get { return System.Windows.Forms.Application.StartupPath; }
		}

		static Common()
		{
			PolPath = Path.Combine(Environment.GetEnvironmentVariable("userprofile"), "polupdate");
			PolExecPath = Path.Combine(PolPath, Path.GetDirectoryName(NowExecPath));

			//Path.GetFileName(Application.ExecutablePath)
		}
	}
}
