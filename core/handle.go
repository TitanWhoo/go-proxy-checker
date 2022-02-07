package core

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func FileDuplicateHandle(file *os.File) int {
	tempFile, err := ioutil.TempFile(filepath.Dir(file.Name()), "goproxychecker_tmp")
	HandleError(err)
	var tempMap = make(map[string]struct{}, 500)

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(tempFile)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		str := string(line)
		if _, ok := tempMap[str]; ok {
			continue
		}
		tempMap[str] = struct{}{}
		_, err = writer.WriteString(str + "\n")
		HandleError(err)
	}
	err = writer.Flush()
	HandleError(err)
	err = file.Close()
	HandleError(err)
	err = tempFile.Close()
	HandleError(err)
	//err = os.Remove(file.Name())
	//HandleError(err)
	err = os.Rename(tempFile.Name(), file.Name())
	HandleError(err)
	file, err = os.Open(file.Name())
	HandleError(err)
	return len(tempMap)
}
