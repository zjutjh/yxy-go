package authManager

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
	"yxy-go/internal/consts"
	"yxy-go/internal/svc"
	"yxy-go/internal/utils/yxyClient"
	"yxy-go/pkg/xerr"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
)

type getAuthTokenResp struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
	Success    bool   `json:"success"`
}

type AuthManager struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthManager(ctx context.Context, svcCtx *svc.ServiceContext) *AuthManager {
	return &AuthManager{
		ctx:    ctx,
		Logger: logx.WithContext(ctx),
		svcCtx: svcCtx,
	}
}

// fetchAuthToken 发送请求获取AuthToken
func (l *AuthManager) fetchAuthToken(uid string) (string, error) {
	authReq, yxyHeaders := yxyClient.GetYxyBaseReqParam("")
	authReq["ymId"] = uid
	var authResp getAuthTokenResp
	r, err := yxyClient.HttpSendPost(consts.GET_AUTH_TOKEN, authReq, yxyHeaders, &authResp)
	ymAuthToken, ok := authResp.Data.(string)
	if err != nil || !ok || ymAuthToken == "" {
		l.Errorf("yxyClient.HttpSendPost err: %v , [%s]", err, consts.GET_AUTH_TOKEN)
		return "", xerr.WithCode(xerr.ErrHttpClient, err.Error())
	}
	yxyReq := map[string]string{
		"bindSkip":    "1",
		"authType":    "2",
		"ymAppId":     "1810181825222034",
		"callbackUrl": consts.APPLICATION_URL + "/",
		"unionid":     uid,
		"schoolCode":  consts.SCHOOL_CODE,
		"ymAuthToken": ymAuthToken,
	}

	client := yxyClient.GetClient()
	r, err = client.R().
		SetHeaders(yxyHeaders).
		SetQueryParams(yxyReq).
		Get(consts.GET_AUTH_CODE_URL)
	if r == nil || (err != nil && r.StatusCode() != 302) {
		l.Errorf("yxyClient.HttpSendPost err: %v , [%s]", err, consts.GET_AUTH_TOKEN)
		return "", xerr.WithCode(xerr.ErrHttpClient, err.Error())
	}

	location := r.Header().Get("Location")
	if location == "" {
		if strings.Contains(r.String(), "用户不存在") {
			return "", xerr.WithCode(xerr.ErrUserNotFound, fmt.Sprintf("User not found, UID: %v", uid))
		}
		return "", xerr.WithCode(xerr.ErrUnknown, fmt.Sprintf("yxy response: %v", r))
	}
	// hack 掉路由 hash模式 下url中的 /#/ 便于 query 参数提取
	location = strings.ReplaceAll(location, "#/", "")
	parsedURL, _ := url.Parse(location)
	ymCode := parsedURL.Query().Get("ymCode")

	r, err = yxyClient.HttpSendPost(consts.GET_AUTH_TOKEN_URL,
		map[string]interface{}{
			"authType": "2",
			"code":     ymCode,
		}, yxyHeaders, &authResp)
	if err != nil {
		return "", xerr.WithCode(xerr.ErrUnknown, fmt.Sprintf("yxy response: %v", r))
	}

	if authResp.StatusCode != 0 {
		return "", xerr.WithCode(xerr.ErrUnknown, fmt.Sprintf("yxy response: %v", r))
	}
	var shiroJID string
	for _, cookie := range r.Cookies() {
		if cookie.Name == "shiroJID" {
			shiroJID = cookie.Value
			// 这里不break是因为会有多个重复的 shiroJID 要拿到最后一个
			// break
		}
	}
	return shiroJID, nil
}

const cacheTTL = 24 * time.Hour

// getCacheKey 获取缓存token的key
func getCacheKey(uid string) string {
	// TODO 考虑后续修改为 auth_token:uid, 因为这个auth目前看来与业务无关, 可以通用
	return "elec:auth_token:" + uid
}

// refreshCachedAuthToken 刷新缓存中的AuthToken
func (l *AuthManager) refreshCachedAuthToken(uid string) (string, error) {
	token, err := l.fetchAuthToken(uid)
	if err != nil {
		return "", err
	}
	key := getCacheKey(uid)
	l.svcCtx.Rdb.Set(l.ctx, key, token, cacheTTL)
	return token, nil
}

// getCachedAuthToken 获取authToken, 优先从缓存中获取
func (l *AuthManager) getCachedAuthToken(uid string) (string, error) {
	key := getCacheKey(uid)
	token, err := l.svcCtx.Rdb.Get(l.ctx, key).Result()
	if err == nil {
		return token, nil
	}

	if errors.Is(err, redis.Nil) {
		return l.refreshCachedAuthToken(uid)
	} else {
		return "", xerr.WithCode(xerr.ErrUnknown, err.Error())
	}
}

type AuthHandler func(token string) (any, error)

// WithAuthToken 包装需要使用token的业务函数, 只需要将其作为回调传入, 以下处理函数会自动处理token的获取和缓存, 并将token注入业务函数
func (l *AuthManager) WithAuthToken(uid string, fn AuthHandler) (any, error) {
	// 1. 从缓存获取 token
	token, err := l.getCachedAuthToken(uid)
	if err != nil {
		return nil, xerr.WithCode(xerr.ErrUnknown, err.Error())
	}

	// 2. 调用回调函数
	result, err := fn(token)
	if err == nil {
		return result, nil
	}

	// 3. token 失效
	l.Logger.Errorf("token: %s 失效, 刷新token", token)
	if token, err = l.refreshCachedAuthToken(uid); err != nil {
		return nil, xerr.WithCode(xerr.ErrUnknown, err.Error())
	}

	return fn(token)
}
