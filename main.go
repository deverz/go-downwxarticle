package main

import (
	"encoding/base64"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
)

/*
 * @Desc:
 * @Author: deverz@qq.com
 * @File: go_downwxarticle/main.go
 * @Date: 2022/1/18 5:19 下午
 */
func main() {
	// url链接
	//url := "https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA%3D%3D&chksm=fa80dd0dcdf7541b641be90bdf39001a8ac417c418bb19b9ee7d9e493a0627cfe7d144bb2163&idx=1&lang=zh_CN&mid=2247486618&scene=21&sn=bb5e76e011ba99ebc2ffb8f9d3c00b89&token=228666042#wechat_redirect"
	url := "https://mp.weixin.qq.com/s/5wrHQz_LqeQn3NLuF8yu0A"
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}
	// 关闭资源流
	defer res.Body.Close()

	data, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	data.Find("img").Each(func(i int, selection *goquery.Selection) {
		// 处理图片
		imgSrc, sOk := selection.Attr("data-src")
		if sOk {
			res2, err := http.Get(imgSrc)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer res2.Body.Close()
			body, err := ioutil.ReadAll(res2.Body)
			if err == nil {
				b64 := base64.StdEncoding.EncodeToString(body)
				imgB64 := fmt.Sprintf("data:image/png;base64,%s", b64)
				selection.SetAttr("src", imgB64)
			}
		}
	})
	str, _ := data.Html()
	ioutil.WriteFile("site3.html", []byte(str), 0644)
}
