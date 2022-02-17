package xhttp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

//
func dialTimeout(network, addr string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, addr, time.Second*10)
	if err != nil {
		return conn, err
	}

	tcpConn := conn.(*net.TCPConn)
	tcpConn.SetKeepAlive(false)

	return tcpConn, err
}

//post
func Post(url string, headers map[string]string, bs []byte) ([]byte, error) {
	return post(url, headers, bs)
}

//post body可以带任意结构体
func PostBody(url string, headers map[string]string, body interface{}) ([]byte, error) {
	bs, err := json.Marshal(body)
	if err != nil || bs == nil {
		return nil, err
	}

	return post(url, headers, bs)
}

//获取response body数据，并解析
func GetDataUnmarshal(url string, headers map[string]string, data interface{}) error {
	bs, status, err := get(url, headers)
	if err == nil && status == 200 && bs != nil {
		if err := json.Unmarshal(bs, &data); err != nil {
			return err
		}
		return nil
	}
	return err
}

//get
func Get(url string, headers map[string]string) ([]byte, int, error) {
	return get(url, headers)
}

//get
func get(url string, headers map[string]string) ([]byte, int, error) {
	transport := http.Transport{
		Dial:              dialTimeout,
		DisableKeepAlives: true, //禁止
	}

	client := http.Client{
		Transport: &transport,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, -1, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, -1, err
	}

	return body, resp.StatusCode, nil
}

//post
func post(url string, headers map[string]string, bs []byte) ([]byte, error) {
	//
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bs))
	if err != nil {
		return nil, err
	}

	//
	transport := http.Transport{
		Dial:              dialTimeout,
		DisableKeepAlives: true, //禁止
	}
	client := http.Client{
		Transport: &transport,
	}
	//
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
