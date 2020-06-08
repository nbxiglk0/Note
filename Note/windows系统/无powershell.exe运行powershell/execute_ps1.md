# 无powershell.exe运行powershell工具
## C#
基本原理 : 使用C#调用system.management.automation  

powershell.exe只是.NET中[system.management.automation](https://docs.microsoft.com/en-us/dotnet/api/system.management.automation?view=pscore-6.2.0)的解释器,所以也可以使用C#直接调用`system.management.automation`来执行相应的powershell命令
### PowerLine
[使用方法](https://github.com/fullmetalcache/PowerLine)
```c#
public static void ExecuteFunc(string[] args)
        {
            string script = encodeString(args[0]);
            string command = decodeString(Functions.Funcs[script]);

            if (args.Length > 1)
            {
                Console.WriteLine("\nCommand Invoked: " + args[1]);
                string parameters = "\n" + args[1] + "\n";
                command += parameters;
            }

            //Runs powershell stuff
            MyPSHost myPSHost = new MyPSHost();
            Runspace rspace = RunspaceFactory.CreateRunspace(myPSHost);
            rspace.Open();
            Pipeline pipeline = rspace.CreatePipeline();
            pipeline.Commands.AddScript(command);
            pipeline.Commands[0].MergeMyResults(PipelineResultTypes.Error, PipelineResultTypes.Output);
            pipeline.Commands.Add("out-default");
            pipeline.InvokeAsync();
```
### nps
[使用方法](https://github.com/Ben0xA/nps)
```c#
PowerShell ps = PowerShell.Create();
                    if (args[0].ToLower() == "-encodedcommand" || args[0].ToLower() == "-enc")
                    {
                        String script = "";
                        for (int argidx = 1; argidx < args.Length; argidx++)
                        {
                            script += System.Text.Encoding.Unicode.GetString(System.Convert.FromBase64String(args[argidx]));
                        }
                        ps.AddScript(script);
                    }
                    else
                    {
                        String script = "";
                        for (int argidx = 0; argidx < args.Length; argidx++)
                        {
                            script += @args[argidx];
                        }
                        ps.AddScript(script);
                    }

                    Collection<PSObject> output = null;
                    try
                    {
                        output = ps.Invoke();
                    }
```
### Powershdll
[使用方法](https://github.com/p3nt4/PowerShdll)
该工具可以使用dll 或者 exe来执行powershell
```C#
        public string exe(string cmd)
        {
            try
            {
                Pipeline pipeline = runspace.CreatePipeline();
                pipeline.Commands.AddScript(cmd);
                pipeline.Commands.Add("Out-String");
                Collection<PSObject> results = pipeline.Invoke();
                StringBuilder stringBuilder = new StringBuilder();
                foreach (PSObject obj in results)
                {
                    stringBuilder.AppendLine(obj.ToString());
                }
                return stringBuilder.ToString();
            }
            catch (Exception e)
            {
                // Let the user know what went wrong.

                string errorText = e.Message + "\n";
                return (errorText);
            }
        }
        public void close()
        {
            this.runspace.Close();
        }
    }
```
### Nopowershell
[使用方法](https://github.com/bitsadmin/nopowershell)