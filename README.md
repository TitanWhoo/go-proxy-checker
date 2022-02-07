# Go Proxy Checker

High Performance HTTP(S) Proxy Checker Written in GO

It can Batch check whether your HTTP/HTTPS proxies is valid and can hide your IP address by request [httpbin.org/get](https://httpbin.org/get)

## Build

clone the project code and use `go build`  command to build

## Usage of Go Proxy Checker

```shell
 Usage of Go Proxy Checker:
  -h    read usage of this tool
  -c int
        concurrent number of proxy checking (default 2000)
  -http
        use http instead of https when checking
  -i string
        the input proxy file (default "input.txt")
  -o string
        the output proxy file (default "output.txt")
```

## Example

### #1 check proxies without HTTPS and input/output file name

```shell
go-proxy-checker -http -i fresh_list.txt -o success.txt
```

### #2 Specify the number of concurrent (default 2000)

```
go-proxy-checker -c 100
```

