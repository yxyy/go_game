package nets

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"context"
)

func Get(urls string) ([]byte,error) {
	return get(urls)
}

func get(urls string) ([]byte,error) {
	resp, err := http.Get(urls)
	if err !=nil {
		return nil,err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func do(method string, urls string) (*http.Response,error) {
	client := http.Client{}
	parseUrl, err := url.Parse(urls)
	if err!=nil {
		return nil,err
	}
	request := &http.Request{
		Method:method,
		URL:parseUrl,
	}

	request.Header["sign"] = []string{}

	return  client.Do(request)
}

func Post(ctx context.Context,urls string , contentType string , data map[string]interface{}) ([]byte,error) {

	tick := strconv.Itoa(int(time.Now().Unix()))
	data["tick"] = tick
	data["sign"] = fmt.Sprintf("%x",md5.Sum([]byte(viper.GetString("log.secret_key")+ tick)))
	body := &bytes.Reader{}
	switch contentType {
	case "application/json":
		jsonbyte, err := json.Marshal(data)
		if err!=nil {
			return nil,err
		}
		body = bytes.NewReader(jsonbyte)
	case "application/x-www-form-urlencoded":
		values := url.Values{}
		for k,v := range data {
			switch i := v.(type) {
			case string:
				values.Set(k,i)
			case int:
				values.Set(k,strconv.Itoa(i))
			case int64:
				values.Set(k,strconv.Itoa(int(i)))
			case float32:
				values.Set(k,strconv.Itoa(int(i)))
			case float64:
				values.Set(k,strconv.Itoa(int(i)))
			}

		}
		body = bytes.NewReader([]byte(values.Encode()))
	}
	resp, err := http.Post(urls, contentType, body)
	if err!=nil {
		return nil,err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}