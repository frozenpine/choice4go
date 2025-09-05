#pragma once
#ifndef CGO_CHOICE_BRIDGE_H
#define CGO_CHOICE_BRIDGE_H

#include <stdlib.h>
#include <stdbool.h>

#ifndef WIN32
#define WIN32
#endif

typedef void* LPVOID;
#define EMQUANTAPI

#ifdef __cplusplus
extern "C"
{
#else
#define NULL ((void *)0)
#endif

#include "EmQuantAPI.h"

int cLogCallback(const char* pLog);
int cDataCallback(const EQMSG* pMsg, LPVOID lpUserParam);

#ifdef __cplusplus
}
#endif

#endif