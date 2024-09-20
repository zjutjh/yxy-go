package consts

const (
	APP_ALL_VERSION = "6.2.8"
	APP_VERSION     = "611"
	CLIENT_ID       = "65l281jbb6wjvuo"
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
)

const (
	GET_CARD_BALANCE_URL             = COMPUS_URL + "/compus/user/getCardMoney"
	GET_CARD_CONSUMPTION_RECORDS_URL = COMPUS_URL + "/routeauth/auth/route/user/cardQuerynoPage"
)

const (
	GET_ELECTRICITY_AUTH_CODE_URL    = AUTH_URL + "/authoriz/getCodeV2"
	GET_ELECTRICITY_AUTH_TOKEN_URL   = APPLICATION_URL + "/app/login/getUser4Authorize"
	QUERY_ELECTRICITY_BIND_URL       = APPLICATION_URL + "/app/electric/queryBind"
	GET_ELECTRICITY_ZHPF_SURPLUS_URL = APPLICATION_URL + "/app/electric/queryISIMSRoomSurplus"
	GET_ELECTRICITY_MGS_SURPLUS_URL  = APPLICATION_URL + "/app/electric/queryRoomSurplus"
)
