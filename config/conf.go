// Copyright 2022 The Yusha Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"yusha/logger"
)

// 默认配置文件路径参数(默认路径为 ./conf/yusha.json)
var defaultProfilePath = "./conf/yusha.json"

// YuShaConf 配置参数结构体
/**
Root 静态资源代理根路径(默认路径 ./html)
Port 监听端口(默认端口 8100)
CertFile TLS 加密需要的证书文件路径
KeyFile TLS 加密需要的密钥文件路径
ProxyAddr 代理地址(可以为ip或者域名)
ProxyApi 代理接口 api 前缀标识
ProxyHeader 代理过程需要传递的 header 内容, 例如 token 等(多个 header 的 key 用 ; 分隔, 例如 token;Content-Type;Connection)
*/
type YuShaConf struct {
	Root        string
	Port        uint16
	CertFile    string
	KeyFile     string
	ProxyAddr   string
	ProxyPort   uint16
	ProxyApi    string
	ProxyHeader string
}

// Yusha 全局配置参数
var Yusha *YuShaConf

// 初始化
func init() {
	Yusha = &YuShaConf{
		Root: "./html",
		Port: 8100,
	}
	_, err := os.Stat(defaultProfilePath)
	if err != nil {
		logger.LogOut(logger.WARN, "No corresponding file found in the default configuration file path : "+defaultProfilePath)
		logger.LogOut(logger.WARN, "Default configuration will be enabled in Yusha")
		return
	}
	b, _ := os.ReadFile(defaultProfilePath)
	err = json.Unmarshal(b, Yusha)
	if err != nil {
		log.Println("Failed to transfer the configuration file content to JSON")
		panic(err)
	}
	if Yusha.ProxyApi != "" && !strings.HasPrefix(Yusha.ProxyApi, "/") {
		Yusha.ProxyApi = "/" + Yusha.ProxyApi
	}
}
