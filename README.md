# httpbingo



[TOC]



### 说明

本程序是类似 `httpbin` 的 Golang 实现。 用于以 JSON 形式回显收到的 HTTP Request 的信息，也可以根据请求指定的参数来返回需要的 HTTP Response，以便 HTTP 客户端可以方便地看到发送请求的信息，也可以用于模拟特定的服务端返回，以便做针对性的处理。



### 使用方法

| 参数 | 说明                     |
| :--: | ------------------------ |
| `-h` | 在控制台显示基础帮助信息 |
| `-p` | 指定监听端口             |



### 功能说明



#### [`/anything/{anything}`](anything)

返回 Request 的详细信息。



#### /cookies

##### [`/`](cookies)

返回 Request 中的 Cookie



##### `/set`

指定服务器通过 Response 返回指定的 Cookie。服务器返回 Status Code 为 302 的 Response，通过 Set-Cookie 头来为客户端设定 Cookie，同时引导 HTTP Client 访问 /cookies 查看本地 Cookie。

举例：

```bash
$ curl -iL -c "cookie.txt" "http://${host}:${port}/cookies/set?key1=val1&key2=val2"
```



##### `/set-detail/{name}/{value}`

基本作用与 `/cookies/set` 相同，但可以进行更加详细的设置。

参数：

| 参数       | 说明                                                         |
| ---------- | ------------------------------------------------------------ |
| `secure`   | 指定 Cookie 的安全性，值为 1 时，Cookie 只有在使用 HTTPS 的时候才能携带到服务端 |
| `httponly` | 指定 Cookie 必须由 HTTP 的 Response 指定，无法通过脚本等手段设置 |


举例：

```bash
$ curl -iL -c "cookie.txt" "http://${host}:${port}/cookies/set-detail/key/value?secure=0&httponly=1"
```



##### `/delete`

删除指定 Cookie。

举例：

```bash
$ curl -iL -c "cookie.txt" "http://${host}:${port}/cookies/delete?key1=&key2="
```



#### [`/redirect-to`](redirect-to?url=http://baidu.com&status_code=302)

指定服务端模拟 30X Redirect

参数：

| 参数          | 说明                         |
| ------------- | ---------------------------- |
| `url`         | 跳转的 URL                   |
| `status_code` | 指定返回的状态麻，必须是 30X |


举例：
```bash
$ curl -iL "http://${host}:${port}/redirect-to?url=http://${host}:${port}/anything&status_code=302"
```


#### [`/basic-auth/{user}/{passwd}`](basic-auth/user/passwd)

提供 Baseic Auth 认证。

举例：

```bash
$ curl -i -u 'user:passwd' "http://${host}:${port}/basic-auth/user/passwd"
```

 

#### [`/delay/{delay}`](delay/3)

让服务器在等待指定时间后返回，如不指定，默认值为 3 秒，最大值为 10 秒。等待时间范围：(0, 10]

举例：

```bash
$ curl "http://${host}:${port}/delay/5"
```

​    

#### [`/base64`，`/base64/{value}`](base64/SFRUUEJJTl9HTyBpcyBhd2Vzb21l)

Base64 解码。

举例：

```bash
# GET 方式
$ curl "http://${host}:${port}/base64/SFRUUEJJTl9HTyBpcyBhd2Vzb21l"

# POST 方式
$ curl --data-urlencode 'base64=SFRUUEJJTl9HTyBpcyBhd2Vzb21l' \
       "http://${host}:${port}/base64"
```



#### [`/data`](download?content=MyTestText&type=application/octet-stream&filename=myfile.txt)

根据指定参数，模拟返回内容。

参数：

| 参数                    | 说明                                                         |
| ----------------------- | ------------------------------------------------------------ |
| `content`               | 内容                                                         |
| `content-type`          | 指定 Content-Type，参考：[HTTP Content-Type 常用对照表](https://tool.oschina.net/commons) |
| `content-type` 默认值   | Content-Type 的默认值为 `application/octet-stream`           |
| `content-type` 自动检测 | Content-Type 为 `auto` 时，自动根据 Body 内容检测类型        |
| `content-file`          | 如果指定了 `content-file`，则从指定文件读取内容              |

举例：

```bash
# GET 方式
$ curl -i "http://${host}:${port}/download?content=abcxyz123&content-type=application/octet-stream"

# POST 方式
$ curl -i "http://${host}:${port}/download" \
       --data-urlencode 'content=abcxyz123' \
       --data-urlencode 'content-type=application/octet-stream' \
       --data-urlencode 'content-file=/tmp/test.txt'
```



#### [`/download`](download?content=MyTestText&type=application/octet-stream&filename=myfile.txt)

模拟内容下载。

参数：

| 参数                    | 说明                                                         |
| ----------------------- | ------------------------------------------------------------ |
| `content`               | 内容                                                         |
| `content-type`          | 指定 Content-Type，参考：[HTTP Content-Type 常用对照表](https://tool.oschina.net/commons) |
| `content-type` 默认值   | Content-Type 的默认值为 `application/octet-stream`           |
| `content-type` 自动检测 | Content-Type 为 `auto` 时，自动根据 Body 内容检测类型        |
| `filename`              | 指定 Content-Disposition 文件名                              |
| `content-file`          | 如果指定了 `content-file`，则从指定文件读取内容              |

举例：

```bash
# GET 方式
$ curl -i "http://${host}:${port}/download?content=abcxyz123&content-type=application/octet-stream&filename=我的文件.txt"

# POST 方式
$ curl -i "http://${host}:${port}/download" \
       --data-urlencode 'content=abcxyz123' \
       --data-urlencode 'content-type=application/octet-stream' \
       --data-urlencode 'filename=我的文件.txt' \
       --data-urlencode 'content-file=/tmp/test.txt'
```



#### `/detect`

检测数据类型。

举例：

```bash
$ curl --data-binary @test.png "http://${host}:${port}/detect"
```



#### `/status/{code}`

返回指定的 Status Code 及描述。

```bash
$ curl -i "http://${host}:${port}/status/500"
```



#### `/response-headers`

指定 Response 的 Header

举例：

```bash
# GET 方式
$ curl -i "http://${host}:${port}/response-headers?k1=v1&k2=v2"

# POST 方式
$ curl -i "http://${host}:${port}/response-headers" \
       --data-urlencode 'k1=v1' \
       --data-urlencode 'k2=v2'
```

