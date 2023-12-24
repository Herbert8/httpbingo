---
title: HTTPBinGo v1.0.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.17"

---

# HTTPBinGo

> v1.0.0

本程序是 [Kenneth Reitz](https://github.com/kennethreitz) 编写的 [HTTPBin](https://httpbin.org/) 的 Golang 实现。

[HTTPBin](https://httpbin.org/) 是一个 `HTTP Request & Response Service`，通过 Response 返回 Request 的信息，也可以根据请求指定的参数来返回需要的 HTTP Response，以便 HTTP 客户端可以方便地看到发送请求的信息，也可以用于模拟特定的服务端返回，以便做针对性的处理。

### 使用方法

| 参数 | 说明                     |
| :--: | ------------------------ |
| `-h` | 在控制台显示基础帮助信息 |
| `-p` | 指定监听端口             |

Base URLs:

* <a href="http://localhost:8080">正式环境: http://localhost:8080</a>

# Authentication

# 认证

## GET Basic Auth 认证

GET /basic-auth/{user}/{passwd}

[Basic Auth](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Authentication) 身份认证测试。

通过 Path 中的 `username` 和 `password` 参数指定服务端按这两个参数进行权限校验。如果没有校验通过，则返回 401。

对于浏览器访问的场景，会弹出输入用户名和密码的输入框。

- cURL 示例

```bash
curl "${base_url}/basic-auth/my-username/my-password" -i -u 'my-username:my-password'
```

- HTTP 示例

```http
GET {{base_url}}/basic-auth/my-username/my-password
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user|path|string| 是 |用户名|
|passwd|path|string| 是 |密码|

> 返回示例

> 成功

```json
{
  "authenticated": true,
  "user": "my_user_name"
}
```

> 没有权限

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|没有权限|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» authenticated|boolean|true|none||none|
|» user|string|true|none||none|

## GET Bearer 认证

GET /bearer

Bearer 身份认证测试。

检查 `Authorization` 头中是否具备值为 `Bearer xxx` 形式的数据。`Bearer` **区分大小写**。

如果没有取到值，则返回 401；否则返回 200 并以 JSON 形式展示数据。

- cURL 示例

```bash
curl "${base_url}/bearer" -i -H 'Authorization: Bearer my-token'
```

- HTTP 示例

```http
GET {{base_url}}/bearer
Authorization: Bearer my-token
```

> 返回示例

> 成功

```json
{
  "authenticated": true,
  "user": "my_user_name"
}
```

> 没有权限

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|没有权限|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» authenticated|boolean|true|none||none|
|» user|string|true|none||none|

# 操作 Cookie

## GET 显示客户端 Cookie

GET /cookies

显示客户端请求携带的 Cookie 信息。

- cURL 示例
```bash
curl -iL -c "cookie.txt" \
  "${base_url}/cookies"
```

- HTTP 示例
```http
GET {{base_url}}/cookies
```

> 返回示例

> 成功

```json
{
  "cookies": [
    {
      "Name": "key1",
      "Value": "value1",
      "Path": "",
      "Domain": "",
      "Expires": "0001-01-01T00:00:00Z",
      "RawExpires": "",
      "MaxAge": 0,
      "Secure": false,
      "HttpOnly": false,
      "SameSite": 0,
      "Raw": ""
    },
    {
      "Name": "key2",
      "Value": "value2",
      "Path": "",
      "Domain": "",
      "Expires": "0001-01-01T00:00:00Z",
      "RawExpires": "",
      "MaxAge": 0,
      "Secure": false,
      "HttpOnly": false,
      "SameSite": 0,
      "Raw": ""
    }
  ]
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» cookies|[object]|true|none||none|
|»» Name|string|true|none||none|
|»» Value|string|true|none||none|
|»» Path|string|true|none||none|
|»» Domain|string|true|none||none|
|»» Expires|string|true|none||none|
|»» RawExpires|string|true|none||none|
|»» MaxAge|integer|true|none||none|
|»» Secure|boolean|true|none||none|
|»» HttpOnly|boolean|true|none||none|
|»» SameSite|integer|true|none||none|
|»» Raw|string|true|none||none|

## GET 为客户端设置 Cookie

GET /cookies/set

通过 Params 参数指定要在 Response 中返回的 Cookie。

其中每个**键值对**表示一个 Cookie 设置。

服务器返回 Status Code 为 302 的 Response，通过 Set-Cookie 头来为客户端设定 Cookie，同时引导 HTTP Client 访问 /cookies 查看本地 Cookie。

- cURL 示例
```bash
curl -iL -c "cookie.txt" \
  "${base_url}/cookies/set?key1=value1&key2=value2"
```

- HTTP 示例
```http
GET {{base_url}}/cookies/set?key1=value1&key2=value2
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|key1|query|string| 否 |none|
|key2|query|string| 否 |none|

> 返回示例

> 成功

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|302|[Found](https://tools.ietf.org/html/rfc7231#section-6.4.3)|成功|Inline|

### 返回数据结构

## GET 设置 Cookie 的详细信息

GET /cookies/set-detail/{key}/{value}

通过 Params 参数指定要在 Response 中返回的 Cookie。

通过 `path` 中的 `key` 和 `value` 指定创建 Cookie 的 `key` 和 `value`，其它参数通过 Query Params 传递。

服务器返回 Status Code 为 302 的 Response，通过 Set-Cookie 头来为客户端设定 Cookie，同时引导 HTTP Client 访问 /cookies 查看本地 Cookie。

- cURL 示例
```bash
curl -iL -c "cookie.txt" \
  "${base_url}/cookies/set-detail/key/value?secure=1&httponly=1"
```

- HTTP 示例
```http
GET {{base_url}}/cookies/set-detail/key/value?secure=1&httponly=1
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|key|path|string| 是 |指定要设置的 Cookie 的 Key。|
|value|path|string| 是 |指定要设置的 Cookie 的 Value。|
|secure|query|integer| 否 |指定 Cookie 的安全性，值为 1 时，Cookie 只有在使用 HTTPS 的时候才能携带到服务端。|
|httponly|query|integer| 否 |指定 Cookie 必须由 HTTP 的 Response 指定，无法通过脚本等手段设置。|

> 返回示例

> 成功

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|302|[Found](https://tools.ietf.org/html/rfc7231#section-6.4.3)|成功|Inline|

### 返回数据结构

## GET 删除客户端 Cookie

GET /cookies/delete

通过 Params 参数指定要在 Response 中进行清理的 Cookie。

其中每个**键值对**中的 Key 表示一个要清理的 Cookie。

- cURL 示例
```bash
curl -iL  "${base_url}/cookies/delete?key1&key2"
```

- HTTP 示例
```http
GET {{base_url}}/cookies/delete?key1&key2
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|key1|query|string| 否 |none|
|key2|query|string| 否 |none|

> 返回示例

> 成功

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|302|[Found](https://tools.ietf.org/html/rfc7231#section-6.4.3)|成功|Inline|

### 返回数据结构

# 获取任意请求的信息

## GET 获取请求的所有信息

GET /anything

获取请求的详细信息。

- cURL 示例
```bash
curl -i "${base_url}/anything"
```

- HTTP 示例
```http
GET {{base_url}}/anything
```

> 返回示例

> 成功

```json
{
  "args": {},
  "data": "",
  "files": {},
  "form": {},
  "headers": {
    "Accept": "*/*",
    "Host": "localhost:8080",
    "User-Agent": "curl/8.1.2"
  },
  "json": {},
  "method": "GET",
  "url": "/anything",
  "server_endpoints": [
    "192.168.50.100:8080",
    "192.168.1.200:8080",
    "192.168.1.7:8080",
    "192.168.101.1:8080"
  ],
  "client": "127.0.0.1:54146"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» args|object|true|none||none|
|» data|string|true|none||none|
|» files|object|true|none||none|
|» form|object|true|none||none|
|» headers|object|true|none||none|
|»» Accept|string|true|none||none|
|»» Host|string|true|none||none|
|»» User-Agent|string|true|none||none|
|» json|object|true|none||none|
|» method|string|true|none||none|
|» url|string|true|none||none|
|» server_endpoints|[string]|true|none||none|
|» client|string|true|none||none|

# 重定向

## GET 重定向

GET /redirect-to

指定服务端模拟 30X Redirect。

- cURL 示例
```bash
curl -iL "${base_url}/redirect-to?url=${base_url}/anything&status_code=302"
```

- HTTP 示例
```http
GET {{base_url}}/redirect-to?url={{base_url}}/anything&status_code=302
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|url|query|string| 否 |跳转到的 URL。|
|status_code|query|integer| 否 |状态码。|

> 返回示例

> 302 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|302|[Found](https://tools.ietf.org/html/rfc7231#section-6.4.3)|成功|Inline|

### 返回数据结构

## GET 页面重定向

GET /web-redirect-to

通过返回 HTML 中的 <meta> 标签使页面重定向。

如：
```html
<html>
<meta http-equiv="refresh" content="0;url=http://www.baidu.com/">
</html>
```

- cURL 示例
```bash
curl -iL "${base_url}/web-redirect-to?url=http://baidu.com&delay=3"
```

- HTTP 示例
```http
GET {{base_url}}/web-redirect-to?url=http://baidu.com&delay=3
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|url|query|string| 否 |跳转到的 URL。|
|delay|query|integer| 否 |延迟的秒数。|

> 返回示例

> 成功

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 动态数据

## GET 按指定时长模拟延时返回

GET /delay/{delay}

让服务器在等待指定时间后返回，如不指定，默认值为 3 秒，最大值为 10 秒。等待时间范围：(0, 10]

- cURL 示例

```bash
curl "${base_url}/delay/5"
```

​    

- HTTP 示例

```http
GET {{base_url}}/delay/5
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|delay|path|integer| 是 |指定延时时间（秒）。|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST Base64 解码（Body 传参）

POST /base64

Base64 解码，对 Body 中传递的 Base64 数据进行解码。

- cURL 示例

```bash
curl --data-urlencode 'base64=SFRUUEJJTl9HTyBpcyBhd2Vzb21l' \
    "${base_url}/base64"
```

​    

- HTTP 示例

```http
POST {{base_url}}/base64
Content-Type: application/x-www-form-urlencoded

base64=SFRUUEJJTl9HTyBpcyBhd2Vzb21l
```

> Body 请求参数

```yaml
base64: SFRUUEJJTl9HTyBpcyBhd2Vzb21l

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» base64|body|string| 否 |Base64 编码的数据|

> 返回示例

> 成功

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET Base64 解码（Path 传参）

GET /base64/{base64-value}

Base64 解码，对 Path 中传递的 Base64 数据进行解码。

- cURL 示例

```bash
curl "${base_url}/base64/SFRUUEJJTl9HTyBpcyBhd2Vzb21l"
```

​    

- HTTP 示例

```http
GET {{base_url}}/base64/SFRUUEJJTl9HTyBpcyBhd2Vzb21l
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|base64-value|path|string| 是 |none|

> 返回示例

> 成功

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 指定 Response Header（GET 传参）

GET /response-headers

指定需要在 Response 中返回的 Header。

- cURL 示例

```bash
curl -i "${base_url}/response-headers?key1=value1&key2=value2"
```

​    

- HTTP 示例

```http
GET {{base_url}}/response-headers?key1=value1&key2=value2
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|key1|query|string| 否 |none|
|key2|query|string| 否 |none|

> 返回示例

> 成功

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 指定 Response Header（POST 传参）

POST /response-headers

指定需要在 Response 中返回的 Header。

- cURL 示例

```bash
curl -i "${base_url}/response-headers" \
     --data-urlencode 'key1=value1' \
     --data-urlencode 'key2=value2'
```

​    

- HTTP 示例

```http
POST {{base_url}}/response-headers
Content-Type: application/x-www-form-urlencoded

key1=value1&key2=value2
```

> Body 请求参数

```yaml
key1: value1
key2: value2

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» key1|body|string| 否 |none|
|» key2|body|string| 否 |none|

> 返回示例

> 成功

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 按配置模拟返回数据

POST /data

根据配置，模拟返回数据。

- cURL 示例

```bash
curl "${base_url}" \
     -F 'content_file=@content_file.txt' \
     -F 'content=content-abc' \
     -F 'as_download=1' \
     -F 'content_type=text/plain' \
     -F 'download_filename=filename.txt'
```

> Body 请求参数

```yaml
content_file: string
content: Test text.
as_download: 1,true
download_filename: 我的文件.txt
content_type: application/octet-stream

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» content_file|body|string(binary)| 否 |使用这里上传文件的内容作为返回。|
|» content|body|string| 否 |需要通过 Response 返回的内容。|
|» as_download|body|integer| 否 |返回的内容以下载形式提供。|
|» download_filename|body|string| 否 |指定下载时使用的文件名。|
|» content_type|body|string| 否 |指定通过 Response 返回内容的 Content-Type。|

#### 详细说明

**» content_type**: 指定通过 Response 返回内容的 Content-Type。
如不指定或指定为 auto，则自动检测数据测 Content-Type。

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET SSE 模拟测试

GET /sse

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 状态码

## GET 指定返回状态码

GET /status/{code}

返回指定的 Status Code 及描述。

如果状态码介于 `[300, 400)` 则将 Response 的 `Location` 头指定为 `/anything`

- cURL 示例

```bash
curl -iL "${base_url}/status/302"
```

​    

- HTTP 示例

```http
GET {{base_url}}/status/302
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|code|path|integer| 是 |指定要返回的状态码|

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 数据探测

## POST 探测数据（multipart/form-data 形式）

POST /detect/multipart

检测数据类型。

- cURL 示例

```bash
curl "${base_url}/detect" -F file=@data_file"
```

​    

- HTTP 示例

```http
POST {{base_url}}/detect
Content-Type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW

------WebKitFormBoundary7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="file"; filename="data_file"
Content-Type: application/octet-stream

< data_file

------WebKitFormBoundary7MA4YWxkTrZu0gW--
```

> Body 请求参数

```yaml
file: string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» file|body|string(binary)| 是 |包含数据的文件|

> 返回示例

> 成功

```json
{
  "size": 64057182,
  "Content-Type": "application/zip",
  "content": "PK\u0003\u0004\u0014\u0000\u0000\u0000\b\u0000{dlQ,�f �\u0000\u0000\u0000\u001f\u0001\u0000\u0000\u0019\u0000\u001c..."
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» size|integer|true|none||none|
|» Content-Type|string|true|none||none|
|» content|string|true|none||none|

## POST 探测数据（application/x-www-form-urlencoded 形式）

POST /detect/www

检测数据类型。

- cURL 示例

```bash
curl "${base_url}/detect" --data-binary @data_file
```

​    

- HTTP 示例

```http
POST {{base_url}}/detect
Content-Type: application/x-www-form-urlencoded

< data_file
```

> Body 请求参数

```yaml
string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Content-Type|header|string| 否 |none|
|body|body|string(binary)| 否 |none|

> 返回示例

> 成功

```json
{
  "size": 64057182,
  "Content-Type": "application/zip",
  "content": "PK\u0003\u0004\u0014\u0000\u0000\u0000\b\u0000{dlQ,�f �\u0000\u0000\u0000\u001f\u0001\u0000\u0000\u0019\u0000\u001c..."
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» size|integer|true|none||none|
|» Content-Type|string|true|none||none|
|» content|string|true|none||none|

# 数据倾卸

## GET 显示请求的原始信息

GET /dump/request

将请求的原始信息作为 Response 的 Body 返回。

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 图像

## GET 返回 JPEG 图像

GET /image/jpeg

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 返回 PNG 图像

GET /image/png

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 返回 SVG 图像

GET /image/svg

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 返回 WebP 图像

GET /image/webp

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 返回 GIF 图像

GET /image/gif

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 文本及编码

## GET 返回 JSON 数据

GET /json

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 返回 XML 数据

GET /xml

> 返回示例

> 200 Response

```xml
<?xml version="1.0" encoding="UTF-8" ?>
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 返回 HTML 数据

GET /html

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 返回编码为 UTF-8 的内容

GET /encoding/utf8

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 返回 Gzip 数据

GET /gzip

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 文件格式

## GET 返回 Word 97 格式文档

GET /file/word97

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 返回 Excel 97 格式文档

GET /file/excel97

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 返回 PowerPoint 97 格式文档

GET /file/ppt97

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 返回 Word 2000 之后格式文档

GET /file/word

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 返回 Excel 2000 之后格式文档

GET /file/excel

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 返回 PowerPoint 2000 之后格式文档

GET /file/ppt

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 返回 PDF 格式文档

GET /file/pdf

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 数据模型

