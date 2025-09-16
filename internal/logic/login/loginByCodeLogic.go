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
		Id string `json:"id"`
		//SchoolCode            string `json:"schoolCode"`
		//SchoolName            string `json:"schoolName"`
		//QrcodePayType         int    `json:"qrcodePayType"`
		//Account               string `json:"account"`
		//AccountEncrypt        string `json:"accountEncrypt"`
		//MobilePhone           string `json:"mobilePhone"`
		//Sex                   int    `json:"sex"`
		//RealNameStatus        int    `json:"realNameStatus"`
		//RegiserTime           string `json:"regiserTime"`
		//NickName              string `json:"nickName"`
		//UserStatus            int    `json:"userStatus"`
		BindCardStatus uint8 `json:"bindCardStatus"`
		//LastLogin             string `json:"lastLogin"`
		//HeadImg               string `json:"headImg"`
		DeviceId string `json:"deviceId"`
		//TestAccount           int    `json:"testAccount"`
		//Token                 string `json:"token"`
		//JoinNewactivityStatus int    `json:"joinNewactivityStatus"`
		//IsNew                 int    `json:"isNew"`
		//CreateStatus          int    `json:"createStatus"`
		//EacctStatus           int    `json:"eacctStatus"`
		//SchoolClasses         int    `json:"schoolClasses"`
		//SchoolNature          int    `json:"schoolNature"`
		//Platform              string `json:"platform"`
		//QrcodePrivateKey      string `json:"qrcodePrivateKey"`
		//BindCardRate          int    `json:"bindCardRate"`
		//Points                int    `json:"points"`
		//CardIdentityType      int    `json:"cardIdentityType"`
		//SchoolIdentityType    int    `json:"schoolIdentityType"`
		//AlumniFlag            int    `json:"alumniFlag"`
		//ExtJson               string `json:"extJson"`
		//AuthType              int    `json:"authType"`
		//JoinChatStatus        int    `json:"joinChatStatus"`
		//QywechatContactStatus int    `json:"qywechatContactStatus"`
		//KeyMap                struct {
		//	Field1 string `json:"730"`
		//} `json:"keyMap"`
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
		} else if strings.HasSuffix(yxyResp.Message, "3次过后将锁定15分钟,请慎重操作") { // 您已输错(1,2,3)次,3次过后将锁定15分钟,请慎重操作
			errCode = xerr.ErrCodeWrong
		} else if yxyResp.Message == "您已输错3次,账户被锁定15分钟" {
			errCode = xerr.ErrCodeWrongThreeTimes
		}
		return nil, xerr.WithCode(errCode, fmt.Sprintf("yxy response: %v", r))
	}

	return &types.LoginByCodeResp{
		UID:            yxyResp.Data.Id,
		BindCardStatus: yxyResp.Data.BindCardStatus,
	}, nil
}
