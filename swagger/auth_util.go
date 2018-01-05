package swagger

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "net/http"
    "net/http/httputil"
    "net/url"
    "strconv"
    "strings"
    "time"
)

func SetAuthHeader(request *http.Request, apiKey APIKey, c *Configuration, httpMethod, path, postBody string,
    queryParams url.Values) {
    var expires = strconv.FormatInt(time.Now().Unix()+c.ExpireTime, 10)
    request.Header.Add("api-key", apiKey.Key)
    request.Header.Add("api-expires", expires)
    request.Header.Add("api-signature", Signature(apiKey.Secret, httpMethod, path[22:], queryParams.Encode(),
        expires, postBody))
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

// Save a copy of this request for debugging.
func DebugHttpRequest(r *http.Request) {
    requestDump, err := httputil.DumpRequest(r, true)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(requestDump))
}
