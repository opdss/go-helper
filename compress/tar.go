package compress

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func Pack(dest, src string) (err error) {
	return packMatch(dest, src, nil)
}

func PackMatch(dest, src string, match Match) (err error) {
	return packMatch(dest, src, match)
}

func packMatch(dest, src string, match Match) (err error) {
	var destF *os.File
	destF, err = os.Create(dest)
	if err != nil {
		return err
	}
	defer func() {
		if _err := destF.Close(); _err != nil && err == nil {
			err = _err
		}
	}()
	gzipWriter := gzip.NewWriter(destF)
	defer func() {
		if _err := gzipWriter.Close(); _err != nil && err == nil {
			err = _err
		}
	}()
	tarWriter := tar.NewWriter(gzipWriter)
	defer func() {
		if _err := tarWriter.Close(); _err != nil && err == nil {
			err = _err
		}
	}()

	var srcF os.FileInfo
	srcF, err = os.Stat(src)
	if err != nil {
		return err
	}
	if !srcF.IsDir() {
		err = tarFile(src, tarWriter, match, filepath.Dir(src)+ps)
	} else {
		err = tarFolder(src, tarWriter, match, strings.TrimRight(src, ps)+ps)
	}
	return err
}

func Unpack(tarFile, dest string) (err error) {
	dest = strings.TrimSuffix(dest, ps)
	var fr *os.File
	fr, err = os.Open(tarFile)
	if err != nil {
		return err
	}
	defer func() {
		if _err := fr.Close(); _err != nil && err == nil {
			err = _err
		}
	}()
	// 使用gzip解压
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}
	defer func() {
		if _err := gr.Close(); _err != nil && err == nil {
			err = _err
		}
	}()
	// 创建tar reader
	tarReader := tar.NewReader(gr)
	// 循环读取
	for {
		header, _err := tarReader.Next()
		switch {
		// 读取结束
		case _err == io.EOF:
			return nil
		case _err != nil:
			return _err
		case header == nil:
			continue
		}
		// 因为指定了解压的目录，所以文件名加上路径
		targetFullPath := filepath.Join(dest, header.Name)
		dir := path.Dir(targetFullPath)
		_, err = os.Stat(dir)
		//如果err 为空说明文件夹已经存在，就不用创建
		if err != nil {
			err = os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				return err
			}
		}
		// 根据文件类型做处理，这里只处理目录和普通文件，如果需要处理其他类型文件，添加case即可
		switch header.Typeflag {
		case tar.TypeDir:
			_, err = os.Stat(targetFullPath)
			//如果err 为空说明文件夹已经存在，就不用创建
			if err != nil {
				if err = os.MkdirAll(targetFullPath, os.ModePerm); err != nil {
					return err
				}
			}
		case tar.TypeReg:
			// 是普通文件，创建并将内容写入
			file, err := os.OpenFile(targetFullPath, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			_, err = io.Copy(file, tarReader)
			// 循环内不能用defer，先关闭文件句柄
			if _err := file.Close(); _err != nil {
				return _err
			}
			// 这里再对文件copy的结果做判断
			if err != nil {
				return err
			}
		}
	}
}

func tarFolder(directory string, zw *tar.Writer, match Match, src string) error {
	return filepath.Walk(directory, func(zipPath string, file os.FileInfo, err error) error {
		if file == nil {
			return err
		}
		if file.IsDir() {
			return nil
		}
		return tarFile(zipPath, zw, match, src)
	})
}

func tarFile(sourceFile string, zw *tar.Writer, match Match, src string) (err error) {
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
	var header *tar.Header
	header, err = tar.FileInfoHeader(info, "")
	if err != nil {
		return err
	}
	header.Name = fileName
	err = zw.WriteHeader(header)
	if err != nil {
		return err
	}
	// 写入文件到压缩包中
	_, err = io.Copy(zw, sFile)
	return err
}
