package login

import (
	"context"
	"fmt"
	"strings"

	"yxy-go/internal/consts"
	"yxy-go/internal/svc"
	"yxy-go/internal/types"
	"yxy-go/internal/utils/yxyClient"
	"yxy-go/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginByCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByCodeLogic {
	return &LoginByCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type LoginByCodeYxyResp struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       struct {
		ID                    string   `json:"id"`
		SchoolCode            string   `json:"schoolCode"`
		BadgeImg              string   `json:"badgeImg"`
		SchoolName            string   `json:"schoolName"`
		QrcodePayType         uint8    `json:"qrcodePayType"`
		Account               string   `json:"account"`
		AccountEncrypt        string   `json:"accountEncrypt"`
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
		IsNew                 uint8    `json:"isNew"`
		CreateStatus          uint8    `json:"createStatus"`
		EacctStatus           uint8    `json:"eacctStatus"`
		SchoolClasses         uint8    `json:"schoolClasses"`
		SchoolNature          uint8    `json:"schoolNature"`
		Platform              string   `json:"platform"`
		CardPhone             string   `json:"cardPhone"`
		QrcodePrivateKey      string   `json:"qrcodePrivateKey"`
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

func (l *LoginByCodeLogic) LoginByCode(req *types.LoginByCodeReq) (resp *types.LoginByCodeResp, err error) {
	yxyReq, yxyHeaders := yxyClient.GetYxyBaseReqParam(req.DeviceID)
	yxyReq["mobilePhone"] = req.PhoneNum
	yxyReq["verificationCode"] = req.Code
	yxyReq["appAllVersion"] = consts.APP_ALL_VERSION
	yxyReq["appPlatform"] = "Android"
	yxyReq["brand"] = "Android"
	yxyReq["clientId"] = consts.CLIENT_ID
	yxyReq["invitationCode"] = ""
	yxyReq["mobileType"] = "Android for arm64"
	yxyReq["osType"] = "Android"
	yxyReq["osVersion"] = "12"

	var yxyResp LoginByCodeYxyResp
	r, err := yxyClient.HttpSendPost(consts.LOGIN_BY_CODE_URL, yxyReq, yxyHeaders, &yxyResp)
	if err != nil {
		return nil, err
	}

	if yxyResp.StatusCode != 0 {
		errCode := xerr.ErrUnknown
		if yxyResp.Message == "手机号格式不正确" {
			errCode = xerr.ErrPhoneNumWrong
		} else if strings.HasSuffix(yxyResp.Message, "3次过后将锁定15分钟,请谨慎操作") { // 您已输错(1,2,3)次, 3次过后将锁定15分钟,请谨慎操作
			errCode = xerr.ErrCodeWrong
		} else if yxyResp.Message == "您已输错3次,账号被锁定15分钟" {
			errCode = xerr.ErrCodeWrongThreeTimes
		}
		return nil, xerr.WithCode(errCode, fmt.Sprintf("yxy response: %v", r))
	}

	return &types.LoginByCodeResp{
		UID:            yxyResp.Data.ID,
		Token:          yxyResp.Data.Token,
		BindCardStatus: yxyResp.Data.BindCardStatus,
	}, nil
}
