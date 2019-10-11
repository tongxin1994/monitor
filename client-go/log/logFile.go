package log

import (
	"bufio"
	"client-go/getflag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

//CreateLogFile create log Files with timestamp
func CreateLogFile() (*os.File, error) {
	//get log dir from cmd

	// flag.Parse()
	fileName := fmt.Sprintf("k8s-check-%s", time.Now().Format("20060102"))
	return os.Create(fileName)
}

//GetLogFile returns pwd of log file
func GetLogFile(file *os.File) string {
	return filepath.Join(getflag.LogDir, file.Name())
}

//WriteLog write check results to log file
func WriteLog(file *os.File, content string) {
	filePath := GetLogFile(file)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(content)
	write.Flush()
}
