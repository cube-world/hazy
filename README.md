# Yusha | 尤莎
***
### 功能描述
```
Web 服务器; 反向代理
```
### 简单使用
###### 注意事项
- 默认静态资源代理根目录是编译后可执行文件同级别的名为 html 的文件夹
- 配置文件的存放位置为编译后可执行文件同级别的名为 conf 的文件夹下的 yusha.json 文件, 具体参数如下
```
  Root 静态资源代理根路径(默认路径 ./html)
  Port 监听端口(默认端口 8100)
  CertFile TLS 加密需要的证书文件路径(也就是添加支持 https)
  KeyFile TLS 加密需要的密钥文件路径
  ProxyAddr 代理地址(可以为ip或者域名)
  ProxyApi 代理接口 api 前缀标识, 默认为 '/api', 如果配置文件中将次选项置为空及'', 则不开启反向代理功能
           (例如  '/api', 那么意味着 http://localhost:8100/api/** 之类的 http url 将通过中间件实现代理转发)
  ProxyCertFile 代理接口加密需要的证书文件路径(也就是添加支持 https)
  ProxyKeyFile  代理接口加密需要的密钥文件路径
  TimeOut http 请求代理转发超时时间参数(单位秒, 默认 3 秒)
 ```
编译运行
```
go build main.go 生成可执行文件运行即可
windows 编译环境下提供了 build.bat 双击执行即可生成 .exe 文件(该 .exe 执行时无 cmd 窗口, 其在后台运行, 可在任务管理器中看到 yusha.exe 服务)
启动后通过 ip (localhost 或者 域名) 加端口号进行访问, 弹出 Yusha 的提示页及为部署成功
```
