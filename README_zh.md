## Go Proxy Checker

中文文档 | [English](https://github.com/titanhw/go-proxy-checker)

基于Golang的高性能的HTTP/HTTPS代理服务器验证工具

它可以批量检查你的HTTP/HTTPS代理是否有效和匿名，通过发送请求到[httpbin.org/get](https://httpbin.org/get)。

## 构建

下载源代码文件并使用`go build`命令进行构建

## 用法

```shell
go-proxy-checker 支持如下参数：
  -h    读取有关此工具的说明文本
  -c int
        同时进行代理服务器验证的并发数目（默认为2000）
  -http
        使用HTTP请求替代HTTPS进行验证（只验证是否支持HTTP）
  -i string
        待验证的代理服务器文本列表 (default "input.txt")
  -o string
        输出验证后的代理服务验证列表 (default "output.txt")
```

## 使用示例

### 输入和输出文件格式

你需要确认你要验证的代理服务器文件的格式如下：

IP地址:端口号（每行一条记录）

```
127.0.0.1:8080
127.0.0.101:3128
127.0.0.102:7890
127.0.0.103:8888
...
```

### 常见的使用示例

#### 1. 直接运行程序，使用默认参数验证代理 

```shell
./go-proxy-checker
```

#### 2. 只检测代理服务器是否支持HTTP，指定输入和输出文件路径

```shell
./go-proxy-checker -http -i fresh_list.txt -o success.txt
```

#### 3. 指定并发的数量（默认为2000）。

```shell
./go-proxy-checker -c 100
```

