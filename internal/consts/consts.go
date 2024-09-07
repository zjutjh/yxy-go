package consts

const (
	APP_ALL_VERSION = "6.2.8"
	APP_VERSION     = "611"
	CLIENT_ID       = "65l281jbb6wjvuo"
	SCHOOL_CODE     = "10337" // ZJUT
	COMPUS_URL      = "https://compus.xiaofubao.com"
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
