package common

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestObtainMessage(t *testing.T) {
	ctx := context.Background()
	callMethod := "GET"
	endPoint := "https://stock.finance.sina.com.cn/fundInfo/api/openapi.php/CaihuiFundInfoService.getNav?symbol=004685&datefrom=2022-06-20&dateto=2022-06-28&page=1"
	header := make(map[string]string)
	header["Cookie"] = "STOCK7-FINANCE-SINA-COM-CN="
	body := []byte{}
	response, err := doRequest(ctx, callMethod, endPoint, header, body)
	if err != nil {
		fmt.Println("error...")
	}
	code := response.StatusCode
	if code != 200 {
		fmt.Println("错误码不对:", code)
	}
	b := response.Body
	defer b.Close()
	all, err := ioutil.ReadAll(b)
	if err != nil {
		fmt.Println("读取错误")
	}
	fmt.Println(string(all))

	//url := "https://stock.finance.sina.com.cn/fundInfo/api/openapi.php/CaihuiFundInfoService.getNav?symbol=004685&datefrom=2022-06-20&dateto=2022-06-28&page=1"
	//method := "GET"
	//
	//client := &http.Client{}
	//req, err := http.NewRequest(method, url, nil)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//req.Header.Add("Cookie", "STOCK7-FINANCE-SINA-COM-CN=")
	//
	//res, err := client.Do(req)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer res.Body.Close()
	//
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(body))
}
