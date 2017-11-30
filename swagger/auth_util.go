package swagger

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "net/url"
    "strconv"
    "strings"
    "time"
    "net/http"
)

func SetAuthHeader(request *http.Request, apiKey APIKey, c *Configuration, httpMethod, path string,
    formParams url.Values, queryParams url.Values) {
    var expires = strconv.FormatInt(time.Now().Unix()+c.ExpireTime, 10)
    bodyStr := ""
    if len(formParams) > 1 {
        bodyBytes, _ := json.Marshal(formParams)
        bodyStr = string(bodyBytes)
    }

    //bodyStr = "{\"channelID\":\"2\",\"message\":\"hello\"}"

    request.Header.Add("api-key", apiKey.Key)
    request.Header.Add("api-expires", expires)
    request.Header.Add("api-signature", Signature(apiKey.Secret, httpMethod, path[22:], queryParams.Encode(),
        expires, bodyStr))
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
