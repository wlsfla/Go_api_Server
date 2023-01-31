using System;
using System.Text;
using System.Net;
using System.IO;
using Newtonsoft.Json;

namespace PolupdateVersionChecker
{
	internal class Program
	{
		static void Main(string[] args)
		{
			Request(GetServerip(args));

			Console.WriteLine("End");
			Console.ReadKey();
		}

		static string GetServerip(string[] args)
		{
			if (args.Length != 1)
				return "127.0.0.1";

			return args[0];
		}

		static void Request(string server_ip)
		{
			string url = $"http://{server_ip}/api/v2/insert/updatelog";

			var httpWebRequest = (HttpWebRequest)WebRequest.Create(url);
			httpWebRequest.ContentType = VersionChkHandler.JSON_HEADER;
			httpWebRequest.Method = "POST";

			using (var streamWriter = new StreamWriter(httpWebRequest.GetRequestStream()))
			{
				string json = JsonConvert.SerializeObject(VersionChkHandler.GetPCinfo());
				streamWriter.Write(json);
			}

			var httpResponse = (HttpWebResponse)httpWebRequest.GetResponse();
		}
	}
}
