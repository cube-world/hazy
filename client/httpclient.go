// Copyright 2022 The Yusha Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client

import (
	"errors"
	"net/http"
	"time"
	"yusha/config"
)

// MethodNotAllowedInProxy 初始化异常信息
var (
	MethodNotAllowedInProxy = errors.New("the reverse proxy does not support request methods other than get and post")
)

// 全局 http client
var yuShaHttpClient = &yuShaClient{http.Client{Timeout: time.Second * time.Duration(config.Yusha.Timeout)}}

// 组合 goland 底层 http.Client 结构, 进行底层方法的重写
type yuShaClient struct {
	http.Client
}

// Proxy 代理转发函数
func Proxy(r *http.Request) (resp *http.Response, err error) {
	switch r.Method {
	case http.MethodGet:
		return yuShaHttpClient.Get(r)
	case http.MethodPost:
		return yuShaHttpClient.Post(r)
	default:
		return nil, MethodNotAllowedInProxy
	}
}

// Get 重新实现 golang 底层 http.Client Get 方法
func (ysc *yuShaClient) Get(r *http.Request) (resp *http.Response, err error) {
	return ysc.Do(r)
}

// Post 重新实现 golang 底层 http.Client Post 方法
func (ysc *yuShaClient) Post(r *http.Request) (resp *http.Response, err error) {
	return ysc.Do(r)
}
