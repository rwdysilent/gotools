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
)

var Client = &http.Client{}

func DoReq(method, url string, params url.Values, client *http.Client) (int, []byte, error) {
	if client != nil {
		Client = client
	}
	if params != nil {
		url += "?" + params.Encode()
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return 0, nil, err
	}
	resp, err := Client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, body, nil
}

func Get(url string, params url.Values, client *http.Client) (int, []byte, error) {
	return DoReq("GET", url, params, client)
}

func Post(url string, params url.Values, contentType string, body []byte, client *http.Client) (int, []byte, error) {
	if contentType == "" {
		contentType = "application/x-www-form-urlencoded;charset=utf-8"
	}

	reqBody := bytes.NewReader(body)
	resp, err := http.Post(url, contentType, reqBody)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}
	return resp.StatusCode, respBody, nil
}

func PostForm(url string, params, forms url.Values) (int, []byte, error){
	if params != nil {
		url += "?" + params.Encode()
	}
	resp, err := Client.PostForm(url, forms)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}

	return resp.StatusCode, body, nil
}
