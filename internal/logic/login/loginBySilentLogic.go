package login

import (
	"context"
	"fmt"

	"yxy-go/internal/consts"
	"yxy-go/internal/svc"
	"yxy-go/internal/types"
	"yxy-go/internal/utils/yxyClient"
	"yxy-go/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginBySilentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginBySilentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginBySilentLogic {
	return &LoginBySilentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type LoginBySilentYxyResp struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       struct {
		ID                    string   `json:"id"`
		SchoolCode            string   `json:"schoolCode"`
		BadgeImg              string   `json:"badgeImg"`
		SchoolName            string   `json:"schoolName"`
		QrcodePayType         uint8    `json:"qrcodePayType"`
		Account               string   `json:"account"`
		UserName              string   `json:"userName"`
		UserType              string   `json:"userType"`
		MobilePhone           string   `json:"mobilePhone"`
		JobNo                 string   `json:"jobNo"`
		UserIdcard            string   `json:"userIdcard"`
		IdentityNo            string   `json:"identityNo"`
		Sex                   uint8    `json:"sex"`
		UserClass             string   `json:"userClass"`
		RealNameStatus        uint8    `json:"realNameStatus"`
		RegisterTime          string   `json:"regiserTime"` // time
		Birthday              string   `json:"birthday"`    // time
		UserStatus            uint8    `json:"userStatus"`
		BindCardStatus        uint8    `json:"bindCardStatus"`
		BindCardTime          string   `json:"bindCardTime"` // time
		LastLogin             string   `json:"lastLogin"`    // time
		HeadImg               string   `json:"headImg"`
		DeviceID              string   `json:"deviceId"`
		TestAccount           uint8    `json:"testAccount"`
		Token                 string   `json:"token"`
		TokenList             []string `json:"tokenList"`
		LastTokenTime         string   `json:"lastTokenTime"` // time
		JoinNewActivityStatus uint8    `json:"joinNewactivityStatus"`
		CreateStatus          uint8    `json:"createStatus"`
		EacctStatus           uint8    `json:"eacctStatus"`
		SchoolClasses         uint8    `json:"schoolClasses"`
		SchoolNature          uint8    `json:"schoolNature"`
		Platform              string   `json:"platform"`
		CardPhone             string   `json:"cardPhone"`
		BindCardRate          uint8    `json:"bindCardRate"`
		Points                uint8    `json:"points"`
		CardIdentityType      uint8    `json:"cardIdentityType"`
		SchoolIdentityType    uint8    `json:"schoolIdentityType"`
		AlumniFlag            uint8    `json:"alumniFlag"`
		ExtJson               string   `json:"extJson"`
		AuthType              uint8    `json:"authType"`
		JoinChatStatus        uint8    `json:"joinChatStatus"`
		QywechatContactStatus uint8    `json:"qywechatContactStatus"`
	} `json:"data"`
	Success bool `json:"success"`
}

func (l *LoginBySilentLogic) LoginBySilent(req *types.LoginBySilentReq) (resp *types.LoginBySilentResp, err error) {
	yxyReq, yxyHeaders := yxyClient.GetYxyBaseReqParam(req.DeviceID)
	yxyReq["appAllVersion"] = consts.APP_ALL_VERSION
	yxyReq["appPlatform"] = "Android"
	yxyReq["brand"] = "Android"
	yxyReq["clientId"] = consts.CLIENT_ID
	yxyReq["mobilePhone"] = req.PhoneNum
	yxyReq["mobileType"] = "Android for arm64"
	yxyReq["osType"] = "Android"
	yxyReq["osVersion"] = "12"
	yxyReq["token"] = req.Token
	yxyReq["ymId"] = req.UID

	var yxyResp LoginBySilentYxyResp
	r, err := yxyClient.HttpSendPost(consts.LOGIN_BY_Silent_URL, yxyReq, yxyHeaders, &yxyResp)
	if err != nil {
		return nil, err
	}

	if yxyResp.StatusCode != 0 {
		errCode := xerr.ErrUnknown
		if yxyResp.Message == "登录已过期，请重新登录[user no find]" {
			errCode = xerr.ErrUserNotFound
		} else if yxyResp.Message == "您的账号已被登出，请重新登录[deviceId changed]" {
			errCode = xerr.ErrAccountLoggedOut
		}
		return nil, xerr.WithCode(errCode, fmt.Sprintf("yxy response: %v", r))
	}

	return nil, nil
}
