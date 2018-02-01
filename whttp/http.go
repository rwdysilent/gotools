// Copyright 2018 TED@Sogou, Inc. All rights reserved.
//
// @Author: wupengfei@sogou-inc.com
// @Date: 2018-01-30 19:09

package whttp

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	http.Client
}

var (
	defaultClient = &Client{}
	defaultType   = "application/x-www-form-urlencoded"
)

//DoReq to request the url
func (c *Client) DoReq(method, url, contentType string, params url.Values, repBody io.Reader) (
	int, []byte, error) {

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
	resp, err := c.Client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, body, nil
}

//WGet to get data
func WGet(url string, params url.Values) (int, []byte, error) {
	return defaultClient.WGet(url, params)
}

func (c *Client) WGet(url string, params url.Values) (int, []byte, error) {
	return c.DoReq("GET", url, "", params, nil)
}

//WPost to post data
func WPost(url string, contentType string, body []byte) (int, []byte, error) {
	return defaultClient.WPost(url, contentType, body)
}

func (c *Client) WPost(url string, contentType string, body []byte) (int, []byte, error) {
	if contentType == "" {
		contentType = defaultType
	}
	return c.DoReq("POST", url, contentType, nil, bytes.NewReader(body))
}

//WPostForm to Post form data
func WPostForm(url string, forms url.Values) (int, []byte, error) {
	return defaultClient.WPostForm(url, forms)
}

func (c *Client) WPostForm(url string, forms url.Values) (int, []byte, error) {
	return c.DoReq("POST", url, defaultType, nil, strings.NewReader(forms.Encode()))
}
