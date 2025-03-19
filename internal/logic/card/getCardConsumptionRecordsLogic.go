package card

import (
	"context"
	"fmt"
	"time"

	"yxy-go/internal/consts"
	"yxy-go/internal/svc"
	"yxy-go/internal/types"
	"yxy-go/internal/utils/yxyClient"
	"yxy-go/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCardConsumptionRecordsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCardConsumptionRecordsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCardConsumptionRecordsLogic {
	return &GetCardConsumptionRecordsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type GetCardConsumptionRecordsYxyResp struct {
	// StatusCode int    `json:"statusCode"` // 由于响应中该字段类型不统一(int/string), 因此不解析, 改由Success判断
	Message string `json:"message"`
	Rows    []struct {
		// Type           string `json:"type"`
		Time string `json:"time"`
		// Dealtime       string `json:"dealtime"`
		Address string `json:"address"`
		// FeeName        string `json:"feeName"`
		// Serialno       string `json:"serialno"`
		Money string `json:"money"`
		// BusinessName   string `json:"businessName"`
		// BusinessNum    string `json:"businessNum"`
		// FeeNum         string `json:"feeNum"`
		// AccName        string `json:"accName"`
		// AccNum         string `json:"accNum"`
		// PerCode        string `json:"perCode"`
		// EWalletId      string `json:"eWalletId"`
		// MonCard        string `json:"monCard"`
		// AfterMon       string `json:"afterMon"`
		// ConcessionsMon string `json:"concessionsMon"`
	} `json:"rows"`
	// Total   int  `json:"total"`
	Success bool `json:"success"`
}

func (l *GetCardConsumptionRecordsLogic) GetCardConsumptionRecords(req *types.GetCardConsumptionRecordsReq) (resp *types.GetCardConsumptionRecordsResp, err error) {
	if _, err := time.Parse("20060102", req.QueryTime); err != nil {
		return nil, xerr.WithCode(xerr.ErrParam, err.Error())
	}

	yxyReq, yxyHeaders := yxyClient.GetYxyBaseReqParam(req.DeviceID)
	yxyReq["ymId"] = req.UID
	yxyReq["schoolCode"] = consts.SCHOOL_CODE
	yxyReq["queryTime"] = req.QueryTime
	yxyReq["token"] = req.Token

	var yxyResp GetCardConsumptionRecordsYxyResp
	r, err := yxyClient.HttpSendPost(consts.GET_CARD_CONSUMPTION_RECORDS_URL, yxyReq, yxyHeaders, &yxyResp)
	if err != nil {
		return nil, err
	}

	if !yxyResp.Success {
		errCode := xerr.ErrUnknown
		if yxyResp.Message == "登录已过期，请重新登录[user no find]" {
			errCode = xerr.ErrUserNotFound
		} else if yxyResp.Message == "您的账号已被登出，请重新登录[deviceId changed]" || yxyResp.Message == "登录已过期，请重新登录[token change]" {
			errCode = xerr.ErrAccountLoggedOut
		} else if yxyResp.Message == "用户还未绑卡" {
			errCode = xerr.ErrNotBindCard
		}
		return nil, xerr.WithCode(errCode, fmt.Sprintf("yxy response: %v", r))
	}

	var records []types.CardConsumptionRecord
	for _, row := range yxyResp.Rows {
		record := types.CardConsumptionRecord{
			Address: row.Address,
			Money:   row.Money,
			Time:    row.Time,
		}
		records = append(records, record)
	}

	return &types.GetCardConsumptionRecordsResp{
		List: records,
	}, nil
}
