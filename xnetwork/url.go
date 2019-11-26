package xnetwork

import "net/url"


// 解析url中的path部分
func ParsePath(httpUrl string) string {
	pUrl, err := url.Parse(httpUrl)
	if err != nil {
		panic(err)
	}

	return pUrl.Path
}


// 简单解析URL地址信息
func SimpleParseUrl(urlStr string) (hostname,path string,values map[string][]string,err error) {
	parseUrl, err := url.Parse(urlStr)
	if err != nil {
		return "","",nil,err
	}
	hostname = parseUrl.Hostname()
	path = parseUrl.Path
	values = parseUrl.Query()
	return
}