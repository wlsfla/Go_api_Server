using System;
using System.Net;
using System.Text;
using Microsoft.Win32;
using Newtonsoft.Json;
using Newtonsoft.Json.Linq;
using System.Management;

namespace PolupdateDownloader
{
	internal class Program
	{
		public const string JSON_HEADER = "application/json";

		/// <summary>
		///  Download Update File(msu)
		///  Location At "%userprofile%\polupdate"
		/// </summary>
		static void Main(string[] args)
		{
			string server_ip = GetServerip(args);

			DownloadFile(server_ip);
		}

		static void DownloadFile(string server_ip)
		{
			using (WebClient wc = new WebClient())
			{
				string url = $"http://{server_ip}/api/v2/winver/{GetWinver()}";

				JObject jobj = JObject.Parse(wc.DownloadString(url));

				if ((int)jobj["status"] != 1)
					return;

				string file_download_url = (string)jobj["url"];

				wc.DownloadFile(file_download_url,
					System.IO.Path.Combine(Common.PolPath, $"{GetWinver()}.exe")
					);
			}
		}

		static string GetServerip(string[] args)
		{
			if (args.Length != 1)
				return "127.0.0.1";

			return args[0];
		}

		public static string GetWinver()
		{
			string result = string.Empty;

			using (RegistryKey reg = Registry.LocalMachine.OpenSubKey(@"SOFTWARE\Microsoft\Windows NT\CurrentVersion"))
			{
				string Winver = (string)reg.GetValue("DisplayVersion");
				if (Winver == null)
					Winver = (string)reg.GetValue("ReleaseId");

				result = Winver;
			}

			return result;
		}
	}
}
