package HttpSmallClient

import (
	"errors"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

var RequestType = struct {
	Get    string
	Post   string
	Delete string
}{
	Get:    "GET",
	Post:   "POST",
	Delete: "DELETE",
}

//func (c *http.Client) Post(url, contentType string, body io.Reader) (resp *Response, err error) {
//	req, err := NewRequest("POST", url, body)
//	if err != nil {
//		return nil, err
//	}
//	req.Header.Set("Content-Type", contentType)
//	return c.Do(req)
//}

func NewClient(timeout time.Duration) (client *http.Client, err error) {
	client = &http.Client{

		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, timeout)
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(timeout))
				return conn, nil
			},

			ResponseHeaderTimeout: timeout,
		},
	}

	//client = &http.Client{
	//	Transport: &http.Transport{
	//		DialContext: (&net.Dialer{
	//			Timeout:   timeout, //连接超时时间
	//			KeepAlive: timeout, //连接保持超时时间
	//			DualStack: true,    //
	//		}).DialContext,
	//		MaxIdleConnsPerHost:   2,       //每个host 最大空闲链接数
	//		MaxIdleConns:          100,     //client对与所有host最大空闲连接数总和
	//		IdleConnTimeout:       timeout, //空闲连接在连接池中的超时时间
	//		TLSHandshakeTimeout:   timeout, //TLS安全连接握手超时时间
	//		ExpectContinueTimeout: timeout, //发送完请求到接收到响应头的超时时间
	//	},
	//	//总的超时时间
	//	Timeout: timeout,
	//}
	return
}

var header sync.Map
var timeout = time.Duration(3) * time.Second
var retryTimes = 3

func SetHeader(key string, value string) {
	header.Store(key, value)
}

func GetHeader(key string) (value string, err error) {
	v, ok := header.Load(key)
	if true != ok {
		return "", errors.New("key not exists")
	}
	if nil == v {
		return "", errors.New("value is empty")
	}
	return v.(string), nil
}

func DelHeader(key string) {
	header.Delete(key)
}

func DellAllHeader() {
	header.Range(func(key, value interface{}) bool {
		header.Delete(key)
		return true
	})
}

func SetTimeOut(_timeout time.Duration) {
	timeout = _timeout
}

func GetTimeOut() time.Duration {
	return timeout
}

func SetRetryTimes(_retryTimes int) {
	retryTimes = _retryTimes
}
func GetRetryTimes() int {
	return retryTimes
}

func DoPost(_url string, data string) (body []byte, err error) {
	return doRequest(RequestType.Post, timeout, retryTimes, _url, data)
}

func DoGet(_url string) (body []byte, err error) {
	return doRequest(RequestType.Get, timeout, retryTimes, _url, "")
}

func DoDelete(_url string) (body []byte, err error) {
	return doRequest(RequestType.Delete, timeout, retryTimes, _url, "")
}

func doRequest(reqType string, timeout time.Duration, retryTimes int, _url string, data string) (body []byte, err error) {
DO_RETRY:
	client, err := NewClient(timeout)
	if nil != err {
		return nil, err
	}

	req, err := http.NewRequest(reqType, _url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	header.Range(func(key, value interface{}) bool {
		req.Header.Set(key.(string), value.(string))
		return true
	})

	resp, err := client.Do(req)
	if nil != err {
		retryTimes = retryTimes - 1
		log.Println("Http-doRequest-failed:", timeout, retryTimes, _url, err)
		if 0 < retryTimes {
			log.Println("Http-doRequest-doretry:", retryTimes)
			goto DO_RETRY
		}
		return nil, err
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if nil != err {
		return nil, err
	}

	return
}

func DoPostForm(timeout time.Duration, retryTimes int, url string, data url.Values) (body []byte, err error) {
DO_RETRY:
	client, err := NewClient(timeout)
	if nil != err {
		return nil, err
	}
	resp, err := client.PostForm(url, data)
	if nil != err {
		retryTimes = retryTimes - 1
		log.Println("HttpPostForm-failed:", url, err)
		if 0 < retryTimes {
			goto DO_RETRY
		}
		return nil, err
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if nil != err {
		return nil, err
	}

	return

}

func GetTCN2Long(timeout time.Duration, url string) (LongURL string, urlList []map[int]string, isSuccess bool, err error) {

	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, timeout)
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(timeout))
				return conn, nil
			},
			ResponseHeaderTimeout: timeout,
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			//log.Println(os.Getpid(), "CheckRedirect:%#v", req.Response.StatusCode, req.URL.String())

			stat := make(map[int]string)
			stat[req.Response.StatusCode] = req.URL.String()
			urlList = append(urlList, stat)
			LongURL = req.URL.String()

			if !strings.Contains(req.URL.Host, "weibo.cn") &&
				!strings.Contains(req.URL.Host, "weibo.com") &&
				!strings.Contains(req.URL.Host, "sina.cn") &&
				!strings.Contains(req.URL.Host, "sina.com") {
				return errors.New("success-jump-out")
			}
			return nil
		},
	}
	stat := make(map[int]string)
	stat[0] = url
	urlList = append(urlList, stat)

	_, err = client.Get(url)
	isSuccess = false
	if nil != err {
		if strings.Contains(err.Error(), "success-jump-out") {
			isSuccess = true
		}
	} else {
		isSuccess = true
	}

	//log.Println(err,resp)
	//log.Println(LongURL)
	//log.Println(urlList)
	return
}

func GetRealTargetURL(timeout time.Duration, url string) (realTargetURL string, urlList []map[int]string, err error) {

	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, timeout)
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(timeout))
				return conn, nil
			},
			ResponseHeaderTimeout: timeout,
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			log.Println(os.Getpid(), "CheckRedirect:%#v", req.Response.StatusCode, req.URL.String())
			stat := make(map[int]string)
			stat[req.Response.StatusCode] = req.URL.String()
			urlList = append(urlList, stat)
			return nil
		},
	}
	stat := make(map[int]string)
	stat[302] = url
	urlList = append(urlList, stat)

	resp, err := client.Get(url)

	if nil == err {
		_lastURL := resp.Request.URL.String()

		if nil != err {
			return "", []map[int]string{}, err
		}
		if _lastURL == url {
			return url, []map[int]string{map[int]string{resp.StatusCode: url}}, nil
		}
		return _lastURL, urlList, nil
	} else {
		return "", urlList, err
	}

	return
}
