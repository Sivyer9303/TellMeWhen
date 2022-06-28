package common

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const secret = "SEC8ecb6c95842ae90335557f4cb974a7fbc836617f4123b50b6be82c42afa8dff1"
const path = "https://oapi.dingtalk.com/robot/send"
const token = "d9770adfe5f2e9a4fba80a989c71bee821703d83e80776854054ed91d4876e71"

func getSign() string {
	// https://oapi.dingtalk.com/robot/send?access_token=d9770adfe5f2e9a4fba80a989c71bee821703d83e80776854054ed91d4876e71
	// 1656406786128
	// 1656406830
	// 1656406849850203500
	// 1656406890000
	ts := time.Now().UnixNano() / 1e6
	fmt.Println(ts)
	strToHash := fmt.Sprintf("%d\n%s", ts, secret)
	hmac256 := hmac.New(sha256.New, []byte(secret))
	hmac256.Write([]byte(strToHash))
	data := hmac256.Sum(nil)
	return base64.StdEncoding.EncodeToString(data)
}

func SendMessage(msg string) error {
	var (
		ctx    context.Context
		cancel context.CancelFunc
		uri    string
		resp   *http.Response
		err    error
	)
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	value := url.Values{}
	value.Set("access_token", token)
	if secret != "" {
		t := time.Now().UnixNano() / 1e6
		value.Set("timestamp", fmt.Sprintf("%d", t))
		value.Set("sign", getSign())

	}
	uri = path + value.Encode()
	header := map[string]string{
		"Content-type": "application/json",
	}
	marshal, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("序列化数据失败")
		return err
	}
	resp, err = doRequest(ctx, "POST", uri, header, marshal)

	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("send msg err: %s, token: %s, msg: %s", string(body), token, marshal)
	}
	return nil
}
