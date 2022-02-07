package main

import (
	"bufio"
	"go-proxy-checker/core"
	"io"
	"log"
	"sync"
)

var workerWait = &sync.WaitGroup{}
var inProxyChan = make(chan string, 1000)
var outProxyChan = make(chan string, 1000)
var resultChan = make(chan CheckResult, 1000)

type CheckResult struct {
	proxy  string
	result bool
}

func startCheck() {
	go inputProxiesFromFile()
	// start workers
	for i := 0; i < currency; i++ {
		workerWait.Add(1)
		go worker()
	}
	go writeResultToFile()
	go printProcess()
	// wait works done
	workerWait.Wait()
	close(outProxyChan)
	close(resultChan)
}

func inputProxiesFromFile() {
	inputReader := bufio.NewReader(InputFile)
	defer func() {
		close(inProxyChan)
		_ = InputFile.Close()
	}()
	for {
		line, _, err := inputReader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		inProxyChan <- string(line)
	}
}
func writeResultToFile() {
	defer func() {
		_ = OutputFile.Close()
		wg.Done()
	}()
	for proxy := range outProxyChan {
		_, err := OutputFile.WriteString(proxy + "\n")
		if err != nil {
			panic(err)
		}
	}
}
func printProcess() {
	defer wg.Done()
	var success int
	var count int
	for r := range resultChan {
		count += 1
		if r.result {
			success += 1
			log.Printf("[now: %d / total: %d / success:%d ] proxy %s is vaild\n", count, total, success, r.proxy)
		} else {
			log.Printf("[now: %d / total: %d / success:%d ] proxy %s is invalid\n", count, total, success, r.proxy)
		}
	}
}
func worker() {
	defer workerWait.Done()
	for proxy := range inProxyChan {
		result := core.CheckProxy(proxy)
		if result {
			outProxyChan <- proxy
		}
		resultChan <- CheckResult{
			proxy:  proxy,
			result: result,
		}
	}
}
