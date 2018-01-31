// Copyright 2018 TED@Sogou, Inc. All rights reserved.
//
// @Author: wupengfei@sogou-inc.com
// @Date: 2018-01-30 19:09

package http

import (
	"net/http"
	"io/ioutil"
	"net/url"
	"bytes"
	"io"
	"strings"
)

var defaultClient = &http.Client{}

func DoReq(method, url, contentType string, params url.Values, client *http.Client, repBody io.Reader) (
	int, []byte, error) {
	if client != nil {
		defaultClient = client
	}
	if params != nil {
		url += "?" + params.Encode()
	}
	req, err := http.NewRequest(method, url, repBody)
	if err != nil {
		return 0, nil, err
	}

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	resp, err := defaultClient.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, body, nil
}

func Get(url string, params url.Values, client *http.Client) (int, []byte, error) {
	return DoReq("GET", url, "", params, client, nil)
}

func Post(url string, params url.Values, contentType string, body []byte, client *http.Client) (int, []byte, error) {
	if contentType == "" {
		contentType = "application/x-www-form-urlencoded"
	}
	reqBody := bytes.NewReader(body)
	return DoReq("POST", url, contentType, params, client, reqBody)
	//resp, err := http.Post(url, contentType, reqBody)
	//if err != nil {
	//	return 0, nil, err
	//}
	//defer resp.Body.Close()
	//
	//respBody, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return resp.StatusCode, nil, err
	//}
	//return resp.StatusCode, respBody, nil
}

func PostForm(url string, forms url.Values) (int, []byte, error) {
	contentType := "application/x-www-form-urlencoded"
	reqData := strings.NewReader(forms.Encode())

	return DoReq("POST", url, contentType, nil, nil, reqData)
	//resp, err := defaultClient.PostForm(url, forms)
	//if err != nil {
	//	return 0, nil, err
	//}
	//defer resp.Body.Close()
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return resp.StatusCode, nil, err
	//}
	//
	//return resp.StatusCode, body, nil
}
