package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	CONN_TIME_OUT = time.Second * 2
)

type HttpTool struct{}

func (s *HttpTool) Get(apiURL string, params url.Values) (resData string, e error) {
	var (
		Url *url.URL
		err error
	)
	Url, err = url.Parse(apiURL)
	if err != nil {
		return "", err
	}
	Url.RawQuery = params.Encode()
	resp, err := http.Get(Url.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

/*
*
Get with TimeOut
Code == 200 则返回,其他请打印错误
第三个参数设置超时  单位:秒 【0:则默认两秒】
*/
func (s *HttpTool) SendGetWithTimeOut(apiUrl string, params url.Values, time_out int) (resData string, e error) {
	var (
		Url *url.URL
		err error
		b   []byte
	)
	Url, err = url.Parse(apiUrl)
	if err != nil {
		return "", err
	}
	Url.RawQuery = params.Encode()
	client := _httpClient(time_out)
	resp, err := client.Get(Url.String())
	if err != nil || resp == nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", nil
	}
	if resp.Body != nil {
		b, err = ioutil.ReadAll(resp.Body)
		return string(b), nil
	}
	return "", err
}

func _httpClient(time_out int) http.Client {
	var (
		_sec time.Duration
	)
	_sec = CONN_TIME_OUT
	if time_out > 0 {
		_sec = time.Second * time.Duration(time_out)
	}
	return http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, _sec)
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(_sec))
				return conn, nil
			},
			ResponseHeaderTimeout: _sec,
		},
	}
}

/*
*
网络请求POST
*/
func (s *HttpTool) Post(apiURL string, params url.Values) (resData string, err error) {
	resp, err := http.PostForm(apiURL, params)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

/*
*
Post with TimeOut
Code == 200 则返回,其他请打印错误
第三个参数设置超时  单位:秒 【0:则默认两秒】
*/
func (s *HttpTool) SendPostWithTimeOut(apiUrl string, params url.Values, time_out int) (resData string, e error) {
	var (
		err error
		b   []byte
	)
	client := _httpClient(time_out)
	resp, err := client.PostForm(apiUrl, params)
	if err != nil || resp == nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", nil
	}
	if resp.Body != nil {
		b, err = ioutil.ReadAll(resp.Body)
		return string(b), nil
	}
	return "", err
}

// OnPostJSON 发送修改密码
func (s *HttpTool) OnPostJSON(url, jsonstr string) []byte {
	//解析这个 URL 并确保解析没有出错。
	body := bytes.NewBuffer([]byte(jsonstr))
	resp, err := http.Post(url, "application/json;charset=utf-8", body)
	if err != nil {
		return []byte("")
	}
	defer resp.Body.Close()
	body1, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		return []byte("")
	}

	return body1
}

// OnGetJSON 发送get 请求
func (s *HttpTool) OnGetJSON(url, params string) string {
	//解析这个 URL 并确保解析没有出错。
	var urls = url
	if len(params) > 0 {
		urls += "?" + params
	}
	resp, err := http.Get(urls)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body1, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		return ""
	}

	return string(body1)
}

// SendGet 发送get 请求 返回对象
func (s *HttpTool) SendGet(url, params string, obj interface{}) bool {
	//解析这个 URL 并确保解析没有出错。
	var urls = url
	if len(params) > 0 {
		urls += "?" + params
	}
	resp, err := http.Get(urls)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	//log.Println((string(body)))
	err = json.Unmarshal([]byte(body), &obj)
	if err != nil {
		return false
	}

	return true
}

// SendGetEx 发送GET请求
func (s *HttpTool) SendGetEx(url string, reponse interface{}) bool {
	resp, e := http.Get(url)
	if e != nil {
		return false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	err = json.Unmarshal(body, &reponse)
	if err != nil {
		return false
	}

	return true
}

// OnPostForm form 方式发送post请求
func (s *HttpTool) OnPostForm(url string, data url.Values) (body []byte) {
	resp, err := http.PostForm(url, data)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}

// SendPost 发送POST请求
func (s *HttpTool) SendPost(requestBody interface{}, responseBody interface{}, url string) bool {
	postData, err := json.Marshal(requestBody)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewReader(postData))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	//	req.Header.Add("Authorization", authorization)
	resp, e := client.Do(req)
	if e != nil {
		return false
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	//	result := string(body)

	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return false
	}

	return true
}

// WriteJSON  像指定client 发送json 包
// msg message.MessageBody
func (s *HttpTool) WriteJSON(w http.ResponseWriter, msg interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	js, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(js))
}
