<a name="zflW8"></a>
# 降级版本
低版本powershell没有AMSI,将Powershell版本降低带2.0即可绕过AMSI检测.<br />查询当前powershell版本<br />`PS: > $host`<br />`PS: > $psversiontable`<br />查询当前系统可使用的Powershell版本:<br />不需要管理员权限:<br />`Get-ChildItem 'HKLM:\SOFTWARE\Microsoft\NET Framework Setup\NDP' -recurse | Get-ItemProperty -name Version -EA 0 | Where { $_.PSChildName -match '^(?!S)\p{L}'} | Select -ExpandProperty Version`​<br />需要管理员权限:<br />`Get-WindowsOptionalFeature -Online -FeatureName MicrosoftWindowsPowerShellV2`<br />`Get-WindowsFeature PowerShell-V2 -->2016/19 `
<a name="FMDVI"></a>
# 注册表禁用
高权限下直接修改注册表相关键值即可:<br />设置注册表“HKCU\Software\Microsoft\Windows Script\Settings\AmsiEnable”设置为0,以禁用 AMSI. 
<a name="exVeD"></a>
# 劫持amsi.dll
AMSI在检测过程中会调用C:\Windows\System32\amsi.dll ,如使用csript.exe或者rundll.exe执行时,直接将其名称改为amsi.dll执行.<br />

