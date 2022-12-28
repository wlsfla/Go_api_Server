using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Net;
using System.Security.Policy;
using System.Text;
using System.Threading.Tasks;
using Microsoft.Win32;

namespace Pol_update
{
	internal class PCinfo_Manager
	{
		private RegistryKey reg;

		public PCinfo_Manager()
		{
			this.reg = Registry.LocalMachine.OpenSubKey(@"SOFTWARE\Microsoft\Windows NT\CurrentVersion");

			this.ServerAddr = "localhost:9999";
		}

		public string CurrentBuildNumber
		{
			get
			{
				return (string)this.reg.GetValue("CurrentBuildNumber");
			}
		}

		public string UBR
		{
			get
			{
				return (string)this.reg.GetValue("UBR");
			}
		}

		public string BuildVer
		{
			get
			{
				return $"{this.CurrentBuildNumber}.{this.UBR}";
			}
		}

		public string DisplayVersion
		{
			get
			{
				string dpver = (string)this.reg.GetValue("DisplayVersion");

				if (dpver == null)
					dpver = this.ReleaseId;

				return dpver;
			}
		}

		public string ReleaseId
		{
			get
			{
				return (string)this.reg.GetValue("ReleaseId");
			}
		}

		public string Hostname
		{
			get
			{
				return Environment.GetEnvironmentVariable("COMPUTERNAME");
			}
		}

		public string ServerAddr { get; set; }

		public void SendPcinfo()
		{
			string Url = $"http://{this.ServerAddr}/api/info_reg/{this.Hostname}/{this.DisplayVersion}/{this.BuildVer}";

			new WebClient().DownloadString(Url);

		}

	
	}
}
