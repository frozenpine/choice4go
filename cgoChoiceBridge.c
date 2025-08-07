#include "cgoChoiceBridge.h"

extern int cgoLogCallback(const char* pLog);
extern int cgoDataCallback(const EQMSG* pMsg, LPVOID lpUserParam);


int cLogCallback(const char* pLog) 
{
    return cgoLogCallback(pLog);
}

int cDataCallback(const EQMSG* pMsg, LPVOID lpUserParam) 
{
    return cgoDataCallback(pMsg, lpUserParam);
}
