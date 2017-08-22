# aliyun-signature
阿里云请求参数签名

## 用法
`import "github.com/Fengxq2014/aliyun-signature/signature"`

### 生成get请求地址
```golang
signature.ComposeURL(inUrlValues, "testKeySecret", "http://vod.cn-shanghai.aliyuncs.com")
```

### 对url.Values排序并转义特殊字符
```golang
signature.SortQueryString(url.Values)
```

### 对string进行签名
```golang
signature.ComputeSignature(sortStr, accessSecret, "")
// 最后一个参数为空时为"GET"
```

## License ##

This library is distributed under the MIT License found in the [LICENSE](./LICENSE)
file.