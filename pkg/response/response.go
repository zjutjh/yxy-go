package response

import (
	"net/http"
	"yxy-go/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type Response struct {
	Code xerr.Code   `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func HttpResponse(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		httpx.WriteJson(w, http.StatusOK, Success(resp))
	} else {
		code := xerr.ErrUnknown

		if e, ok := err.(*xerr.ErrCode); ok {
			code = e.Code()
		}

		logc.Errorf(r.Context(), "[HTTP] %d - %s %s - %v - %s - %s", http.StatusOK, r.Method, r.URL.Path, err, r.RemoteAddr, r.UserAgent())
		httpx.WriteJson(w, http.StatusOK, Error(code))
	}
}

func ParamErrorResponse(r *http.Request, w http.ResponseWriter, err error) {
	logc.Errorf(r.Context(), "[HTTP] %d - %s %s - %v - %s - %s", http.StatusOK, r.Method, r.URL.Path, err, r.RemoteAddr, r.UserAgent())
	httpx.WriteJson(w, http.StatusOK, Error(xerr.ErrParam))
}

func Success(data interface{}) *Response {
	code := xerr.ErrSuccess
	return &Response{code, code.String(), data}
}

func Error(code xerr.Code) *Response {
	return &Response{code, code.String(), nil}
}
