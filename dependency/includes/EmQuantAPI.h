/**************************************************************************************************
*   EmQuantAPI.h    version  2.6.0.0                                                              *                                                                      *
*   Copyright(c)2016-2025,  Shanghai Choice Information Technology Co,. Ltd. All Rights Reserved. *
*   Eastmoney API must not be used without authorzation, unauthorized user shall be held liable.  *
***************************************************************************************************/
#ifndef _EMQUANTAPI_H_
#define _EMQUANTAPI_H_ 

#include <stdint.h>

#define EMQUNANTAPIVER 2600

// #ifdef WIN32
// #ifdef EMQUANTAPI_EXPORTS
// #define EMQUANTAPI extern "C" __declspec(dllexport)
// #else
// #define EMQUANTAPI extern "C" __declspec(dllimport)
// #endif /*#ifdef EMQUANTAPI_EXPORTS*/
// #else
// #define EMQUANTAPI extern "C"
// #endif /*#ifdef WIN32*/

//错误码的类型定义
typedef int EQErr;

//请求ID的类型定义
typedef int EQID;

#ifndef WIN32
typedef void* LPVOID;
#endif //#ifndef WIN32
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//成功
#define EQERR_SUCCESS                  (0)
//错误基数                             
#define EQERR_BASE                     (10000000)
//一般性错误                           
#define EQERR_BASE_GENERAL             (EQERR_BASE)
//账户相关错误						 
#define EQERR_BASE_ACCOUT              (EQERR_BASE + 1000)
//网络相关错误                         
#define EQERR_BASE_NET                 (EQERR_BASE + 2000)
//参数或请求错误                       
#define EQERR_BASE_PARAM               (EQERR_BASE + 3000)


#pragma region 账户相关错误 
//用户未登录
#define EQERR_NO_LOGIN                 (EQERR_BASE_ACCOUT + 1)
//用户名或密码错误					   
#define EQERR_USERNAMEORPASSWORD_ERR   (EQERR_BASE_ACCOUT + 2)
//用户无API权限						  
#define EQERR_NO_ACCESS                (EQERR_BASE_ACCOUT + 3)
//用户API权限过期						 
#define EQERR_ACCESS_EXPIRE            (EQERR_BASE_ACCOUT + 4)
//获取用户信息失败					   
#define EQERR_GETUSERINFO_FAIL         (EQERR_BASE_ACCOUT + 5)
//DLL版本号过期						  
#define EQERR_DLLVESION_EXPIRE         (EQERR_BASE_ACCOUT + 6)
//用户无API_LV2权限					  
#define EQERR_NO_LV2_ACCESS            (EQERR_BASE_ACCOUT + 7)
//用户API_LV2权限过期					 
#define EQERR_LV2_ACCESS_EXPIRE        (EQERR_BASE_ACCOUT + 8)
//账号登录数达到上限					  
#define EQERR_LOGIN_COUNT_LIMIT        (EQERR_BASE_ACCOUT + 9)
//用户登录失败						 
#define EQERR_LOGIN_FAIL               (EQERR_BASE_ACCOUT + 10)
//用户登录掉线						 
#define EQERR_LOGIN_DISCONNECT         (EQERR_BASE_ACCOUT + 11)
//用户权限不足						 
#define EQERR_ACCESS_INSUFFICIENCE     (EQERR_BASE_ACCOUT + 12)
//用户正在登录						 
#define EQERR_IS_LOGIN                 (EQERR_BASE_ACCOUT + 13)
//需要登录激活						 
#define EQERR_NEED_ACTIVATE            (EQERR_BASE_ACCOUT + 14)
//登录服务异常						 
#define EQERR_LOGIN_SERVICE_ERR        (EQERR_BASE_ACCOUT + 15)
//正在人工激活						 
#define EQERR_IS_MANUAL_ACTIVATE       (EQERR_BASE_ACCOUT + 16)
//无需人工激活						 
#define EQERR_NOTNEED_MANUAL_ACTIVATE  (EQERR_BASE_ACCOUT + 17)
//人工激活失败						 
#define EQERR_MANUAL_ACTIVATE_FAIL     (EQERR_BASE_ACCOUT + 18)
//激活设备与登录设备不一致			   
#define EQERR_DIFFRENT_DEVICE          (EQERR_BASE_ACCOUT + 19)
//userInfo已失效需重新激活			   
#define EQERR_USERINFO_EXPIRED         (EQERR_BASE_ACCOUT + 20)
//行情服务登录验证失败
#define EQERR_QUOTE_LOGIN_FAIL         (EQERR_BASE_ACCOUT + 21)
//行情服务流量验证失败
#define EQERR_QUOTE_FLOW_FAIL          (EQERR_BASE_ACCOUT + 22)
//资讯查询服务登录验证失败
#define EQERR_INFOQUERY_LOGIN_FAIL     (EQERR_BASE_ACCOUT + 23)
//资讯订阅服务登录验证失败
#define EQERR_INFOSUB_LOGIN_FAIL       (EQERR_BASE_ACCOUT + 24)
//资讯服务流量验证失败
#define EQERR_INFO_FLOW_FAIL           (EQERR_BASE_ACCOUT + 25)
//无效的上行短信
#define EQERR_SMS_INVALIED             (EQERR_BASE_ACCOUT + 26)
//专项服务登录验证失败
#define EQERR_CHQQUOTE_LOGIN_FAIL      (EQERR_BASE_ACCOUT + 27)
//专项服务权限验证失败
#define EQERR_CHQQUOTE_ACCESS_FAIL     (EQERR_BASE_ACCOUT + 28)
#pragma  endregion    

#pragma region 一般性错误             
//获取交易日失败                      
#define EQERR_GET_TRADE_FAIL           (EQERR_BASE_GENERAL + 1)
//初始化主类失败                	    
#define EQERR_INIT_OBTAIN_CLASS_FAIL   (EQERR_BASE_GENERAL + 2)
//申请内存失败                         
#define EQERR_NEW_MEM_FAIL             (EQERR_BASE_GENERAL + 3)
//解析数据错误                         
#define EQERR_PARSE_DATA_ERR           (EQERR_BASE_GENERAL + 4)
//gzip解压失败                         
#define EQERR_UNGZIP_DATA_FAIL         (EQERR_BASE_GENERAL + 5)
//未知错误							   
#define EQERR_UNKNOWN_ERR              (EQERR_BASE_GENERAL + 6)
//函数内部错误						 
#define EQERR_FUNCTION_INTERNAL_ERR    (EQERR_BASE_GENERAL + 7)
//数组越界							   
#define EQERR_OUTOF_BOUNDS             (EQERR_BASE_GENERAL + 8)
//无数据							    
#define EQERR_NO_DATA                  (EQERR_BASE_GENERAL + 9)
//系统级别错误						 
#define EQERR_SYSTEM_ERROR             (EQERR_BASE_GENERAL + 10)
//服务器列表错误                          
#define EQERR_SERVERLIST_ERROR         (EQERR_BASE_GENERAL + 11)
//操作失败							   
#define EQERR_OPERATION_FAILURE        (EQERR_BASE_GENERAL + 12)
//服务出错							   
#define EQERR_SERVICE_ERROR            (EQERR_BASE_GENERAL + 13)
//获取服务器列表失败
#define EQERR_GETSERVERLIST_FAIL       (EQERR_BASE_GENERAL + 14)
//服务超时
#define EQERR_SERVICE_TIMEOUT          (EQERR_BASE_GENERAL + 15)
//请求频次过高
#define EQERR_FREQUENCY_OVER           (EQERR_BASE_GENERAL + 16)
//海外IP受限
#define EQERR_OVERSEAS_IP_RESTRICTED   (EQERR_BASE_GENERAL + 17)
//POP组合不支持此操作
#define EQERR_POP_GROUP_NOT_SUPPORT    (EQERR_BASE_GENERAL + 18)
#pragma  endregion                    

#pragma region 网络相关错误           
//网络错误                            
#define EQERR_SOCKET_ERR               (EQERR_BASE_NET + 1)
//网络连接失败                         
#define EQERR_CONNECT_FAIL             (EQERR_BASE_NET + 2)
//网络连接超时                         
#define EQERR_CONNECT_TIMEOUT          (EQERR_BASE_NET + 3)
//网络接收时连接断开					  
#define EQERR_RECVCONNECTION_CLOSED    (EQERR_BASE_NET + 4)
//网络发送失败                         
#define EQERR_SENDSOCK_FAIL            (EQERR_BASE_NET + 5)
//网络发送超时                         
#define EQERR_SENDSOCK_TIMEOUT         (EQERR_BASE_NET + 6)
//网络接收错误                         
#define EQERR_RECVSOCK_FAIL            (EQERR_BASE_NET + 7)
//网络接收超时                         
#define EQERR_RECVSOCK_TIMEOUT         (EQERR_BASE_NET + 8)
//行情服务器连续重连失败
#define EQERR_QUOTE_RECONNECT_FAIL     (EQERR_BASE_NET + 9)
//http访问失败
#define EQERR_HTTP_FAIL                (EQERR_BASE_NET + 10)
//等待网络响应超时 
#define EQERR_WAIT_NET_RES_TIMEOUT     (EQERR_BASE_NET + 11)
//行情服务器重连
#define EQERR_QUOTE_RECONNECT          (EQERR_BASE_NET + 12)
//资讯服务器重连
#define EQERR_INFO_RECONNECT           (EQERR_BASE_NET + 13)
//资讯服务器连续重连失败
#define EQERR_INFO_RECONNECT_FAIL      (EQERR_BASE_NET + 14)
//专项服务器连续重连失败
#define EQERR_CHQQUOTE_RECONNECT_FAIL  (EQERR_BASE_NET + 15)
//专项服务器重连
#define EQERR_CHQQUOTE_RECONNECT       (EQERR_BASE_NET + 16)
#pragma  endregion

#pragma region 参数或请求错误
//传入参数为空
#define EQERR_INPARAM_EMPTY            (EQERR_BASE_PARAM + 1)
//传出参数为空					       
#define EQERR_OUTPARAM_EMPTY           (EQERR_BASE_PARAM + 2)
//参数错误						       
#define EQERR_PARAM_ERR                (EQERR_BASE_PARAM + 3)
//起始日期格式不正确				      
#define EQERR_START_DATE_ERR           (EQERR_BASE_PARAM + 4)
//截止日期格式不正确  			       
#define EQERR_END_DATE_ERR             (EQERR_BASE_PARAM + 5)
//起始日期大于截至日期			       
#define EQERR_START_BIGTHAN_END        (EQERR_BASE_PARAM + 6)
//日期格式不正确  				        
#define EQERR_DATE_ERR                 (EQERR_BASE_PARAM + 7)
//无效的证券代码				        
#define EQERR_CODE_INVALIED            (EQERR_BASE_PARAM + 8)
//证券代码重复					       
#define EQERR_CODE_REPEAT              (EQERR_BASE_PARAM + 9)
//无效的指标						      
#define EQERR_INDICATOR_INVALIED       (EQERR_BASE_PARAM + 10)
//用户名为空						      
#define EQERR_USERNAME_EMPTY           (EQERR_BASE_PARAM + 11)
//密码为空						       
#define EQERR_PASSWORD_EMPTY           (EQERR_BASE_PARAM + 12)
//订阅数或股票总数达到上限		       
#define EQERR_TO_UPPER_LIMIT           (EQERR_BASE_PARAM + 13)
//不支持的混合指标				       
#define EQERR_MIXED_INDICATOR          (EQERR_BASE_PARAM + 14)
//单次订阅指标达到上限				 
#define EQERR_INDICATOR_TO_UPPER_LIMIT (EQERR_BASE_PARAM + 15)
//超出日期支持范围					   
#define EQERR_BEYOND_DATE_SUPPORT      (EQERR_BASE_PARAM + 16)
//复权基期小于截止日期			       
#define EQERR_BASE_LESS_THAN_END       (EQERR_BASE_PARAM + 17)
//不支持的混合证券品种			       
#define EQERR_MIXED_CODES_MARKET       (EQERR_BASE_PARAM + 18)
//不支持的证券代码品种				 
#define EQERR_NO_SUPPORT_CODES_MARKET  (EQERR_BASE_PARAM + 19)
//交易条数超过上限					   
#define EQERR_ORDER_TO_UPPER_LIMIT     (EQERR_BASE_PARAM + 20)
//不支持的交易信息				       
#define EQERR_NO_SUPPORT_ORDERINFO     (EQERR_BASE_PARAM + 21)
//指标重复						 
#define EQERR_INDICATOR_REPEAT         (EQERR_BASE_PARAM + 22)
//资讯板块代码错误					 
#define EQERR_INFOBKCODE_INVALIED      (EQERR_BASE_PARAM + 23)
//资讯数据量过大
#define EQERR_INFOSIZE_TOOLARGE        (EQERR_BASE_PARAM + 24)
//资讯查询不到数据
#define EQERR_INFO_SEARCH_NODATA       (EQERR_BASE_PARAM + 25)
//资讯板块代码重复
#define EQERR_INFOBKCODE_REPEAT		   (EQERR_BASE_PARAM + 26)

#pragma  endregion

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//错误信息语言类型枚举
typedef enum
{
    eLang_ch,
    eLang_en,
} EQLang;

//EQVARTYPE数据类型枚举
typedef enum 
{
    eVT_null,
    eVT_char,
    eVT_byte,
    eVT_bool,
    eVT_short,
    eVT_ushort,
    eVT_int,
    eVT_uInt,
    eVT_int64,
    eVT_uInt64,
    eVT_float,
    eVT_double,
    eVT_byteArray,
    eVT_asciiString,
    eVT_unicodeString
} EQVarType;

// 返回Events结构类型枚举
typedef enum 
{
    eMT_err,                             // 出错信息
    eMT_response,                        // 数据返回
    eMT_partialResponse,                 // 部分数据返回（订阅时适用）
    eMT_others,                          // 其他信息
} EQMsgType;

//网络代理方式枚举
typedef enum 
{
	ePT_NONE,                            // 不使用代理  
	ePT_HTTP,						     // HTTP代理
	ePT_HTTPS,						     // HTTPS代理
	ePT_SOCK4,						     // SOCK4代理
	ePT_SOCK5                            // SOCK5代理
} ProxyType;

//下单操作类型枚举
typedef enum 
{
	eOT_default,                         // 默认(默认则根据传入数量的正负标志买入eOT_buy卖出eOT_Sell,其余类型对数量作正负转换)
	eOT_buy,						     // 买入
	eOT_sell,						     // 卖出
	eOT_purchase,						 // 申购
	eOT_redemption,                      // 赎回
} OperateType;

//资讯查询类型枚举
typedef enum
{
	eCfnMode_StartToEnd = 1,             //starttime和endtime中间的所有资讯
	eCfnMode_EndCount,                   //提取endtime的近count条数据
}eCfnMode;

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////

#pragma pack(push)
#pragma pack(8)

//char数组封装结构体
typedef struct _tagEQCHAR
{
    char*        pChar;                            //字符指针数组
    unsigned int nSize;                            //字符数组长度，(存储字符串时包括0结尾长度）
} EQCHAR;

//char二维数组封装结构体
typedef struct _tagEQCHARARRAY
{
    EQCHAR*      pChArray;                         //EQCHAR指针数组
    unsigned int nSize;                            //EQCHAR数组长度
} EQCHARARRAY;

//EQVARIENT数据类型封装结构体，可存储多种数据类型
typedef struct _tagEQVARIENT
{
    EQVarType          vtype;                      //值类型
    union 
    {
        char           charValue;
        bool           boolValue;
        short          shortValue;
        unsigned short uShortValue;
        int            intValue;
        unsigned int   uIntValue;
        int64_t        int64Value;
        uint64_t       uInt64Value;
        float          floatValue;
        double         doubleValue;
    } unionValues;                                 //联合体
    EQCHAR eqchar;                                 //字符数组
} EQVARIENT;

//EQVARIENT的数组封装结构体
typedef struct _tagEQVARIENTARRAY
{
    EQVARIENT*     pEQVarient;                       //EQVARIENT指针数组
    unsigned int   nSize;                            //EQVARIENT数组长度
} EQVARIENTARRAY;

//EQDATA数据返回结构体
typedef struct _tagEQDATA
{
    EQCHARARRAY    codeArray;                        //股票代码数组
    EQCHARARRAY    indicatorArray;                   //指标数组
    EQCHARARRAY    dateArray;                        //日期数组
    EQVARIENTARRAY valueArray;                       //值数组

    // //operator ()(int,int,int)
    // EQVARIENT* operator()( const unsigned int code, const unsigned int indicator,const unsigned int date )
    // {
    //     EQVARIENT* pEQVar = NULL;
    //     if(valueArray.nSize == 0) 
    //     {
    //         return pEQVar;
    //     }

    //     if(valueArray.nSize != codeArray.nSize * dateArray.nSize * indicatorArray.nSize)
    //     {
    //         return pEQVar;
    //     }

    //     if(code < codeArray.nSize && indicator < indicatorArray.nSize && date < dateArray.nSize)
    //     {
    //         int nIndex = codeArray.nSize * indicatorArray.nSize * date + indicatorArray.nSize * code + indicator;
    //         pEQVar = valueArray.pEQVarient;
    //         pEQVar += nIndex;
    //     }

    //     return pEQVar;
    // }
    // //EQVARIENT*
    // EQVARIENT* GetAt(const unsigned int pos)
    // {
    //     if(pos >= valueArray.nSize)
    //     {
    //         return NULL;
    //     }
    //     return &(valueArray.pEQVarient[pos]);
    // }

} EQDATA, *PEQDATA;

//EQMSG数据结构体(异步返回回调时使用)
typedef struct _tagEQMSG
{
    int            version;						     //版本号
    EQMsgType      msgType;                          //Msg类型
    EQErr          err;							     //错误码
    EQID           requestID;					     //请求ID
    EQID           serialID;					     //流水号
    EQDATA*        pEQData;      		             //包含的数据
} EQMSG, *PEQMSG;

//报表输出结构体
typedef struct _tagEQCtrData
{
	int            row;                              //行数
	int            column;                           //列数
	EQCHARARRAY    indicatorArray;                   //指标名数组
	EQVARIENTARRAY valueArray;                       //值数组

	// //operator ()(int,int)
	// EQVARIENT* operator()( const unsigned int rowIndex, const unsigned int columnIndex)
	// {
	// 	EQVARIENT* pEQVar = NULL;
	// 	if(valueArray.nSize == 0) 
	// 	{
	// 		return pEQVar;
	// 	}

	// 	if(valueArray.nSize != row * column)
	// 	{
	// 		return pEQVar;
	// 	}

	// 	if(rowIndex < row && columnIndex < column)
	// 	{
	// 		int nIndex = column * rowIndex + columnIndex;
	// 		pEQVar = valueArray.pEQVarient;
	// 		pEQVar += nIndex;
	// 	}

	// 	return pEQVar;
	// }
	// //EQVARIENT*
	// EQVARIENT* GetAt(const unsigned int pos)
	// {
	// 	if(pos >= valueArray.nSize)
	// 	{
	// 		return NULL;
	// 	}
	// 	return &(valueArray.pEQVarient[pos]);
	// }

} EQCTRDATA;

// 登录信息
#define MAXSTRING 255
typedef struct _tagEQLOGININFO
{
	char userName[MAXSTRING];                        //登录用户名
	char password[MAXSTRING];                        //登录密码
} EQLOGININFO;

//下单信息
typedef struct _tagORDERINFO
{
	char           code[20];                         //证券代码
	double         volume;                           //(OrderMode 0:交易数量  OrderMode 1:持仓目标数量  OrderMode 2:持仓目标权重)
	float          price;                            //交易价格
	int            date;                             //交易日期 yyyymmdd  8位数
	int            time;                             //交易时间 hhmmss    6位数
	OperateType    optype;                           //操作类型
	float          cost;                             //费用
	float          rate;                             //费率
	int            reserve;                          //保留字段
} ORDERINFO, *PORDERINFO;

#pragma pack(pop)

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//日志回调函数(本接口中的日志通过此回调函数输出)
typedef int (*logcallback)(const char* pLog);

//请求回调函数(异步请求时使用)
typedef int (*datacallback) (const EQMSG* pMsg, LPVOID lpUserParam);

//设置主回调函数(一定要设置一个主回调函数，可在调用start之前调用，否则收不到账号掉线通知)
EMQUANTAPI EQErr setcallback(datacallback pfnCallback);

//用户可以自定义"ServerList.json.e"和"userInfo"文件的存放目录(如果不调用此函数或者dir传空则默认当前目录)
EMQUANTAPI void setserverlistdir(const char* dir);

//获取错误码文本说明（返回指针不要释放）
EMQUANTAPI const char* geterrstring(EQErr errcode, EQLang lang);

//证券与指标校验函数，获取相匹配的csd/css/cses的证券和指标请求参数，并按证券品种区分 options必须传入FunType=CSD 或CSS 或 CSES之一
EMQUANTAPI EQErr cfc(const char* codes, const char* indicators, const char* options, EQCTRDATA** pEQCtrData);

//校验或补全东财证券代码函数 options传入ReturnType=0/1  0:返回并标记代码是否正确 1:根据SecuType与SecuMarket补全代码后缀(有可能返回多个不同的后缀)
EMQUANTAPI EQErr cec(const char* codes, const char* options, EQCTRDATA** pEQCtrData);

//设置网络代理 注：如需使用代理，需要在调用所有接口之前设置
EMQUANTAPI EQErr setproxy(ProxyType type, const char* proxyip, unsigned short proxyport, bool verify, const char* proxyuser, const char* proxypwd);

/**人工激活 适用于无界面运行环境（如远程linux）或无法运行LoginActivator程序的情况，激活成功后将获得的激活文件"userInfo"放到"ServerList.json.e"同级目录，再调用start登录
*  参数说明：
*  pLoginInfo：账户名密码结构体指针（必传项）   options:可传邮箱,人工激活后会将令牌文件"userInfo"发送至您传入的邮箱 例如："email=who@what.com"  
*  pfnCallback：日志回调函数 */
EMQUANTAPI EQErr manualactivate(EQLOGININFO* pLoginInfo, const char* options, logcallback pfnCallback);
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/**登录(开始时或掉线后调用) 
*  参数说明：
*  pLoginInfo：保留参数,无需传(2.0.0.0版本之后改为令牌自动登陆,保留此参数以兼容旧版本)  
*  options：附加参数,用半角逗号隔开   现开放 TestLatency=1（服务器测速,默认为0不测速） 
*           ForceLogin=1 （强制登录，默认为0普通登录） LogLevel=2(日志级别 1:Debug 2:Info 3:Error)
*  pfnCallback：日志回调函数*/
EMQUANTAPI EQErr start(EQLOGININFO* pLoginInfo, const char* options, logcallback pfnCallback);
//退出(结束退出时调用，只需调用一次)
EMQUANTAPI EQErr stop();
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//参数说明:
//    codes: 东财代码  多个代码间用半角逗号隔开，支持大小写。如 "300059.SZ,000002.SZ,000003.SZ,000004.SZ"
//    indicators: 东财指标 多个指标间用半角逗号隔开，支持大小写。如 "open,close,high"
//    startdate和enddate: 开始日期和截止日期。如无分隔符，则必须为8位数字。格式支持:YYYYMMDD YYYY/MM/DD YYYY/M/D YYYY-MM-DD YYYY-M-D
//    options: 附加参数  多个参数以半角逗号隔开，"Period=1,Market=CNSESH,Order=1,Adjustflag=1,Curtype=1,Pricetype=1,Type=1"
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//指标服务数据查询(同步请求)
EMQUANTAPI EQErr csd(const char* codes, const char* indicators, const char* startDate, const char* endDate, const char* options, EQDATA** pEQData);
//截面数据查询(同步请求)
EMQUANTAPI EQErr css(const char* codes, const char* indicators, const char* options, EQDATA** pEQData);
//板块截面数据查询(同步请求)
EMQUANTAPI EQErr cses(const char* blockcodes, const char* indicators, const char* options, EQDATA** pEQData);
//获取区间日期内的交易日(同步请求)
EMQUANTAPI EQErr tradedates(const char* startDate, const char* endDate, const char* options, EQDATA** pEQData);
//获取系统板块成分(同步请求)
EMQUANTAPI EQErr sector(const char* pukeyCode, const char* tradeDate, const char* options, EQDATA** pEQData);
//获取偏移N的交易日(同步请求)
#ifdef WIN32 
EMQUANTAPI EQErr getdate(const char* tradeDate, int offDay, const char* options, EQDATA** pEQData);
#else
EMQUANTAPI EQErr gettradedate(const char* tradeDate, int offDay, const char* options, EQDATA** pEQData);
#endif
//获取区间日期内的交易日天数(同步请求)
EMQUANTAPI EQErr tradedatesnum(const char* startDate, const char* endDate, const char* options, int* num);

//分钟K线(同步请求) //code只支持单个股票/期货/指数
EMQUANTAPI EQErr csc(const char* code, const char* indicators, const char* startDate, const char* endDate, const char* options, EQDATA** pEQData);
//分钟K线(同步请求) //code只支持单个股票/期货/指数
EMQUANTAPI EQErr cmc(const char* code, const char* indicators, const char* startDate, const char* endDate, const char* options, EQDATA** pEQData);
//分钟K线(同步请求) //code只支持单个股票/期货/指数
EMQUANTAPI EQErr chmc(const char* code, const char* indicators, const char* startDate, const char* endDate, const char* options, EQDATA** pEQData);

//行情快照(同步请求) 每次indicators最多为64个
EMQUANTAPI EQErr csqsnapshot(const char* codes, const char* indicators, const char* options, EQDATA** pEQData);
//专项快照(同步请求)
EMQUANTAPI EQErr chqsnapshot(const char* codes, const char* indicators, const char* options, EQDATA** pEQData);
//获取专题报表(同步请求)
EMQUANTAPI EQErr ctr(const char* ctrName, const char* indicators, const char* options, EQCTRDATA** pEQCtrData);

/*条件选股(同步请求)
* 参数说明：
* cpsCodes:      证券代码范围,必传,支持两种模式 1.板块代码 以 B_ 开头，如 "B_001004"; 2. 东财代码，多个代码间用半角逗号隔开，如"000001.SZ,000002.SZ,600000.SH"
* cpsIndicators: 表达式参数,表达式之前用英文分号隔开,内部各参数用半角逗号隔开 如   s1,open,2016/12/13,1;s2,close,2017-02-25,1;s3,listdate
* cpsConditions: 条件表达式
* cpsOptions:    其他附加条件,如排序,取前N条选股结果等
* pEQData:       选股结果*/
EMQUANTAPI EQErr cps(const char* cpsCodes, const char* cpsIndicators, const char* cpsConditions, const char* cpsOptions, EQDATA** pEQData);
//组合账户关系查询(同步请求)
EMQUANTAPI EQErr pquery(const char* options, EQDATA** pEQData);
//批量下单(同步请求)  pOrderInfo:下单信息指针数组  orderInfoSize：数组元素个数  combinCode：组合代码  remark：组合说明
EMQUANTAPI EQErr porder(ORDERINFO* pOrderInfo, int orderInfoSize, const char* combinCode, const char* remark, const char* options);
//新建组合(同步请求) combinCode:组合代码  combinName：组合名称  initialFound：初始资金（最大99999999999）  remark：组合说明
EMQUANTAPI EQErr pcreate(const char* combinCode, const char* combinName, int64_t initialFound, const char* remark, const char* options);
//删除组合(同步请求) combinCode:组合代码
EMQUANTAPI EQErr pdelete(const char* combinCode, const char* options);
//组合报表查询(同步请求) combinCode:组合代码  indicator:报表名称
EMQUANTAPI EQErr preport(const char* combinCode, const char* indicator, const char* options, EQDATA** pEQData);
//组合资金调配(同步请求) combinCode:组合代码  transferdirect:IN 增加资金 OUT 减少资金  date：调配日期  opCash：增加或减少的资金量 remark：操作说明
EMQUANTAPI EQErr pctransfer(const char* combinCode, const char* transferdirect, const char* date, double opCash, const char* remark, const char* options);

//仅供本API中同步接口返回数据指针释放内存(EQDATA* 或 EQCTRDATA* 或者 EQCHAR*, 不可传入其他指针，异步函数回调中的指针也不可传入)
EMQUANTAPI EQErr releasedata(void* pEQData);
//宏观指标服务(同步请求)
EMQUANTAPI EQErr edb(const char* edbids, const char* options, EQDATA** pEQData);
//宏观指标id详情查询(同步请求)
EMQUANTAPI EQErr edbquery(const char* edbids, const char* indicators, const char* options, EQDATA** pEQData);

//资讯数据查询(同步请求) codes：东财代码或板块代码（不可混合） content：查询内容
EMQUANTAPI EQErr cfn(const char* codes, const char* content, eCfnMode emode, const char* options, EQDATA** pEQData);
//板块树查询（同步请求）
EMQUANTAPI EQErr cfnquery(const char* options, EQDATA** pEQData);


//↑↑↑↑↑↑↑以上为同步函数↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑
//↓↓↓↓↓↓↓以下为异步函数↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
//返回： EQID  流水号
//pfnCallback：不同的接口可以设定不同的回调，传NULL则使用默认的主回调函数
//lpUserParam: 用户参数,回调时原样返回

//实时行情(异步)  每次indicators最多为64个 options: Pushtype=0 增量推送  1全量推送  2增量推送2(证券增量，指标值全量)
EMQUANTAPI EQID csq(const char* codes, const char* indicators, const char* options, datacallback pfnCallback, LPVOID lpUserParam, EQErr* nErrorID);
//取消实时行情订阅   serialID为0时 取消所有订阅
EMQUANTAPI EQErr csqcancel(EQID serialID);

//日内跳价服务(异步)  startdatetime和enddatetime格式(YYYYMMDDHHMMSS或HHMMSS表示系统日期当天的时间，两者需使用同一种格式)
EMQUANTAPI EQID cst(const char* codes, const char* indicators, const char* startdatetime, const char* enddatetime, const char* options, datacallback pfnCallback, LPVOID lpUserParam, EQErr* nErrorID);

//资讯订阅函数(异步)  codes：东财代码或板块代码（不可混合） content：订阅内容
EMQUANTAPI EQID cnq(const char* codes, const char* content, const char* options, datacallback pfnCallback, LPVOID lpUserParam, EQErr* nErrorID);
////资讯取消订阅  serialID为0时 取消所有订阅
EMQUANTAPI EQErr cnqcancel(EQID serialID);

//专项订阅(异步)  options: Pushtype=0 增量推送  1全量推送  2增量推送2(证券增量，指标值全量)
EMQUANTAPI EQID chq(const char* codes, const char* indicators, const char* options, datacallback pfnCallback, LPVOID lpUserParam, EQErr* nErrorID);
//取消专项订阅   serialID为0时 取消所有订阅
EMQUANTAPI EQErr chqcancel(EQID serialID);

//ai债券分析 文本  text:文本内容 options:RECVTIMEOUT接收超时默认10秒 pEQChar:返回字符串，用完需使用releasedata(pEQChar)
EMQUANTAPI EQErr aibondtext(const char* text, const char* options, EQCHAR** pEQChar);
//ai债券分析 图片  context:图片字节流 contentLen：字节流长度   options:RECVTIMEOUT接收超时默认10秒 PICMODE:0(content传字节流，默认) 1(content传文件名全路径)   pEQChar:返回字符串，用完需使用releasedata(pEQChar)
EMQUANTAPI EQErr aibondpic(const char* context, int contentLen, const char* options, EQCHAR** pEQChar);

//流量查询函数   funcname:单个函数名称，不传或为空时表示查询所有   indicators：查询的指标列，多个用英文逗号分隔，不传或为空则查询所有字段   options:可传STARTDATE ENDDATE    用完需使用releasedata(pEQCtrData)
EMQUANTAPI EQErr datastatistics(const char* funcname, const char* indicators, const char* options, EQCTRDATA** pEQCtrData);


#endif /*#ifndef _EMQUANTAPI_H_*/