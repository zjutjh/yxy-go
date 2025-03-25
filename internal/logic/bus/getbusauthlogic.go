package bus

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"yxy-go/internal/consts"
	"yxy-go/internal/svc"
	"yxy-go/internal/types"
	"yxy-go/internal/utils/yxyClient"
	"yxy-go/pkg/xerr"

	"github.com/PuerkitoBio/goquery"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetBusAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBusAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBusAuthLogic {
	return &GetBusAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBusAuthLogic) GetBusAuth(req *types.GetBusAuthReq) (resp *types.GetBusAuthResp, err error) {
	_, yxyHeaders := yxyClient.GetYxyBaseReqParam("")
	yxyReq := map[string]string{
		"authType":    "2",
		"ymAppId":     "2011112043190345310",
		"authAppid":   "10337",
		"callbackUrl": "https://api.pinbayun.com/api/v1/zjgd_interface/?schoolCode=10337",
		"unionid":     req.UID,
		"schoolCode":  consts.SCHOOL_CODE,
	}

	client := yxyClient.GetClient()

	r, err := client.R().
		SetHeaders(yxyHeaders).
		SetQueryParams(yxyReq).
		Get(consts.GET_BUS_AUTH_CODE_URL)

	if err != nil && (r.StatusCode() != 302 && r.StatusCode() != 200) {
		return nil, xerr.WithCode(xerr.ErrHttpClient, err.Error())
	}

	location := r.Header().Get("Location")
	fmt.Println("location: " + location)
	if location == "" {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(r.String()))
		if err != nil {
			return nil, xerr.WithCode(xerr.ErrUnknown, err.Error())
		}
		stateCode := doc.Find("input.stateCode").AttrOr("value", "")
		if stateCode == "" {
			return nil, xerr.WithCode(xerr.ErrUserNotFound, "登录失败")
		}

		r, err = client.R().
			SetHeaders(yxyHeaders).
			SetQueryParams(map[string]string{
				"stateCode": stateCode,
				"appid":     "2011112043190345310",
			}).
			Get(consts.GET_BUS_ACCESS_URL)
		if err != nil && r.StatusCode() != 302 {
			return nil, xerr.WithCode(xerr.ErrUserNotFound, "登录失败")
		}

		location = r.Header().Get("Location")
	}

	r, err = client.R().
		SetHeaders(yxyHeaders).
		Get(location)

	if err != nil && r.StatusCode() != 302 {
		return nil, xerr.WithCode(xerr.ErrUserNotFound, "登录失败")
	}

	redirectURL := r.Header().Get("Location")
	fmt.Println(redirectURL)

	if redirectURL == "" {
		return nil, xerr.WithCode(xerr.ErrUserNotFound, "登录失败")
	}

	parsedURL, err := url.Parse(redirectURL)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrUnknown, err.Error())
	}

	queryParams := parsedURL.Query()
	openid := queryParams.Get("openid")
	corpcode := queryParams.Get("corpcode")

	if openid == "" || corpcode == "" {
		return nil, xerr.WithCode(xerr.ErrUserNotFound, "登录失败")
	}

	r, err = client.R().
		SetBody(map[string]string{
			"openid":   openid,
			"corpcode": corpcode,
		}).
		Post(consts.GET_BUS_AUTH_TOKEN_URL)

	if err != nil {
		return nil, xerr.WithCode(xerr.ErrHttpClient, err.Error())
	}

	var result map[string]interface{}
	if err := json.Unmarshal(r.Body(), &result); err != nil {
		return nil, xerr.WithCode(xerr.ErrUnknown, err.Error())
	}

	token, ok := result["token"].(string)
	if !ok || token == "" {
		return nil, xerr.WithCode(xerr.ErrTokenInvalid, "返回token为空")
	}

	return &types.GetBusAuthResp{
		Token: token,
	}, nil
}
