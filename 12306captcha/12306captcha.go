package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	router "github.com/gorilla/mux"
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
	<div><img src="%s" alt=""></div>
	<br>
	%s
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
	var port string
	if len(os.Args) >= 2 && os.Args[1] != "" {
		port = os.Args[1]
	} else {
		fmt.Println("没有指定端口，默认启动:12306")
		port = "12306"
	}
	url := fmt.Sprintf("%s%s", popup_passport_captcha, strconv.FormatInt(time.Now().UnixNano()/1e6, 10))

	mux := router.NewRouter()
	// mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.String())
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
		img, _, _ := ReadImageFromString(captcha12306.Image)
		pics := cut12306(img)
		var imglist string
		for _, v := range pics {
			bf := bytes.NewBuffer(nil)
			jpeg.Encode(bf, v, &jpeg.Options{Quality: 80})
			imglist += fmt.Sprintf(`<img src=data:image/jpg;base64,%s><br>`, base64.StdEncoding.EncodeToString(bf.Bytes()))
		}

		io.WriteString(w, fmt.Sprintf(tpl, attr+captcha12306.Image, imglist))
	})

	log.Fatalln(http.ListenAndServe(":"+port, mux))
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
