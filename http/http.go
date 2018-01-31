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

var (
	defaultClient = &http.Client{}
	defaultType   = "application/x-www-form-urlencoded"
)

//DoReq to request the url
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

//WGet to get data
func WGet(url string, params url.Values, client *http.Client) (int, []byte, error) {
	return DoReq("GET", url, "", params, client, nil)
}

//WPost to post data
func WPost(url string, params url.Values, contentType string, body []byte, client *http.Client) (
	int, []byte, error) {
	if contentType == "" {
		contentType = defaultType
	}
	return DoReq("POST", url, contentType, params, client, bytes.NewReader(body))
}

//WPostForm to Post form data
func WPostForm(url string, forms url.Values) (int, []byte, error) {
	return DoReq("POST", url, defaultType, nil, nil, strings.NewReader(forms.Encode()))
}
