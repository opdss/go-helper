package files

import (
	"errors"
	"io"
	"testing"
)

// ReadLineString 按行读取文件，返回io.EOF 表示读完
func TestReadLineString(t *testing.T) {
	err := ReadLineString("./read.go", func(data string, line int) error {
		//fmt.Println(line, " => ", data)
		return nil
	})
	if err != nil && !errors.Is(err, io.EOF) {
		t.Errorf("ReadLineString error:%v", err)
	}
}
