package utils

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteAppendCache(file_path string, data string) bool {
	var compressedData bytes.Buffer
	gz := gzip.NewWriter(&compressedData)
	_, gerr := gz.Write([]byte(data))
	check(gerr)
	fileHandle, err := os.OpenFile(file_path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	check(err)
	_, err = fileHandle.Write(compressedData.Bytes())
	check(err)
	err = fileHandle.Close()
	return true
}

func ReadCacheData(file_path string) string {
	data, err := ioutil.ReadFile(file_path)
	check(err)
	return string(data)
}
