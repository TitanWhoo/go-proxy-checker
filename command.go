package main

import (
	"flag"
	"fmt"
	"go-proxy-checker/core"
	"os"
)

var InputFile *os.File
var OutputFile *os.File
var UseHTTP bool
var inputFilePath string
var outputFilePath string
var currency int
var help bool

func init() {
	flag.BoolVar(&help, "h", false, "read usage of this tool")
	flag.StringVar(&inputFilePath, "i", "input.txt", "the input proxy file")
	flag.StringVar(&outputFilePath, "o", "output.txt", "the output proxy file")
	flag.IntVar(&currency, "c", 2000, "concurrent number of proxy checking")
	flag.BoolVar(&UseHTTP, "http", false, "use http instead of https when checking")
}
func initCheckUrl() {
	if UseHTTP {
		core.CheckURL = "http://httpbin.org/get"
	} else {
		core.CheckURL = "https://httpbin.org/get"
	}
}
func startCommand() {
	var err error
	flag.Parse()
	InputFile, err = os.Open(inputFilePath)
	if err != nil {
		fmt.Println(err.Error())
		flag.Usage()
		os.Exit(-1)
	}
	OutputFile, err = os.OpenFile(outputFilePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err.Error())
		flag.Usage()
		os.Exit(-1)
	}
	if help {
		flag.Usage()
		os.Exit(0)
	}
	initCheckUrl()
}
