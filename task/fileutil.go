package task

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

//newreader 是最快的io方式

func FileRead() {
	open, _ := os.Open("/Users/oker/data.exp")
	fmt.Println(open.Name())
	defer open.Close()
	inputReader := bufio.NewReader(open)
	i := 0
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			return
		}
		i++
		fmt.Printf("第 %v 行:%s", i, inputString)
	}
}

func FileWrite() {
	name := "output.dat"
	if Exist(name) {
		os.Remove(name)
	}
	outputFile, outputError := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	for i := 0; i < 10; i++ {
		now := time.Now()
		unix := now.UnixNano()
		outputWriter.WriteString(strconv.FormatInt(unix, 10) + "\n")
	}
	outputWriter.Flush()
}

func FileWriteAppend() {
	outputFile, outputError := os.OpenFile("output.dat", os.O_RDWR|os.O_APPEND, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"
	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}

func IOread() {
	if contents, err := ioutil.ReadFile("/Users/oker/data.exp"); err == nil {
		fmt.Println(string(contents))
	}
}

func IOWrite() {
	data := []byte("funtstet")
	if ioutil.WriteFile("output.dat", data, 0644) == nil {
		fmt.Println("写入文件成功:", data)
	}
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func CreateDir(path string) {
	os.MkdirAll(path, os.ModePerm)
}

func CreateFileDir(filePath string) string {
	err := os.MkdirAll(filepath.Dir(filePath), 0666)
	if err == nil {
		return filePath
	}
	return filePath
}

func CreateDateDir(Path string, time time.Time) string {
	folderName := time.Format("20060102")
	folderPath := filepath.Join(Path, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.MkdirAll(folderPath, os.ModePerm)
		os.Chmod(folderPath, os.ModePerm)
	}
	return folderPath
}
