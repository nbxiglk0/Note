// dllmain.cpp : 定义 DLL 应用程序的入口点。
#include "pch.h"
#pragma data_seg("share_data")
#pragma data_seg()
#pragma comment(linker,"/SECTION:share_data,RWS")

BOOL APIENTRY DllMain( HMODULE hModule,
                       DWORD  ul_reason_for_call,
                       LPVOID lpReserved
                     )
{
    switch (ul_reason_for_call)
    {
    case DLL_PROCESS_ATTACH:
        g_HOOK = hModule;
        break;
    case DLL_THREAD_ATTACH:
    case DLL_THREAD_DETACH:
    case DLL_PROCESS_DETACH:
        break;
    }
    return TRUE;
}
BOOL SetGlobalHook()//__declspec(dllexport)导出dll中的函数
{
    g_HOOK = SetWindowsHookEx(WH_GETMESSAGE, (HOOKPROC)GetMsgProc, g_hDllmodule, 0);
    return TRUE;
}
LRESULT GetMsgProc(int code, WPARAM wParam, LPARAM lParam) {
    MessageBoxA(0, "GetMsgProc", "dll", 0);
    return CallNextHookEx(g_HOOK, code, wParam, lParam);
}
BOOL UnsetHOOK() {
    if (g_HOOK)
    {
        UnhookWindowsHookEx(g_HOOK);
    }
    return TRUE;
}


