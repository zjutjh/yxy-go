package consts

const (
	APP_ALL_VERSION = "7.3.6"
	APP_VERSION     = "730"
	CLIENT_ID       = "65l3attk4r095ib"
	SCHOOL_CODE     = "10337" // ZJUT
	COMPUS_URL      = "https://compus.xiaofubao.com"
	APPLICATION_URL = "https://application.xiaofubao.com"
	AUTH_URL        = "https://auth.xiaofubao.com"
)

const (
	GET_SECURITY_TOKEN_URL = COMPUS_URL + "/common/security/token"
	GET_CAPTCHA_IMAGE_URL  = COMPUS_URL + "/common/security/imageCaptcha"
	SEND_CODE_URL          = COMPUS_URL + "/compus/user/sendLoginVerificationCode"
	LOGIN_BY_CODE_URL      = COMPUS_URL + "/login/doLoginByVerificationCode"
	LOGIN_BY_Silent_URL    = COMPUS_URL + "/login/doLoginBySilent"
	GET_AUTH_TOKEN         = COMPUS_URL + "/compus/user/getAuthToken"
)

const (
	GET_CARD_BALANCE_URL             = COMPUS_URL + "/compus/user/getCardMoney"
	GET_CARD_CONSUMPTION_RECORDS_URL = COMPUS_URL + "/routeauth/auth/route/user/cardQuerynoPage"
)

const (
	GET_ELECTRICITY_AUTH_CODE_URL             = AUTH_URL + "/authoriz/getCodeV2"
	GET_ELECTRICITY_AUTH_TOKEN_URL            = APPLICATION_URL + "/app/login/getUser4Authorize"
	QUERY_ELECTRICITY_BIND_URL                = APPLICATION_URL + "/app/electric/queryBind"
	GET_ELECTRICITY_ZHPF_SURPLUS_URL          = APPLICATION_URL + "/app/electric/queryISIMSRoomSurplus"
	GET_ELECTRICITY_MGS_SURPLUS_URL           = APPLICATION_URL + "/app/electric/queryRoomSurplus"
	GET_ELECTRICITY_ZHPF_RECHARGE_RECORDS_URL = APPLICATION_URL + "/app/electric/queryISIMSRoomBuyRecord"
	GET_ELECTRICITY_MGS_RECHARGE_RECORDS_URL  = APPLICATION_URL + "/app/electric/roomBuyRecord"
	GET_ELECTRICITY_ZHPF_USAGE_RECORDS_URL    = APPLICATION_URL + "/app/electric/getISIMSRecords"
	GET_ELECTRICITY_MGS_USAGE_RECORDS_URL     = APPLICATION_URL + "/app/electric/queryUsageRecord"
)
