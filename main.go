package main

import (
	"go-proxy-checker/core"
	"log"
	"os"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}
var err error
var total int

func handleInputFile() {
	log.Println("start handle with the input file.")
	total = core.FileDuplicateHandle(InputFile)
	InputFile, err = os.Open(InputFile.Name())
	core.HandleError(err)
	log.Printf("total unchecked proxies: %d, start check after 3 seconds...\n", total)
	if currency > total {
		currency = total
	}
	time.Sleep(3 * time.Second)
}

func handleOutputFile() {
	log.Println("start handle with the output file.")
	OutputFile, err = os.Open(OutputFile.Name())
	core.HandleError(err)
	total = core.FileDuplicateHandle(OutputFile)
	log.Printf("all done! total checked proxies: %d\n", total)
}

func main() {
	wg.Add(2)
	startCommand()
	handleInputFile()
	startCheck()
	wg.Wait()
	handleOutputFile()
}
