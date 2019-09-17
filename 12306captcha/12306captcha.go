package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	// "github.com/Sakurasan/to"
)

// "https://kyfw.12306.cn/passport/captcha/captcha-image64?login_site=E&module=login&rand=sjrand&1568713748079"
var (
	attr                   = "data:image/jpg;base64,"
	popup_passport_captcha = "https://kyfw.12306.cn/passport/captcha/captcha-image64?login_site=E&module=login&rand=sjrand&"
)

const tpl = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>12306</title>
</head>
<body>
    <div class="pacman"></div>
    <img src="%s" alt="">
</body>
</html>
`

type captcha struct {
	Image         string `json:"image"`
	ResultMessage string `json:"result_message"`
	ResultCode    string `json:"result_code"`
}

func (t *captcha) fmtstruct(data []byte) error {
	return json.Unmarshal(data, t)

}

func main() {
	url := fmt.Sprintf("%s%s", popup_passport_captcha, strconv.FormatInt(time.Now().UnixNano()/1e6, 10))

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		captcha12306 := captcha{}
		data, err := DoGet(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = json.Unmarshal(data, &captcha12306)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println("看这里:", captcha12306.Image)
		//io.WriteString(w, attr+captcha12306.Image)
		io.WriteString(w, fmt.Sprintf(tpl, attr+captcha12306.Image))
	})

	http.ListenAndServe(":8080", mux)
}

func DoGet(urlstr string) ([]byte, error) {
	resp, err := http.Get(urlstr)
	if err != nil {
		log.Fatalf("HTTP通信，Get 请求出错[%s]", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("HTTP通信，Get 请求出错")
	}
	//fmt.Println(string(body))
	return body, nil
}
