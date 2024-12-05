package request

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var transport http.RoundTripper

const ApplicationJson = "application/json"
const ApplicationForm = "application/x-www-form-urlencoded"

// init 初始化http连接池
func init() {
	transport = http.DefaultTransport
}

// addToQuery 将参数值转换为字符串后添加到查询参数中
func addToQuery(query url.Values, key string, value interface{}) {
	switch v := value.(type) {
	case int, int64, float64:
		query.Add(key, cast.ToString(v))
	case string:
		query.Add(key, v)
	default:
		b, _ := json.Marshal(v)
		query.Add(key, string(b))
	}
}

// getRequest get请求
func getRequest(ctx context.Context, reqUrl string, reqData map[string]any, header map[string]string, timeout time.Duration) ([]byte, error) {
	client := http.Client{
		Timeout:   timeout,
		Transport: transport,
	}
	req, err := http.NewRequestWithContext(ctx, "GET", reqUrl, nil)
	if err != nil {
		logc.Errorf(ctx, "get请求初始化失败: %s", err)
		return nil, err
	}
	query := req.URL.Query()
	for k, v := range reqData {
		addToQuery(query, k, v)
	}
	req.URL.RawQuery = query.Encode()
	for hk, hv := range header {
		req.Header.Add(hk, hv)
	}
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		logc.Errorf(ctx, "get请求失败: %s", err)
		return nil, err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			logc.Errorf(ctx, "http.body资源关闭失败: %s", err)
		}
	}()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logc.Errorf(ctx, "get请求返回数据读取失败: %s", err)
		return nil, err
	}

	return respBody, nil
}

// postRequest post请求
func postRequest(ctx context.Context, reqUrl string, reqData map[string]any, header map[string]string, timeout time.Duration) ([]byte, error) {
	var (
		err  error
		data io.Reader
	)

	// 根据 header 来判断是否为 json 请求
	ct, ok := header["Content-Type"]
	if ok && ct == ApplicationJson {
		jsonData, err := json.Marshal(reqData)
		if err != nil {
			return nil, err
		}

		data = bytes.NewBuffer(jsonData)
	} else {
		params := url.Values{}
		for k, v := range reqData {
			params.Set(k, cast.ToString(v))
		}

		data = strings.NewReader(params.Encode())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", reqUrl, data)
	if err != nil {
		logc.Errorf(ctx, "post请求初始化失败: %s", err)
		return nil, err
	}
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
	for hk, hv := range header {
		req.Header.Add(hk, hv)
	}

	client := http.Client{
		Timeout:   timeout,
		Transport: transport,
	}
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		logc.Errorf(ctx, "post请求失败: %s", err)
		return nil, err
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			logc.Errorf(ctx, "http.body资源关闭失败: %s", err)
		}
	}()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logc.Errorf(ctx, "post请求返回数据读取失败: %s", err)
		return nil, err
	}

	return respBody, nil
}

// DoRequest 发起get/post请求
func DoRequest(ctx context.Context, reqUrl, method string, reqData map[string]any, header map[string]string, timeout time.Duration) ([]byte, error) {
	var resp []byte
	var err error
	if method == http.MethodGet {
		resp, err = getRequest(ctx, reqUrl, reqData, header, timeout)
	} else {
		// post 请求默认为 form 类型
		if header == nil {
			header = map[string]string{"Content-Type": ApplicationForm}
		}
		resp, err = postRequest(ctx, reqUrl, reqData, header, timeout)
	}
	var params = map[string]any{
		"url":     reqUrl,
		"method":  method,
		"data":    reqData,
		"header":  header,
		"timeout": timeout,
	}
	if err != nil {
		logc.Errorf(ctx, "接口请求失败，请求内容：%+v，返回错误：%v", params, err)
	} else {
		logc.Infof(ctx, "接口请求成功，请求内容：%+v，返回数据：%+v", params, string(resp))
	}
	return resp, err
}
