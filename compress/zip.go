package compress

import (
	"archive/zip"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func Zip(dest, src string) error {
	return zipMatch(dest, src, nil)
}

func ZipMatch(dest, src string, match Match) error {
	return zipMatch(dest, src, match)
}

func zipMatch(dest, src string, match Match) error {
	targetFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer func() {
		if _err := targetFile.Close(); _err != nil && err == nil {
			err = _err
		}
	}()
	zipWriter := zip.NewWriter(targetFile)
	defer func() {
		if _err := zipWriter.Close(); _err != nil && err == nil {
			err = _err
		}
	}()
	var srcF os.FileInfo
	srcF, err = os.Stat(src)
	if err != nil {
		return err
	}
	if !srcF.IsDir() {
		err = zipFile(src, zipWriter, match, strings.TrimRight(src, ps)+ps)
	} else {
		err = zipFolder(src, zipWriter, match, filepath.Dir(src)+ps)
	}
	return err
}

func UnZip(zipFile, dest string) error {
	dest = strings.TrimSuffix(dest, ps)
	//打开要解包的文件，tarFile是要解包的 .tar 文件的路径
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer func() {
		if _err := reader.Close(); _err != nil && err == nil {
			err = _err
		}
	}()
	//用 tr.Next() 来遍历包中的文件，然后将文件的数据保存到磁盘中
	for _, file := range reader.File {
		err = unzip(file, dest)
		if err != nil {
			return err
		}
	}
	return nil
}

func zipFolder(directory string, zw *zip.Writer, match Match, src string) error {
	return filepath.Walk(directory, func(zipPath string, file os.FileInfo, err error) error {
		if file == nil {
			return err
		}
		if file.IsDir() {
			return nil
		}
		return zipFile(zipPath, zw, match, src)
	})
}

func zipFile(sourceFile string, zw *zip.Writer, match Match, src string) (err error) {
	fileName := strings.TrimPrefix(sourceFile, src)
	if match != nil && !match(fileName) {
		return nil
	}
	var sFile *os.File
	sFile, err = os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer func() {
		if _err := sFile.Close(); _err != nil && err == nil {
			err = _err
		}
	}()

	var info os.FileInfo
	info, err = sFile.Stat()
	if err != nil {
		return err
	}
	// 获取压缩头信息
	var header *zip.FileHeader
	header, err = zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	// 指定文件压缩方式 默认为 Store 方式 该方式不压缩文件 只是转换为zip保存
	header.Method = zip.Deflate
	header.Name = fileName
	var fw io.Writer
	fw, err = zw.CreateHeader(header)
	if err != nil {
		return err
	}
	// 写入文件到压缩包中
	_, err = io.Copy(fw, sFile)
	return err
}

func unzip(file *zip.File, dest string) (err error) {
	var rc io.ReadCloser
	rc, err = file.Open()
	if err != nil {
		return err
	}
	defer func() {
		if _err := rc.Close(); _err != nil && err == nil {
			err = _err
		}
	}()
	//先创建目录
	fileName := dest + ps + strings.TrimPrefix(file.Name, ps)
	dir := path.Dir(fileName)
	_, err = os.Stat(dir)
	//如果err 为空说明文件夹已经存在，就不用创建
	if err != nil {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	fw, er := os.Create(fileName)
	if er != nil {
		return er
	}
	defer func() {
		if _err := fw.Close(); _err != nil && err == nil {
			err = _err
		}
	}()
	// 写入解压后的数据
	_, err = io.CopyN(fw, rc, int64(file.UncompressedSize64))
	return err
}
