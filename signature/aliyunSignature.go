package signature

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
	"sort"
	"strings"
)

// ComposeURL 生成带签名的url
func ComposeURL(urlValues url.Values, accessSecret, hostString string) string {
	sortStr := SortQueryString(urlValues)
	Signature := ComputeSignature(sortStr, accessSecret, "")
	_url := hostString + "/?Signature=" + Signature + "&" + sortStr
	return _url
}

// ComputeSignature 生成签名
func ComputeSignature(sortQueryString, accessSecret, method string) string {
	var popBuf bytes.Buffer
	if method == "" {
		popBuf.WriteString("GET")
	} else {
		popBuf.WriteString(method)
	}
	popBuf.WriteString("&")
	popBuf.WriteString(specialURLEncode("/"))
	popBuf.WriteString("&")
	popBuf.WriteString(specialURLEncode(sortQueryString))
	return specialURLEncode(signString(popBuf.String(), accessSecret+"&"))
}

// signString 用指定的access_secret 对source进行签名
func signString(source string, accessSecret string) string {
	h := hmac.New(sha1.New, []byte(accessSecret))
	h.Write([]byte(source))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// SortQueryString 排序url.Values
func SortQueryString(preSingURL url.Values) string {
	var buffer bytes.Buffer
	keys := make([]string, 0, len(preSingURL))
	for k := range preSingURL {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if buffer.Len() > 0 {
			buffer.WriteString("&")
		}
		buffer.WriteString(specialURLEncode(k))
		buffer.WriteString("=")
		buffer.WriteString(specialURLEncode(preSingURL.Get(k)))
	}
	return buffer.String()
}

// specialURLEncode 特殊的URLEncode
func specialURLEncode(str string) string {
	str = url.QueryEscape(str)
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)
	return str
}
