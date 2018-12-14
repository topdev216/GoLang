package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

const (
	debug   = true
	https   = false
	app_key = ""
	secret  = ""
	session = ""
)

func encode(src string) string {
	data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewEncoder()))
	if err != nil {
		log.Fatalln(err)
	}
	return string(data)
}

// decode gbk to utf8
func decode(src string) string {
	data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewDecoder()))
	if err != nil {
		log.Fatalln(err)
	}
	return string(data)
}

func postResp(strUrl string, postDict map[string]string, method string) map[string]interface{} {
	var postBytesReader *bytes.Reader

	// form request params
	if postDict != nil {
		postValues := url.Values{}
		for postKey, postValue := range postDict {
			postValues.Set(postKey, postValue)
		}
		postBytesReader = bytes.NewReader([]byte(postValues.Encode()))
	}

	// set post body
	var httpReq *http.Request
	httpReq, _ = http.NewRequest(method, strUrl, postBytesReader)
	httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// do post
	httpClient := &http.Client{}
	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		log.Fatalln(err)
	}
	defer httpResp.Body.Close()

	// parse response data
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var f map[string]interface{}
	if err := json.Unmarshal([]byte(decode(string(body))), &f); err != nil {
		return nil
	}
	return f
}

func getSigned(param map[string]string) string {
	// sort
	var keys []string
	for k := range param {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// join
	var str string
	for _, k := range keys {
		str += fmt.Sprintf("%s%s", k, param[k])
	}

	// encrypt
	x := md5.New()
	x.Write([]byte(secret + str + secret))
	return hex.EncodeToString(x.Sum(nil))
}

func getTaobaoApi(debug bool, https bool) string {
	var uri string
	if https {
		uri = "https://eco.taobao.com/router/rest"
		if debug {
			uri = "https://gw.api.tbsandbox.com/router/rest"
		}
	} else {
		uri = "http://gw.api.taobao.com/router/rest"
		if debug {
			uri = "http://gw.api.tbsandbox.com/router/rest"
		}
	}
	return uri
}

func TaobaoRequest(method string, need_auth bool, param map[string]string) map[string]interface{} {
	data := map[string]string{}

	data["method"] = method
	data["app_key"] = app_key
	if need_auth {
		data["session"] = session
	}
	data["timestamp"] = time.Now().Format("2006-01-02 15:04:05") // 时间戳，格式为yyyy-MM-dd HH:mm:ss，时区为GMT+8，例如：2015-01-01 12:00:00。淘宝API服务端允许客户端请求最大时间误差为10分钟。
	data["format"] = "json"                                      // 响应格式。默认为xml格式，可选值：xml，json。
	data["v"] = "2.0"                                            // API协议版本，可选值：2.0。
	data["partner_id"] = ""                                      // 合作伙伴身份标识。
	data["target_app_key"] = ""                                  // 被调用的目标AppKey，仅当被调用的API为第三方ISV提供时有效。
	data["simplify"] = ""                                        // 是否采用精简JSON返回格式，仅当format=json时有效，默认值为：false。
	data["sign_method"] = "md5"                                  // 签名的摘要算法，可选值为：hmac，md5。

	// 拼接上接口附加的参数
	for k, v := range param {
		data[k] = v
	}

	data["sign"] = getSigned(data)

	return postResp(getTaobaoApi(debug, https), data, "POST")
}

/**
 * @desc taobao.shop.getbytitle (根据店铺名称获取店铺信息)
 * @param param
 *    title 店铺名称，必须严格匹配（不支持模糊查询）
 *    fields sid,cid,title,nick,desc,bulletin,pic_path,created,modified,shop_score
 */
func Shop_GetByTitle(param map[string]string) map[string]interface{} {
	return TaobaoRequest("taobao.shop.getbytitle", false, param)
}

func main() {
	params := map[string]string{}
	params["title"] = "XXXX官方旗舰店"
	params["fields"] = "sid,cid,title,nick,desc,bulletin,pic_path,created,modified,shop_score"
	log.Println(Shop_GetByTitle(params))
}
