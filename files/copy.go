package files

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func Copy(dst, src string) (int64, error) {
	destF, destE := os.Stat(dst)
	srcF, err := os.Stat(src)
	if err != nil {
		return 0, err
	}
	if srcF.IsDir() {
		if destE == nil {
			if !destF.IsDir() {
				return 0, fmt.Errorf("the target [%s] already exists", dst)
			}
		}
		return copyDirToDir(dst, src)
	} else {
		if destE == nil {
			if destF.IsDir() {
				return copyFileToDir(dst, src)
			}
		}
		return copyFileToFile(dst, src)
	}
}

func CopyFileToDir(destDir, srcFile string) (n int64, err error) {
	var srcF os.FileInfo
	srcF, err = os.Stat(srcFile)
	if err != nil || srcF.IsDir() {
		return 0, fmt.Errorf("copy source[%s] error", srcFile)
	}

	var destF os.FileInfo
	destF, err = os.Stat(destDir)
	if err != nil {
		err = os.MkdirAll(destDir, os.ModePerm)
		if err != nil {
			return 0, fmt.Errorf("create dir[%s] error: %s", destDir, err)
		}
	} else {
		if !destF.IsDir() {
			return 0, fmt.Errorf("object [%s] not dir", destDir)
		}
	}
	return copyFileToDir(destDir, srcFile)
}

func CopyFileToFile(destFile, srcFile string) (n int64, err error) {
	var srcF os.FileInfo
	srcF, err = os.Stat(srcFile)
	if err != nil || srcF.IsDir() {
		return 0, fmt.Errorf("copy source[%s] error", srcFile)
	}
	return copyFileToFile(destFile, srcFile)
}

func CopyDirToDir(destDir, srcDir string) (n int64, err error) {
	var srcF os.FileInfo
	srcF, err = os.Stat(srcDir)
	if err != nil || !srcF.IsDir() {
		return 0, fmt.Errorf("copy source[%s] error", srcDir)
	}
	return copyDirToDir(destDir, srcDir)
}

func copyFileToDir(destDir, srcFile string) (n int64, err error) {
	var srcFp *os.File
	srcFp, err = os.Open(srcFile)
	if err != nil {
		return
	}
	defer func() {
		if _err := srcFp.Close(); _err != nil || err == nil {
			err = _err
		}
	}()
	var destFp *os.File
	destFp, err = os.Create(filepath.Join(destDir, filepath.Base(srcFile)))
	if err != nil {
		return
	}
	defer func() {
		if _err := destFp.Close(); _err != nil || err == nil {
			err = _err
		}
	}()
	n, err = io.Copy(destFp, srcFp)
	return
}

func copyDirToDir(destDir, srcDir string) (n int64, err error) {
	var destF os.FileInfo
	destF, err = os.Stat(destDir)
	if err != nil {
		err = os.MkdirAll(destDir, os.ModePerm)
		if err != nil {
			return 0, fmt.Errorf("create dir [%s] error: %s", destDir, err)
		}
	} else if !destF.IsDir() {
		return 0, errors.New("object dir error：" + destDir)
	}
	err = filepath.Walk(srcDir, func(path string, info fs.FileInfo, err error) error {
		if srcDir == path {
			return nil
		}
		if err != nil {
			return err
		}
		if !info.IsDir() {
			_n, _err := copyFileToFile(filepath.Join(destDir, strings.TrimPrefix(path, srcDir)), path)
			n += _n
			return _err
		}
		return nil
	})
	return
}

func copyFileToFile(destFile, srcFile string) (n int64, err error) {
	_, err = os.Stat(destFile)
	if err == nil {
		return 0, fmt.Errorf("the target [%s] already exists", destFile)
	}

	dir := filepath.Dir(destFile)
	var dirF os.FileInfo
	dirF, err = os.Stat(dir)
	if err != nil {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return 0, fmt.Errorf("create dir[%s] error: %s", dir, err)
		}
	} else if !dirF.IsDir() {
		return 0, fmt.Errorf("object file[%s] error", destFile)
	}

	var srcFp *os.File
	srcFp, err = os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	defer func() {
		if _err := srcFp.Close(); _err != nil || err == nil {
			err = _err
		}
	}()

	var destFp *os.File
	destFp, err = os.Create(destFile)
	if err != nil {
		return 0, err
	}
	defer func() {
		if _err := destFp.Close(); _err != nil || err == nil {
			err = _err
		}
	}()

	n, err = io.Copy(destFp, srcFp)
	return
}
