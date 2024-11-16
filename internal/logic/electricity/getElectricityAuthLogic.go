package electricity

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"yxy-go/internal/consts"
	"yxy-go/internal/svc"
	"yxy-go/internal/types"
	"yxy-go/internal/utils/yxyClient"
	"yxy-go/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetElectricityAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetElectricityAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetElectricityAuthLogic {
	return &GetElectricityAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type GetElectricityAuthTokenYxyResp struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       struct {
		ID             string `json:"id"`
		SchoolCode     string `json:"schoolCode"`
		SchoolName     string `json:"schoolName"`
		UserName       string `json:"userName"`
		UserType       string `json:"userType"`
		MobilePhone    string `json:"mobilePhone"`
		JobNo          string `json:"jobNo"`
		UserIdcard     string `json:"userIdcard"`
		Sex            uint8  `json:"sex"`
		UserClass      string `json:"userClass"`
		BindCardStatus uint8  `json:"bindCardStatus"`
		TestAccount    uint8  `json:"testAccount"`
		Platform       string `json:"platform"`
		ThirdOpenid    string `json:"thirdOpenid"`
	} `json:"data"`
	Success bool `json:"success"`
}

func (l *GetElectricityAuthLogic) GetElectricityAuth(req *types.GetElectricityAuthReq) (resp *types.GetElectricityAuthResp, err error) {
	_, yxyHeaders := yxyClient.GetYxyBaseReqParam("")
	yxyReq := map[string]string{
		"bindSkip":    "1",
		"authType":    "2",
		"ymAppId":     "1810181825222034",
		"callbackUrl": consts.APPLICATION_URL + "/",
		"unionid":     req.UID,
		"schoolCode":  consts.SCHOOL_CODE,
	}

	client := yxyClient.GetClient()
	r, err := client.R().
		SetHeaders(yxyHeaders).
		SetQueryParams(yxyReq).
		Get(consts.GET_ELECTRICITY_AUTH_CODE_URL)
	if err != nil && r.StatusCode() != 302 {
		return nil, xerr.WithCode(xerr.ErrHttpClient, err.Error())
	}

	location := r.Header().Get("Location")
	if location == "" {
		if strings.Contains(r.String(), "用户不存在") {
			return nil, xerr.WithCode(xerr.ErrUserNotFound, fmt.Sprintf("User not found, UID: %v", req.UID))
		}
		return nil, xerr.WithCode(xerr.ErrUnknown, fmt.Sprintf("yxy response: %v", r))
	}
	parsedURL, _ := url.Parse(location)
	ymCode := parsedURL.Query().Get("ymCode")

	var yxyResp GetElectricityAuthTokenYxyResp
	r, err = yxyClient.HttpSendPost(consts.GET_ELECTRICITY_AUTH_TOKEN_URL,
		map[string]interface{}{
			"code": ymCode,
		}, yxyHeaders, &yxyResp)
	if err != nil {
		return nil, err
	}

	if yxyResp.StatusCode != 0 {
		return nil, xerr.WithCode(xerr.ErrUnknown, fmt.Sprintf("yxy response: %v", r))
	}
	var shiroJID string
	for _, cookie := range r.Cookies() {
		if cookie.Name == "shiroJID" {
			shiroJID = cookie.Value
			break
		}
	}

	return &types.GetElectricityAuthResp{
		Token: shiroJID,
	}, nil
}
