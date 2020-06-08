# include <stdio.h>
# include <windows.h>
# include "pch.h"
HMODULE hint;
typedef BOOL(*HILB)();
void main() {
	HMODULE hint = LoadLibrary("inject.dll");
	if (NULL == hint)
	{	
		MessageBoxA(0, "Loaddll Fail!", "load", 1);
		return;
	}
	HILB fundll = NULL;
	printf("%p", hint);
	fundll = (HILB)(GetProcAddress(hint, "SetGlobalHOOK"));
	if (NULL != fundll) {
		(*fundll)();
	}
}