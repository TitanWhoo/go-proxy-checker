# Go Proxy Checker

[中文版文档](https://github.com/titanhw/go-proxy-checker/blob/master/README_zh.md)

High Performance HTTP(S) Proxy Checker Written in GO

It can Batch check whether your proxies is valid and anonymous, by sending request to [httpbin.org/get](https://httpbin.org/get) .

Support **HTTP(S)**, **SOCKS4/5** and Authentication

## Build

download the source code files and use `go build`  command to build

## Usage of Go Proxy Checker

```shell
 Usage of Go Proxy Checker:
  -h    read usage of this tool
  -c int
        concurrent number of proxy checking (default 2000)
  -http
        use http instead of https when checking
  -i string
        the input unchecked proxy list file (default "input.txt")
  -o string
        the output checked proxy list file (default "output.txt")
```

## Example

### Input and output file format

Simple Format

IP Address: Port Number (one record per line)

```
127.0.0.1:8080
127.0.0.101:3128
127.0.0.102:7890
127.0.0.103:8888
...
```

Or with Scheme and Authentication

```
foo:2138
http://foo:2138
http://usename:password@ip:port
socks5://foo:1080
https://bar:3128
```

### Common Usage Example

#### 1. Just check you proxy with default settings

```shell
./go-proxy-checker
```

#### 2. Use http instead of https when checking and specify the input/output filename

```shell
./go-proxy-checker -http -i fresh_list.txt -o success.txt
```

#### 3. Specify the number of concurrent (default 2000)

```shell
./go-proxy-checker -c 100
```
