using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Net;
using System.Net.Http.Headers;
using System.Security.Policy;
using System.Text;
using System.Threading.Tasks;
using Microsoft.Win32;

namespace Pol_update
{
	internal class PCinfo_Manager
	{
		private RegistryKey reg;

		public PCinfo_Manager() : this("localhost:9999")
		{
			
		}

		public PCinfo_Manager(string server_addr)
		{
			this.reg = Registry.LocalMachine.OpenSubKey(@"SOFTWARE\Microsoft\Windows NT\CurrentVersion");
			this.ServerAddr = server_addr;
			this.winverList = new Dictionary<string, double>()
			{
				{ "1803", 17134.2208},
				{ "1809", 17763.3653},
				{ "1903", 18362.1256},
				{ "1909", 18363.2274},
				{ "20H2", 19042.2311},
				{ "21H1", 19043.2311},
				{ "21H2", 19044.2311}
			};
		}

		private Dictionary<string, double> winverList { get; set; }

		private string CurrentBuildNumber
		{
			get
			{
				return Convert.ToString(reg.GetValue("CurrentBuildNumber"));
			}
		}

		private string UBR
		{
			get
			{
				return Convert.ToString(Convert.ToDouble(this.reg.GetValue("UBR")));
			}
		}

		private double CurrentBuildVer
		{
			get
			{
				return Convert.ToDouble($"{this.CurrentBuildNumber}.{this.UBR}");
			}
		}

		private string DisplayVersion
		{
			get
			{
				string dpver = (string)this.reg.GetValue("DisplayVersion");

				if (dpver == null)
					dpver = this.ReleaseId;

				return dpver;
			}
		}

		private string ReleaseId
		{
			get
			{
				return (string)this.reg.GetValue("ReleaseId");
			}
		}

		private string Hostname
		{
			get
			{
				return Environment.GetEnvironmentVariable("COMPUTERNAME");
			}
		}

		public string ServerAddr { get; set; }

		private bool IsContainsDpVer
		{
			get { return this.winverList.ContainsKey(this.DisplayVersion); }
		}

		public void ChkBuildVer()
		{
			if ((this.DisplayVersion == "1803") || !IsContainsDpVer || (this.CurrentBuildVer >= this.winverList[this.DisplayVersion])) // 1803 버전이거나 현재 버전과 다르면
			{
				Console.WriteLine("[info] Cannot Update.");
				new WebClient().DownloadString($"http://{this.ServerAddr}/api/result/2");
				return;
			}

			if (this.CurrentBuildVer == this.winverList[this.DisplayVersion])
			{
				Console.WriteLine("[info] Cannot Update.");
				new WebClient().DownloadString($"http://{this.ServerAddr}/api/result/1");
				return;
			}

			Console.WriteLine($"[info] Access This URL. And Download File In Internet Explorer \n\n >>> http://{this.ServerAddr}/file/{this.DisplayVersion}\n");
		}

		public void SendPcinfo()
		{
			string Url = $"http://{this.ServerAddr}/api/info_reg/{this.Hostname}/{this.DisplayVersion}/{this.CurrentBuildVer}";
			new WebClient().DownloadString(Url);
		}

	
	}
}
