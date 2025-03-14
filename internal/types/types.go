// Code generated by goctl. DO NOT EDIT.
package types

type CardConsumptionRecord struct {
	Address string `json:"address"`
	Money   string `json:"money"`
	Time    string `json:"time"`
}

type ElectricityRechargeRecord struct {
	Money    string `json:"money"`
	Datetime string `json:"datetime"`
}

type ElectricityUsageRecord struct {
	Usage    string `json:"usage"`
	Datetime string `json:"datetime"`
}

type GetCaptchaImageReq struct {
	DeviceID      string `form:"device_id"`
	SecurityToken string `form:"security_token"`
}

type GetCaptchaImageResp struct {
	Img string `json:"img"`
}

type GetCardBalanceReq struct {
	UID      string `form:"uid"`
	DeviceID string `form:"device_id"`
	Token    string `form:"token"`
}

type GetCardBalanceResp struct {
	Balance string `json:"balance"`
}

type GetCardConsumptionRecordsReq struct {
	UID       string `form:"uid"`
	DeviceID  string `form:"device_id"`
	Token     string `form:"token"`
	QueryTime string `form:"query_time"`
}

type GetCardConsumptionRecordsResp struct {
	List []CardConsumptionRecord `json:"list"`
}

type GetElectricityAuthReq struct {
	UID string `form:"uid"`
}

type GetElectricityAuthResp struct {
	Token string `json:"token"`
}

type GetElectricityRechargeRecordsReq struct {
	Token         string `form:"token"`
	Campus        string `form:"campus,options=zhpf|mgs"`
	Page          string `form:"page"`
	RoomStrConcat string `form:"room_str_concat"`
}

type GetElectricityRechargeRecordsResp struct {
	List []ElectricityRechargeRecord `json:"list"`
}

type GetElectricitySurplusReq struct {
	Token  string `form:"token"`
	Campus string `form:"campus,options=zhpf|mgs"`
}

type GetElectricitySurplusResp struct {
	DisplayRoomName string  `json:"display_room_name"`
	RoomStrConcat   string  `json:"room_str_concat"`
	Surplus         float64 `json:"surplus"`
}

type GetElectricityUsageRecordsReq struct {
	Token         string `form:"token"`
	Campus        string `form:"campus,options=zhpf|mgs"`
	RoomStrConcat string `form:"room_str_concat"`
}

type GetElectricityUsageRecordsResp struct {
	List []ElectricityUsageRecord `json:"list"`
}

type GetSecurityTokenReq struct {
	DeviceID string `form:"device_id"`
}

type GetSecurityTokenResp struct {
	Level         uint8  `json:"level"`
	SecurityToken string `json:"security_token"`
}

type LoginByCodeReq struct {
	DeviceID string `json:"device_id"`
	PhoneNum string `json:"phone_num"`
	Code     string `json:"code"`
}

type LoginByCodeResp struct {
	UID            string `json:"uid"`
	Token          string `json:"token"`
	BindCardStatus uint8  `json:"bind_card_status"`
}

type LoginBySilentReq struct {
	UID      string `json:"uid"`
	DeviceID string `json:"device_id"`
	PhoneNum string `json:"phone_num,optional"`
	Token    string `json:"token"`
}

type LoginBySilentResp struct {
}

type SendCodeReq struct {
	DeviceID      string `json:"device_id"`
	SecurityToken string `json:"security_token"`
	Captcha       string `json:"captcha,optional"`
	PhoneNum      string `json:"phone_num"`
}

type SendCodeResp struct {
	UserExists bool `json:"user_exists"`
}
