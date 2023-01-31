using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace PolupdateDownloader
{
	//Status:   0,
	//	Winver:   "",
	//	Buildver: "",
	//	Url:      "",


	public class Winverinfo
	{
		int status { get; set; }
		string winver { get; set; }
		string buildver { get; set; }
		string url { get; set; }

		public Winverinfo(int status, string winver, string buildver, string url)
		{
			this.status = status;
			this.winver = winver;
			this.buildver = buildver;
			this.url = url;
		}
	}
}
