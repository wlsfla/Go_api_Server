using System;
using System.Collections.Generic;
using System.Linq;
using System.Net;
using System.Text;
using System.Threading.Tasks;

namespace PolupdateDownloader
{
	internal class Program
	{
		static void Main(string[] args)
		{
			if (args.Length > 1 || args.Length == 0)
				return;

			string server_ip = args[0];
			string output = string.Empty;

			WebClient wc = new WebClient();
			wc.DownloadFile(server_ip, output);
			wc.DownloadFileCompleted += Wc_DownloadFileCompleted;

			Console.ReadKey();
		}

		private static void Wc_DownloadFileCompleted(object sender, System.ComponentModel.AsyncCompletedEventArgs e)
		{
			
			
		}
	}
}
