// Author stone-bird created on 2021/8/8 11:50.
// Email 1245863260@qq.com or g1245863260@gmail.com.
// Use of more goroutine write string line to file
package writefile

import (
	"os"
	"sync"
)

// make a writer use to many go goroutine write string in target file
type FileWriter struct {
	FileName string
	Lock     sync.Mutex
	File     *os.File
}

// make a new file writer struck
func NewWriter(fileName string) (*FileWriter, error) {
	var fw FileWriter
	fw.FileName = fileName
	file, err := os.OpenFile(fw.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	fw.File = file
	return &fw, nil
}

// write info into the fileName
func (fw *FileWriter) WriteString(info string) {
	fw.Lock.Lock()
	_, _ = fw.File.WriteString(info + "\n")
	fw.Lock.Unlock()
}
