package yxyClient

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

// Sign 易校园签名算法
func Sign(params map[string]any) string {
	// 1. 对 payload 基于key字典序排序, 并以|链接
	keys := make([]string, 0, len(params))
	values := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v, ok := params[k].(string)
		if ok && len(v) == 0 {
			// 空字符串
			continue
		}
		values = append(values, fmt.Sprintf("%v", params[k]))
	}
	payload := strings.Join(values, "|")

	//2. hmac 加密
	hashKey := getHMACKey(params)
	mac := hmac.New(sha256.New, hashKey)
	mac.Write([]byte(payload))
	resultBytes := mac.Sum(nil)
	resultBase64 := base64.StdEncoding.EncodeToString(resultBytes)
	return resultBase64
}

// getHMACKey 获取Sign中HMAC的密钥
// 这个由 ymId 的前13位和 sha256(deviceId) 的后29位组成
func getHMACKey(params map[string]any) []byte {
	keyBuilder := ""
	if v, ok := params["ymId"]; ok {
		ymId := v.(string)
		if len(ymId) >= 13 {
			keyBuilder += ymId[len(ymId)-13:]
		}
	}
	if v, ok := params["deviceId"]; ok {
		deviceID := v.(string)
		if len(deviceID) >= 29 {
			keyBuilder += deviceID[len(deviceID)-29:]
		}
	}
	keyBytes := sha256.Sum256([]byte(keyBuilder))
	keyHex := hex.EncodeToString(keyBytes[:])
	return []byte(keyHex[32:])
}
