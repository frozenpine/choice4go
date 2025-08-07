#ifndef CGO_DYN_LIB_H
#define CGO_DYN_LIB_H

#include <stdlib.h>
#include <stdbool.h>
#include <dlfcn.h>

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

//设置主回调函数(一定要设置一个主回调函数，可在调用start之前调用，否则收不到账号掉线通知)
const char* CALLBACK_SETTER_NAME = "setcallback";

//用户可以自定义"ServerList.json.e"和"userInfo"文件的存放目录(如果不调用此函数或者dir传空则默认当前目录)
const char* SET_SERVER_LIST_NAME = "setserverlistdir";

//获取错误码文本说明（返回指针不要释放）
const char* ERR_GETTER_NAME = "geterrstring";

//证券与指标校验函数，获取相匹配的csd/css/cses的证券和指标请求参数，并按证券品种区分 options必须传入FunType=CSD 或CSS 或 CSES之一
const char* CFC_VERIFIER_NAME = "cfc";

//校验或补全东财证券代码函数 options传入ReturnType=0/1  0:返回并标记代码是否正确 1:根据SecuType与SecuMarket补全代码后缀(有可能返回多个不同的后缀)
const char* CEC_VERIFIER_NAME = "cec";

//设置网络代理 注：如需使用代理，需要在调用所有接口之前设置
const char* SET_PROXY_NAME = "setproxy";

/**登录(开始时或掉线后调用) 
*  参数说明：
*  pLoginInfo：保留参数,无需传(2.0.0.0版本之后改为令牌自动登陆,保留此参数以兼容旧版本)  
*  options：附加参数,用半角逗号隔开   现开放 TestLatency=1（服务器测速,默认为0不测速） 
*           ForceLogin=1 （强制登录，默认为0普通登录） LogLevel=2(日志级别 1:Debug 2:Info 3:Error)
*  pfnCallback：日志回调函数*/
const char* STARTER_NAME = "start";

//退出(结束退出时调用，只需调用一次)
const char* STOPPER_NAME = "stop";

//指标服务数据查询(同步请求)
const char* CSD_QUERIER_NAME = "csd";

//截面数据查询(同步请求)
const char* CSS_QUERIER_NAME = "css";

//板块截面数据查询(同步请求)
const char* CSES_QUERIER_NAME = "cses";

//获取区间日期内的交易日(同步请求)
const char* TRADEDATE_QUERIER_NAME = "tradedates";

//获取系统板块成分(同步请求)
const char* SECTOR_QUERIER_NAME = "sector";

//获取专题报表(同步请求)
const char* CTR_QUERIER_NAME = "ctr";

//仅供本API中同步接口返回数据指针释放内存(EQDATA* 或 EQCTRDATA* 或者 EQCHAR*, 不可传入其他指针，异步函数回调中的指针也不可传入)
const char* DATA_RELEASER_NAME = "releasedata";

//宏观指标服务(同步请求)
const char* EDB_QUERIER_NAME = "edb";

//宏观指标id详情查询(同步请求)
const char* DEB_DTL_QUERIER_NAME = "edbquery";

//资讯数据查询(同步请求) codes：东财代码或板块代码（不可混合） content：查询内容
const char* CFN_QUERIER_NAME = "cfn";

//板块树查询（同步请求）
const char* CFN_DTL_QUERIER_NAME = "cfnquery";

typedef EQErr (*callback_setter)(datacallback);
typedef const char* (*err_getter)(EQErr, EQLang);
typedef EQErr (*starter)(EQLOGININFO*, const char*, logcallback);
typedef EQErr (*stopper)();
typedef EQErr (*data_releaser)(void*);
typedef EQErr (*query_pchar2_pdata)(const char*, const char*, EQDATA**);
typedef EQErr (*query_pchar3_pdata)(const char*, const char*, const char*, EQDATA**);
typedef EQErr (*query_pchar5_pdata)(const char*, const char*, const char*, const char*, const char*, EQDATA**);

int CallCbSetter(callback_setter fn, datacallback cb)
{
	return fn(cb);
}

const char* CallErrGetter(err_getter fn, EQErr code, EQLang lang)
{
	return fn(code, lang);
}

int CallStarter(starter fn, EQLOGININFO* login, const char* options, logcallback cb)
{
	return fn(login, options, cb);
}

int CallStopper(stopper fn) { return fn(); }

int CallDataReleaser(data_releaser fn, void* data) { return fn(data); }

int CallPChar2PData(
	query_pchar2_pdata fn, const char* p1, const char* p2, EQDATA** data
)
{
	return fn(p1, p2, data);
}

int CallPChar3PData(
	query_pchar3_pdata fn, const char* p1, const char* p2,
	const char* p3, EQDATA** data
)
{
	return fn(p1, p2, p3, data);
}

int CallPChar5PData(
	query_pchar5_pdata fn, const char* p1, const char* p2,
	const char* p3, const char* p4, const char* p5, EQDATA** data
)
{
	return fn(p1, p2, p3, p4, p5, data);
}

#ifdef __cplusplus
}
#endif

#endif