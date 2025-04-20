package consts

const (
	APP_ALL_VERSION = "6.2.8"
	APP_VERSION     = "611"
	CLIENT_ID       = "65l281jbb6wjvuo"
	SCHOOL_CODE     = "10337" // ZJUT
	COMPUS_URL      = "https://compus.xiaofubao.com"
	ELECTRICITY_URL = "https://application.xiaofubao.com"
	AUTH_URL        = "https://auth.xiaofubao.com"
	BUS_URL         = "https://api.pinbayun.com"
	BUS_AUTH_URL    = "https://open.xiaofubao.com"
)

const (
	GET_SECURITY_TOKEN_URL = COMPUS_URL + "/common/security/token"
	GET_CAPTCHA_IMAGE_URL  = COMPUS_URL + "/common/security/imageCaptcha"
	SEND_CODE_URL          = COMPUS_URL + "/compus/user/sendLoginVerificationCode"
	LOGIN_BY_CODE_URL      = COMPUS_URL + "/login/doLoginByVerificationCode"
	LOGIN_BY_Silent_URL    = COMPUS_URL + "/login/doLoginBySilent"
)

const (
	GET_CARD_BALANCE_URL             = COMPUS_URL + "/compus/user/getCardMoney"
	GET_CARD_CONSUMPTION_RECORDS_URL = COMPUS_URL + "/routeauth/auth/route/user/cardQuerynoPage"
)

const (
	GET_ELECTRICITY_AUTH_CODE_URL             = AUTH_URL + "/authoriz/getCodeV2"
	GET_ELECTRICITY_AUTH_TOKEN_URL            = ELECTRICITY_URL + "/app/login/getUser4Authorize"
	QUERY_ELECTRICITY_BIND_URL                = ELECTRICITY_URL + "/app/electric/queryBind"
	GET_ELECTRICITY_ZHPF_SURPLUS_URL          = ELECTRICITY_URL + "/app/electric/queryISIMSRoomSurplus"
	GET_ELECTRICITY_MGS_SURPLUS_URL           = ELECTRICITY_URL + "/app/electric/queryRoomSurplus"
	GET_ELECTRICITY_ZHPF_RECHARGE_RECORDS_URL = ELECTRICITY_URL + "/app/electric/queryISIMSRoomBuyRecord"
	GET_ELECTRICITY_MGS_RECHARGE_RECORDS_URL  = ELECTRICITY_URL + "/app/electric/roomBuyRecord"
	GET_ELECTRICITY_ZHPF_USAGE_RECORDS_URL    = ELECTRICITY_URL + "/app/electric/getISIMSRecords"
	GET_ELECTRICITY_MGS_USAGE_RECORDS_URL     = ELECTRICITY_URL + "/app/electric/queryUsageRecord"
)

const (
	GET_BUS_AUTH_CODE_URL  = BUS_AUTH_URL + "/routeauth/auth/route/ua/authorize/getCodeV2"
	GET_BUS_AUTH_TOKEN_URL = BUS_URL + "/api/v1/staff/auths/wx_auth/"
	GET_BUS_ACCESS_URL     = AUTH_URL + "/auth/route/authorize/agreementAuth"
	GET_BUS_INFO_URL       = BUS_URL + "/api/v2/staff/shuttlebus/"
	GET_BUS_TIME_URL       = BUS_URL + "/api/v2/staff/shuttlebus/{id}/bustimes/"
	GET_BUS_DATE_URL       = BUS_URL + "/api/v2/staff/shuttlebus/{id}/dates/"
	GET_BUS_RECORD_URL     = BUS_URL + "/api/v1/staff/busorders/"
	GET_BUS_QRCODE_URL     = BUS_URL + "/api/v3/pos/staff_qrcode/"
	GET_BUS_MESSAGE_URL    = BUS_URL + "/api/v1/staff/messages/"
	// GET_BUS_MESSAGE_UNREAD_COUNT_URL = BUS_URL + "/api/v1/staff/messages/unread_count/"
)
