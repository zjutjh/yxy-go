package yxyClient

import (
	"testing"
)

func TestSign(t *testing.T) {
	tests := []struct {
		params map[string]any
		want   string
	}{
		{
			map[string]any{
				"appVersion": "730",
				"sceneCode":  "app_user_login",
				"nt":         1756225768826,
				"deviceId":   "ym-57766d45173c643469eadb10a8766666",
				"platform":   "YUNMA_APP",
				"schoolCode": "10337",
			}, "i85zWFEF0nb8BR8VxQGw7IqkZvT9C5CUqyt6KteO+4U=",
		},
		{
			map[string]any{
				"appVersion": "730",
				"nt":         1756228132412,
				"mt":         1756228132115,
				"deviceId":   "ym-57766d45173c643469eadb10a8766666",
				"platform":   "YUNMA_APP",
			}, "dv2UcggaYmdhK73jlGxZhF4EdvGnR1UW6g8TnZIU7RM=",
		},
		{
			map[string]any{
				"appVersion":       "730",
				"clientId":         "65l3attk4r095ib",
				"appAllVersion":    "7.3.6",
				"appPlatform":      "Android",
				"nt":               1756229573789,
				"mobileType":       "23013RK75C",
				"deviceId":         "ym-bcbcbb20801444c96978b0c72f114514",
				"platform":         "YUNMA_APP",
				"verificationCode": "123456",
				"mobilePhone":      "13355225522",
				"osVersion":        "15",
				"osType":           "Android",
				"brand":            "Redmi",
				"invitationCode":   "",
				"schoolCode":       "",
			}, "eR/TDASY3yy8AlHlhqjDRK7/FrVclhp4T2xyyIyq2qI=",
		},
		{
			map[string]any{
				"appVersion": "730",
				"sceneCode":  "app_user_login",
				"nt":         1756230992802,
				"deviceId":   "ym-bcbcbb20801444c96978b0c72f114514",
				"platform":   "YUNMA_APP",
				"schoolCode": "",
			}, "kXAbBGxdQjpl5kxiYDmVD76nOAbQ7PImqECdy4m7K2k=",
		},
		{
			map[string]any{
				"securityToken": "6d023e2660e76aafd16a7e93c8295bc3aZjutjhwlwCnpH706AwSfQ==",
				"appVersion":    "730",
				"nt":            1756230993220,
				"deviceId":      "ym-bcbcbb20801444c96978b0c72f114514",
				"platform":      "YUNMA_APP",
				"schoolCode":    "",
			}, "ME4fuAlO9zF0I1PanCUdddLbMaVXFOG5unjnw+j3+qo="},
		{
			map[string]any{
				"securityToken":     "6d023e2660e76aafd16a7e93c8295bc3aZjutjhwlwCnpH706AwSfQ==",
				"appVersion":        "730",
				"mobilePhone":       "13355225522",
				"nt":                1756231211379,
				"appSecurityToken":  "/G/IJOQnR7sjohNq/JFtVhl07BNq797nleNaFfZZDB29HkAhPmDNjGf6wNcA0C3HfZJ12345B84TH3jDC1VDjcUuONlBsk/A15Z1co+buQXOavNWEhwrSJgFDDycYBT5xiGdYB/6EVa7/o5/CJmDLoU1ikYKbrD9EY69juphqq4=",
				"imageCaptchaValue": "y4pm",
				"sendCount":         1,
				"deviceId":          "ym-bcbcbb20801444c96978b0c72f114514",
				"platform":          "YUNMA_APP",
				"schoolCode":        "",
			}, "KJrBJlLwNZcXXLmHtcHO6LR+tCsxwXUd9yF6ChUmpGU=",
		},
		{
			map[string]any{
				"appVersion":       "730",
				"clientId":         "65l3attk4r095ib",
				"appAllVersion":    "7.3.6",
				"appPlatform":      "Android",
				"nt":               1756231286750,
				"mobileType":       "23013RK75C",
				"deviceId":         "ym-bcbcbb20801444c96978b0c72f114514",
				"platform":         "YUNMA_APP",
				"verificationCode": "123456",
				"mobilePhone":      "13355225522",
				"osVersion":        "15",
				"osType":           "Android",
				"brand":            "Redmi",
				"invitationCode":   "",
				"schoolCode":       "",
			}, "chMojn114ebG3PR1BmiYz7YLu2AzeING+guiJs7cds8=",
		},
		{
			map[string]any{
				"appVersion": "730",
				"nt":         1756730606020,
				"ymId":       "2408157831570432666",
				"deviceId":   "ym-bcbcbb20801444c96978b0c72f114514",
				"platform":   "YUNMA_APP",
			}, "9buzHakN0foXRu30s7CyB3B8zH9epcBe5vB+zaBC4s4=",
		},
		{
			map[string]any{
				"appVersion":    "730",
				"clientId":      "65l3ben7ub3x6vs",
				"appAllVersion": "7.3.6",
				"appPlatform":   "Android",
				"nt":            1756734241205,
				"mobileType":    "23013RK75C",
				"deviceId":      "ym-bcbcbb20801444c96978b0c72f114514",
				"platform":      "YUNMA_APP",
				"osVersion":     "15",
				"ymId":          "2408157831570432666",
				"osType":        "Android",
				"brand":         "Redmi",
				"schoolCode":    "10337",
			}, "eDI0hQQ673SN/4m9dTpixTLYpsSm7VgnUtWXXQd5vmw=",
		},
	}

	for _, tt := range tests {
		got := Sign(tt.params)
		if got != tt.want {
			t.Errorf("[Error] sign(%v) = %s; want %s", tt.params, got, tt.want)
		}
	}
}
