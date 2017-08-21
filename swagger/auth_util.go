package swagger

import (
    "net/url"
    "time"
    "encoding/json"
    "strconv"
    "strings"
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
)

func SetApiHeader(headerParams map[string]string, c *Configuration, httpMethod, path string,
    formParams map[string]string, queryParams url.Values) {
    var expires = strconv.FormatInt(time.Now().Unix()+c.ExpireTime, 10)
    bodyStr := ""
    if len(formParams) > 1 {
        bodyBytes, _ := json.Marshal(formParams)
        bodyStr = string(bodyBytes)
    }
    //bodyStr = "{\"channelID\":2,\"message\":\"hello\"}"
    headerParams["api-expires"] = expires
    headerParams["api-key"] = c.ApiKey
    headerParams["api-signature"] = Signature(c.SecretKey, httpMethod, path[22:], queryParams.Encode(), expires, bodyStr)
}

/**
 *  nonce: nonce or expires
 */
func Signature(apiSecret, method, path, query, nonce, bodyStr string) string {
    str := ""
    if "" == query {
        str = strings.ToUpper(method) + path + nonce + bodyStr
    } else {
        str = strings.ToUpper(method) + path + "?" + query + nonce + bodyStr
    }
    return CalSignature(apiSecret, str)
}

func CalSignature(apiSecret, payload string) string {
    sig := hmac.New(sha256.New, []byte(apiSecret))
    sig.Write([]byte(payload))
    return hex.EncodeToString(sig.Sum(nil))
}
