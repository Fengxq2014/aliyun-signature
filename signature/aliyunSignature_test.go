package signature

import (
	"net/url"
	"testing"
)

func TestComputeSignature(t *testing.T) {
	want := "UI%2FwKfuvTtphzCKHwPhP0ErtLnc%3D"
	in := url.Values{
		"Action":           {"GetVideoPlayAuth"},
		"Version":          {"2017-03-21"},
		"AccessKeyId":      {"testId"},
		"Timestamp":        {"2017-03-29T12:09:11Z"},
		"SignatureMethod":  {"HMAC-SHA1"},
		"SignatureVersion": {"1.0"},
		"SignatureNonce":   {"578a50c1-280d-4a34-bffc-e06aa6b2df76"},
		"Format":           {"JSON"},
		"VideoId":          {"68a4d2629a339db3207963ac073a88cd"},
	}
	sortQueryString := SortQueryString(in)
	if result := ComputeSignature(sortQueryString, "testKeySecret", ""); result != want {
		t.Errorf("want:%v,but got:%v", want, result)
	}
}

func TestComposeURL(t *testing.T) {
	want := "http://vod.cn-shanghai.aliyuncs.com?Signature=UI%2FwKfuvTtphzCKHwPhP0ErtLnc%3D&SignatureVersion=1.0&Action=GetVideoPlayAuth&Format=JSON&VideoId=68a4d2629a339db3207963ac073a88cd&SignatureNonce=578a50c1-280d-4a34-bffc-e06aa6b2df76&Version=2017-03-21&AccessKeyId=testId&SignatureMethod=HMAC-SHA1&Timestamp=2017-03-29T12%3A09%3A11Z"
	inUrlValues := url.Values{
		"Action":           {"GetVideoPlayAuth"},
		"Version":          {"2017-03-21"},
		"AccessKeyId":      {"testId"},
		"Timestamp":        {"2017-03-29T12:09:11Z"},
		"SignatureMethod":  {"HMAC-SHA1"},
		"SignatureVersion": {"1.0"},
		"SignatureNonce":   {"578a50c1-280d-4a34-bffc-e06aa6b2df76"},
		"Format":           {"JSON"},
		"VideoId":          {"68a4d2629a339db3207963ac073a88cd"},
	}
	result := ComposeURL(inUrlValues, "testKeySecret", "http://vod.cn-shanghai.aliyuncs.com")
	gotURI,err := url.ParseRequestURI(result)
	if err != nil{
		t.Error("Not a valid url")
	}
	wantURI,_:=url.ParseRequestURI(want)
	if gotURI.Query().Encode() != wantURI.Query().Encode(){
		t.Error("not equal")
	}
}
