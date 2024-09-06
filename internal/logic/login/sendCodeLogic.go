package login

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"yxy-go/internal/consts"
	"yxy-go/internal/svc"
	"yxy-go/internal/types"
	"yxy-go/internal/utils/yxyClient"
	"yxy-go/pkg/xerr"

	"github.com/forgoer/openssl"
	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type SendCodeYxyResp struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       struct {
		UserExists bool `json:"userExists"`
	} `json:"data"`
	Success bool `json:"success"`
}

func (l *SendCodeLogic) SendCode(req *types.SendCodeReq) (resp *types.SendCodeResp, err error) {
	yxyReq, yxyHeaders := yxyClient.GetYxyBaseReqParam(req.DeviceID)
	yxyReq["mobilePhone"] = req.PhoneNum
	yxyReq["securityToken"] = req.SecurityToken
	yxyReq["sendCount"] = 1

	appSecurityToken, err := getAppSecurityToken(yxyReq["deviceId"].(string), req.SecurityToken)
	if err != nil {
		return nil, err
	}
	yxyReq["appSecurityToken"] = appSecurityToken

	if req.Captcha != "" {
		yxyReq["imageCaptchaValue"] = req.Captcha
	}

	var yxyResp SendCodeYxyResp
	r, err := yxyClient.HttpSendPost(consts.SEND_CODE_URL, yxyReq, yxyHeaders, &yxyResp)
	if err != nil {
		return nil, err
	}

	if yxyResp.StatusCode != 0 {
		errCode := xerr.ErrUnknown
		switch yxyResp.Message {
		case "验证码已失效":
			errCode = xerr.ErrCaptchaInvalid
		case "验证码错误":
			errCode = xerr.ErrCaptchaWrong
		case "encryptedDeviceId不一致":
			errCode = xerr.ErrDeviceIDInconsistent
		case "请输入正确的手机号":
			errCode = xerr.ErrPhoneNumWrong
		case "一分钟内只能发送一次短信,请稍后再试", "短信发送超限，请一分钟后再试":
			errCode = xerr.ErrSendLimit
		}
		return nil, xerr.WithCode(errCode, fmt.Sprintf("yxy response: %v", r))
	}

	return &types.SendCodeResp{
		UserExists: yxyResp.Data.UserExists,
	}, nil
}

func getAppSecurityToken(deviceID, securityToken string) (appSecurityToken string, err error) {
	if len(securityToken) != 56 {
		return "", xerr.WithCode(xerr.ErrTokenInvalid, "Invalid security token length")
	}

	key := []byte(securityToken[:16])
	token := securityToken[32:]

	cipherText, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", xerr.WithCode(xerr.ErrTokenInvalid, err.Error())
	}

	plainText, err := openssl.AesECBDecrypt(cipherText, key, openssl.PKCS7_PADDING)
	if err != nil {
		return "", xerr.WithCode(xerr.ErrTokenInvalid, err.Error())
	}
	t := string(plainText)

	timestampNano := time.Now().UnixNano()
	seconds := timestampNano / 1e9
	microseconds := (timestampNano % 1e9) / 1e2
	ts := fmt.Sprintf("%v.%v", seconds, microseconds)

	md5Hash1 := md5.Sum([]byte(deviceID + "|YUNMA_APP|" + t + "|" + ts + "|" + consts.APP_ALL_VERSION))
	md5HashStrUpper1 := strings.ToUpper(hex.EncodeToString(md5Hash1[:]))
	md5Hash2 := md5.Sum([]byte(md5HashStrUpper1))
	s := strings.ToUpper(hex.EncodeToString(md5Hash2[:]))

	encrypted, err := openssl.AesECBEncrypt([]byte(deviceID+"|YUNMA_APP|"+t+"|"+ts+"|"+consts.APP_ALL_VERSION+"|"+s), key, openssl.PKCS7_PADDING)
	if err != nil {
		return "", xerr.WithCode(xerr.ErrTokenInvalid, err.Error())
	}

	appSecurityToken = base64.StdEncoding.EncodeToString(encrypted)

	return appSecurityToken, nil
}
