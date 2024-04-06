package files

import (
	"bufio"
	"os"
)

// ReadLineString 按行读取文件，返回io.EOF 表示读完
func ReadLineString(fileName string, fn func(line string, idx int) error) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()
	buf := bufio.NewReader(f)
	var line string
	idx := 0
	for {
		line, err = buf.ReadString('\n')
		if err != nil {
			return err
		}
		if err = fn(line[:len(line)-1], idx); err != nil {
			return err
		}
		idx++
	}
}
