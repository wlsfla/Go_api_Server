using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Microsoft.Win32;

namespace PolupdateVersionChecker
{
	internal class VersionChkHandler
	{
		public const string JSON_HEADER = "application/json";

		/// <summary>
		/// Get Local Host PC's info
		/// </summary>
		public static Updatelog GetPCinfo()
		{
			Updatelog result;

			using (RegistryKey reg = Registry.LocalMachine.OpenSubKey(@"SOFTWARE\Microsoft\Windows NT\CurrentVersion"))
			{
				string Winver = (string)reg.GetValue("DisplayVersion");
				if (Winver == null)
					Winver = (string)reg.GetValue("ReleaseId");

				string CurrBuild = $"{(string)reg.GetValue("CurrentBuild")}.{(int)reg.GetValue("ubr")}";
				string host_name = Environment.MachineName;

				result = new Updatelog(
					String.Empty, host_name, Winver, CurrBuild, 0
					);
			}
			
			return result;
		}
	}

	// json model
	public class Updatelog
	{
		public string host_ip { get; set; }
		public string host_name { get; set; }
		public string winver { get; set; }
		public string buildver { get; set; }
		public int result { get; set; }

		public Updatelog(string host_ip, string host_name, string winver, string buildver, int result)
		{
			this.host_ip = host_ip;
			this.host_name = host_name;
			this.winver = winver;
			this.buildver = buildver;
			this.result = result;
		}
	}
}
